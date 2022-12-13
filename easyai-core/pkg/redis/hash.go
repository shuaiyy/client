package redis

import (
	"time"

	"github.com/go-redis/redis"
)

// HGet get a key value from hash
func HGet(c *redis.Client, hashKey, filed string) (value string) {
	cmd := c.HGet(hashKey, filed)
	if cmd.Err() != nil {
		return ""
	}
	return cmd.Val()
}

// HSet set a key value into a hash
func HSet(c *redis.Client, hashKey, filed, value string) (ok bool) {
	cmd := c.HSet(hashKey, filed, value)
	// add: cmd.Val() == 1;
	// update: cmd.Val() == 0
	// we don't care add or update, just set if ok
	return cmd.Err() == nil
}

// HExistKey if a key in hash
func HExistKey(c *redis.Client, hashKey, filed string) bool {
	cmd := c.HExists(hashKey, filed)
	return cmd.Val()
}

// HGetAll get all key-values
func HGetAll(c *redis.Client, hashKey string) (value map[string]string) {
	cmd := c.HGetAll(hashKey)
	if cmd.Err() != nil {
		return nil
	}
	return cmd.Val()
}

// HDeleteAll delete a hash
func HDeleteAll(c *redis.Client, hashKey string) (ok bool) {
	cmd := c.Del(hashKey)
	if cmd.Err() != nil {
		return false
	}
	return cmd.Val() > 0
}

// HSetTTL update a hash ttl
func HSetTTL(c *redis.Client, hashKey string, ttl time.Duration) (ok bool) {
	cmd := c.Expire(hashKey, ttl)
	if cmd.Err() != nil {
		return false
	}
	return cmd.Val()
}

// HExist if a hash exist
func HExist(c *redis.Client, hashKey string) bool {
	cmd := c.Exists(hashKey)
	return cmd.Val() > 0
}
