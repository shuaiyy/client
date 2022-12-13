package auth

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
)

// tokenInfo 令牌信息
type tokenInfo struct {
	AccessToken string `json:"access_token"` // 访问令牌
	TokenType   string `json:"token_type"`   // 令牌类型
	ExpiresAt   int64  `json:"expires_at"`   // 令牌到期时间
}

func (t *tokenInfo) GetAccessToken() string {
	return t.AccessToken
}

func (t *tokenInfo) GetTokenType() string {
	return t.TokenType
}

func (t *tokenInfo) GetExpiresAt() int64 {
	return t.ExpiresAt
}

func (t *tokenInfo) EncodeToJSON() ([]byte, error) {
	return json.Marshal(t)
}

const defaultKey = "easyai-core"

var defaultOptions = options{
	tokenType:     "Bearer",
	expired:       3600 * 24 * 365 * 2,
	signingMethod: jwt.SigningMethodHS512,
	signingKey:    []byte(defaultKey),
	keyfunc: func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(defaultKey), nil
	},
}

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    interface{}
	keyfunc       jwt.Keyfunc
	expired       int
	tokenType     string
}

// Option 定义参数项
type Option func(*options)

// SetSigningMethod 设定签名方式
func SetSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// SetSigningKey 设定签名key
func SetSigningKey(key interface{}) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

// SetKeyfunc 设定验证key的回调函数
func SetKeyfunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyfunc = keyFunc
	}
}

// SetExpired 设定令牌过期时长(单位秒，默认7200)
func SetExpired(expired int) Option {
	return func(o *options) {
		o.expired = expired
	}
}

// New 创建认证实例
func New(cli *redis.Client, opts ...Option) *JWTAuth {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}

	return &JWTAuth{
		opts:  &o,
		redis: cli,
	}
}

// JWTAuth jwt认证
type JWTAuth struct {
	opts  *options
	redis *redis.Client
}

// UserClaims user
type UserClaims struct {
	UserUID  string `json:"user_id"`
	RealName string `json:"real_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken 生成令牌
func (a *JWTAuth) GenerateToken(ctx context.Context, userUID, realName, phone, email string) (TokenInfo, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(a.opts.expired) * time.Second).Unix()

	token := jwt.NewWithClaims(a.opts.signingMethod, &UserClaims{
		UserUID:  userUID,
		RealName: realName,
		Phone:    phone,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt,
			NotBefore: now.Unix(),
			Subject:   userUID,
		},
	})

	tokenString, err := token.SignedString(a.opts.signingKey)
	if err != nil {
		return nil, err
	}

	tokenInfo := &tokenInfo{
		ExpiresAt:   expiresAt,
		TokenType:   a.opts.tokenType,
		AccessToken: tokenString,
	}
	return tokenInfo, nil
}

// 解析令牌
func (a *JWTAuth) parseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, a.opts.keyfunc)
	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	return token.Claims.(*UserClaims), nil
}

// DestroyToken 销毁令牌
func (a *JWTAuth) DestroyToken(ctx context.Context, tokenString string) error {
	claims, err := a.parseToken(tokenString)
	if err != nil {
		return err
	}

	// 如果设定了redis存储，则将失效的token标记在redis中
	key := "destroyed-token:" + tokenString
	if a.redis != nil {
		cmd := a.redis.Set(key, claims.Subject+"-deleted", time.Hour*24*180)
		return cmd.Err()
	}

	return nil
}

// ParseUserInfo 解析用户ID
func (a *JWTAuth) ParseUserInfo(ctx context.Context, tokenString string, noCheckInRedis ...bool) (userUID, realName, phone, email string, err error) {
	if tokenString == "" {
		return "", "", "", "", ErrInvalidToken
	}

	claims, err := a.parseToken(tokenString)
	if err != nil {
		return "", "", "", "", err
	}
	if len(noCheckInRedis) == 0 || !noCheckInRedis[0] {
		key := "destroyed-token:" + tokenString
		cmd := a.redis.Exists(key)
		if err := cmd.Err(); err != nil {
			return "", "", "", "", err
		}
		if cmd.Val() > 0 {
			return "", "", "", "", ErrDestroyedToken
		}
	}
	return claims.UserUID, claims.RealName, claims.Phone, claims.Email, nil
}

// Release 释放资源
func (a *JWTAuth) Release() error {
	return a.redis.Close()
}
