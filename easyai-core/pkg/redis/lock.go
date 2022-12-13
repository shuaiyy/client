package redis

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"time"

	"github.com/go-redis/redis"
)

var luaRelease = redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("del", KEYS[1]) else return 0 end`)

// LockRes ...
type LockRes struct {
	key string
	val string
}

var (
	// ErrNotGetLock is returned when a lock cannot be obtained.
	ErrNotGetLock = errors.New("redis lock: not get lock")
	// ErrLockNotHeld is returned when trying to release an inactive lock.
	ErrLockNotHeld = errors.New("redis lock: not held lock")
)

// RetryLock 获取锁，可指定重试次数，每次重试间隔时间
func RetryLock(c *redis.Client, key string, ttl time.Duration, retryCount int, backOff time.Duration) (*LockRes, error) {
	val, err := randomVal()
	if err != nil {
		return nil, err
	}

	if retryCount == 0 {
		return tryLock(c, key, ttl, val)
	}

	for ; retryCount > 0; retryCount-- {
		res, err := tryLock(c, key, ttl, val)
		if err != nil {
			if err == ErrNotGetLock {
				if backOff > 0 {
					time.Sleep(backOff)
				}
				continue
			}

			return nil, err
		}

		// get lock
		return res, nil
	}

	return nil, ErrNotGetLock
}

// GetLock 获取锁，获取不到立即返回
func GetLock(c *redis.Client, key string, ttl time.Duration) (*LockRes, error) {

	val, err := randomVal()
	if err != nil {
		return nil, err
	}

	return tryLock(c, key, ttl, val)
}

// ReleaseLock 释放锁
func (l *LockRes) ReleaseLock(c *redis.Client) error {

	res, err := luaRelease.Run(c, []string{l.key}, l.val).Int64()
	if err != nil {
		return err
	}

	if res != 1 {
		return ErrLockNotHeld
	}
	return nil
}

func tryLock(c *redis.Client, key string, ttl time.Duration, val string) (*LockRes, error) {
	ok, err := lock(c, key, val, ttl)
	if err != nil {
		return nil, err
	}

	if ok {
		return &LockRes{
			key: key,
			val: val,
		}, nil
	}
	return nil, ErrNotGetLock
}

func lock(c *redis.Client, key, val string, ttl time.Duration) (bool, error) {
	ok, err := c.SetNX(key, val, ttl).Result()
	if err != nil {
		return false, err
	}

	return ok, nil
}

func randomVal() (string, error) {

	tmp := make([]byte, 16)

	if _, err := io.ReadFull(rand.Reader, tmp); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(tmp), nil
}
