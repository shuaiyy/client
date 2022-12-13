package sdk

// Config defines config
type Config struct {
	Endpoint  string // server域名
	APIPrefix string // API 前缀，default /api/v1
	Token     string // 用户的token
	IsDebug   bool   // 是否开启调试模式，默认false
	Timeout   uint   // 超时时间，默认60秒
}

// NewConfig get default config
func NewConfig() *Config {
	config := Config{}
	config.Endpoint = ""
	config.IsDebug = false
	config.Timeout = 60
	config.APIPrefix = "/api/v1"
	return &config
}
