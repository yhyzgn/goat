// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-11-19 11:05
// version: 1.0.0
// desc   : redis 分布式锁

package godis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/yhyzgn/goat/date"
	"github.com/yhyzgn/goat/rand"
)

// Locker ...
type Locker struct {
	pool       *RedisPool    // redis 连接池
	key        Key           // 锁名称
	token      string        // 锁编号
	expiry     time.Duration // 过期时间
	tries      int           // 重试次数
	retryDelay time.Duration // 重试间隔
	start      time.Time     // 被锁时间
	locked     bool          // 是否被锁住
}

// LockerOption ...
type LockerOption interface {
	apply(*Locker)
}

type lockerOption func(*Locker)

func (opt lockerOption) apply(locker *Locker) {
	opt(locker)
}

var unlockScript = redis.NewScript(1, `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	else
		return 0
	end
`)

var touchScript = redis.NewScript(1, `
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("PEXPIRE", KEYS[1], ARGV[2])
	else
		return 0
	end
`)

// NewLocker ...
func NewLocker(pool *RedisPool, key Key, options ...LockerOption) *Locker {
	locker := &Locker{
		pool:       pool,
		key:        key,
		token:      lockerToken(),
		expiry:     6 * time.Second,
		tries:      3,
		retryDelay: 600 * time.Millisecond,
		start:      date.Now(),
		locked:     false,
	}

	// 加载配置项
	for _, opt := range options {
		opt.apply(locker)
	}
	return locker
}

// Lock ...
func (lk *Locker) Lock() (bool, error) {
	// 实现重试功能
	var err error
	for i := 0; i < lk.tries; i++ {
		err = lk.acquire()
		if err == nil {
			break
		}
		time.Sleep(lk.retryDelay)
	}
	if err != nil {
		return false, err
	}

	lk.locked = true
	// 由于上述过期时间并不能解决那种客户端异常退出，又没删除锁的情况
	// 所以需要另起协程设置
	go lk.touch()
	return true, nil
}

// Unlock ...
func (lk *Locker) Unlock() error {
	_, err := lk.pool.call(lk.key, func(conn redis.Conn, realKey string) (interface{}, error) {
		return unlockScript.Do(conn, realKey, lk.token)
	})
	if err != nil {
		return err
	}
	lk.locked = false
	return nil
}

func (lk *Locker) acquire() error {
	_, err := lk.pool.call(lk.key, func(conn redis.Conn, realKey string) (interface{}, error) {
		return conn.Do("SET", realKey, lk.token, "PX", int(lk.expiry/time.Millisecond), "NX")
	})
	if err != nil {
		return err
	}
	if err == redis.ErrNil {
		return fmt.Errorf("the lock [%s]=[%s] is already exist", lk.key, lk.token)
	}
	return err
}

func (lk *Locker) touch() {
	for lk.locked {
		// 计算所剩超时时长：设置的超时时长 - 已过的时长
		expiry := lk.expiry - time.Since(lk.start)
		if expiry < 0 {
			// 已超时，直接删除锁
			_ = lk.Unlock()
			return
		}

		// 重新设置剩余的到期时间
		_, _ = lk.pool.call(lk.key, func(conn redis.Conn, realKey string) (interface{}, error) {
			return redis.Int64(touchScript.Do(conn, realKey, lk.token, int(expiry/time.Millisecond)))
		})

		// 定时设置，默认 1s 设置一次，如果之前设置的到期时间 < 1s，则以到期时间为准
		interval := 1 * time.Second
		if lk.expiry < interval {
			interval = lk.expiry
		}
		time.Sleep(interval)
	}
}

// SetLockerKey ...
func SetLockerKey(key Key) LockerOption {
	return lockerOption(func(locker *Locker) {
		locker.key = key
	})
}

// SetLockerToken ...
func SetLockerToken(token string) LockerOption {
	return lockerOption(func(locker *Locker) {
		locker.token = token
	})
}

// SetLockerExpiry ...
func SetLockerExpiry(expiry time.Duration) LockerOption {
	return lockerOption(func(locker *Locker) {
		locker.expiry = expiry
	})
}

// SetLockerTries ...
func SetLockerTries(tries int) LockerOption {
	return lockerOption(func(locker *Locker) {
		locker.tries = tries
	})
}

// SetLockerRetryDelay ...
func SetLockerRetryDelay(delay time.Duration) LockerOption {
	return lockerOption(func(locker *Locker) {
		locker.retryDelay = delay
	})
}

func lockerToken() string {
	return rand.String(32)
}
