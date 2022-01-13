// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-11-19 11:05
// version: 1.0.0
// desc   : redis 连接池

package godis

import (
	"context"
	"fmt"
	"github.com/yhyzgn/gog"
	"regexp"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/pkg/errors"
)

// RedisPool ...
type RedisPool struct {
	KeyPrefix string
	Pool      *redis.Pool
}

// 回调执行，以便 connection 的统一获取和关闭
func (rp *RedisPool) call(key Key, cb func(conn redis.Conn, realKey string) (interface{}, error)) (interface{}, error) {
	conn := rp.Pool.Get()
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			gog.Error(err)
		}
	}(conn)
	return cb(conn, rp.withPrefix(key))
}

// 带前缀的 key
func (rp *RedisPool) withPrefix(key Key) string {
	if rp.KeyPrefix == "" || strings.HasPrefix(string(key), rp.KeyPrefix) {
		return string(key)
	}
	return fmt.Sprintf("%s:%s", rp.KeyPrefix, string(key))
}

// Ping ping检测
func (rp *RedisPool) Ping() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	coon, err := rp.Pool.GetContext(ctx)
	if err != nil {
		err = errors.WithMessage(err, "Pool.GetContext")
		return
	}
	defer func(coon redis.Conn) {
		err := coon.Close()
		if err != nil {
			gog.Error(err)
		}
	}(coon)

	const content = "Hello Redis! Here goes gateway"
	response, err := redis.String(coon.Do("PING", content))
	if err != nil {
		err = errors.WithMessage(err, "Redis Ping")
		return
	}
	if response != content {
		err = errors.Errorf("wrong response, want %q, got %q", content, response)
		return
	}
	return
}

// Locker 获取分布式锁
func (rp *RedisPool) Locker(key Key, options ...LockerOption) *Locker {
	return NewLocker(rp, key, options...)
}

// Limiter 获取限流器
func (rp *RedisPool) Limiter(client *Client, key Key, quota uint64, period time.Duration, options ...LimiterOption) *Limiter {
	return NewLimiter(client, key, quota, period, options...)
}

// Exists 判断 key 是否存在
func (rp *RedisPool) Exists(key Key) (bool, error) {
	res, err := rp.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		return exists(conn, realKey)
	})
	return res.(bool), err
}

// Expire 设置过期时间
func (rp *RedisPool) Expire(key Key, expiry time.Duration) error {
	_, err := rp.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		return expire(conn, realKey, expiry)
	})
	return err
}

// Keys 按通配符 * 查询一批 key
//
// 查询结果自动去掉已配置的前缀
func (rp *RedisPool) Keys(key Key) ([]string, error) {
	res, err := rp.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		return redis.Strings(conn.Do("KEYS", realKey))
	})
	if err != nil {
		return nil, err
	}
	// 此处获取到的key带有前缀，需要把前缀都去掉
	temp := res.([]string)
	keys := make([]string, len(temp))
	rpx := regexp.MustCompile("^" + rp.KeyPrefix + ":?")
	for i, k := range temp {
		if rpx.MatchString(k) {
			keys[i] = rpx.ReplaceAllString(k, "")
		}
	}
	return keys, nil
}

// Del 删除一些 key
func (rp *RedisPool) Del(key ...Key) error {
	// key 有待处理
	_, err := rp.call("", func(conn redis.Conn, realKey string) (interface{}, error) {
		keys := make([]string, len(key))
		for i, k := range key {
			keys[i] = rp.withPrefix(k)
		}
		return nil, del(conn, keys...)
	})
	return err
}

// NewRedisPool 创建连接池
func NewRedisPool(host string, port, db, maxIdle, maxActive int, password, keyPrefix string) *RedisPool {
	pool := newPool(host, port, db, maxIdle, maxActive, password)
	return &RedisPool{
		KeyPrefix: keyPrefix,
		Pool:      pool,
	}
}

func newPool(host string, port, db, maxIdle, maxActive int, password string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return nil, errors.WithMessage(err, "redis.Dial")
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					_ = c.Close()
					return nil, errors.WithMessage(err, "Do AUTH")
				}
			}
			if db > 0 {
				if _, err := c.Do("SELECT", db); err != nil {
					_ = c.Close()
					return nil, errors.WithMessage(err, "Do SELECT")
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         maxIdle,
		MaxActive:       maxActive,
		IdleTimeout:     240 * time.Second,
		Wait:            true,
		MaxConnLifetime: 0,
	}
}
