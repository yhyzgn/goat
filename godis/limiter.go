// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-28 14:11
// version: 1.0.0
// desc   : 分布式限流器 - 令牌桶方案

package godis

import (
	"github.com/yhyzgn/gog"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var limiterScript = redis.NewScript(1, `
	redis.replicate_commands()

	-- 接收并初始化一些参数
	local key = KEYS[1] -- 令牌桶标识
	local capacity = tonumber(ARGV[1]) -- 令牌桶容量
	local period = tonumber(ARGV[2]) -- 规定一定数量令牌的单位时间，同时也是生成一批令牌的单位时间（s）
	local quota = tonumber(ARGV[3]) -- 单位时间内生成令牌的数量
	local quantity = tonumber(ARGV[4]) or 1 -- 每次需要的令牌数，默认为 1
	local timestamp = tonumber(redis.call('time')[1]) -- 当前时间戳（s）

	-- 判断令牌桶是否存在
	if (redis.call('exists', key) == 0) then
		-- 初始化
		redis.call('hmset', key, 'tokens', capacity, 'timestamp', timestamp)
		-- 设置自动过期失效
		redis.call('expire', key, period)
	else
		-- 计算从上一次更新到现在这段时间内应该要生成的令牌数
		local tokens = tonumber(redis.call('hget', key, 'tokens'))
		local last = tonumber(redis.call('hget', key, 'timestamp'))
		local supply = ((timestamp - last) / period) * quota
		if (supply > 0) then
			-- 重置令牌数量
			tokens = math.min(tokens + supply, capacity)
			redis.call('hmset', key, 'tokens', tokens, 'timestamp', timestamp)
			-- 设置自动过期失效
			redis.call('expire', key, period)
		end
	end

	local result = {}
	local tokens = tonumber(redis.call('hget', key, 'tokens'))
	if (tokens < quantity) then
		-- 令牌数量不足，返回0表示已超过限流，同时返回剩余令牌数
		result = {0, tokens}
	else
		-- 令牌充足
		-- 重置剩余令牌数
		tokens = tokens - quantity
		redis.call('hmset', key, 'tokens', tokens, 'timestamp', timestamp)
		-- 设置自动过期失效
		redis.call('expire', key, period)
		-- 返回当前所需要的令牌数量，同时返回剩余令牌数
		result = {quantity, tokens}
	end

	return result
`)

// Limiter 限流器
type Limiter struct {
	sync.Mutex
	client   *Client       // client
	key      Key           // 令牌桶唯一标识
	capacity uint          // 令牌桶容量，默认值 = 单位时间内生成的令牌数
	period   time.Duration // 规定一定数量令牌的单位时间
	quota    uint          // 单位时间内生成令牌的数量
	quantity uint          // 每次需要的令牌数，默认为 1
}

// LimiterOption ...
type LimiterOption interface {
	apply(*Limiter)
}

type limiterOption func(*Limiter)

func (opt limiterOption) apply(limiter *Limiter) {
	opt(limiter)
}

// NewLimiter 创建一个限流器
func NewLimiter(client *Client, key Key, quota uint, period time.Duration, options ...LimiterOption) *Limiter {
	limiter := &Limiter{
		client:   client,
		key:      key,
		capacity: quota, // 默认值 = 单位时间内生成的令牌数
		period:   period,
		quota:    quota,
		quantity: 1,
	}

	// 加载配置项
	for _, opt := range options {
		opt.apply(limiter)
	}
	return limiter
}

// Acquire 申请令牌
//
// 不需要子key
func (lm *Limiter) Acquire() bool {
	return lm.AcquireWith("")
}

// AcquireWith 申请令牌
//
// 需要子key
//
// key 子key
func (lm *Limiter) AcquireWith(key Key) bool {
	// key = 父key + : + 子key
	if key != "" && !strings.HasPrefix(string(key), ":") && !strings.HasSuffix(string(lm.key), ":") {
		key = ":" + key
	}

	key = lm.key + key

	// 优先使用 lua 脚本
	res, err := lm.acquireWithLuaScript(key)
	if nil != err {
		// 有些版本不支持 lua 脚本，此处进行兼容性处理
		res, err = lm.acquireWithNative(key)
	}
	if nil != err {
		gog.Error(err)
	}
	// res = [quantity, tokens]，即 [当前已申请到的令牌数，令牌桶中还剩余的令牌数]
	return (nil == err || err == redis.ErrNil) && len(res) == 2 && res[0] > 0
}

func (lm *Limiter) acquireWithLuaScript(key Key) ([]int64, error) {
	return redis.Int64s(lm.client.Pool.call(lm.key+key, func(conn redis.Conn, realKey string) (interface{}, error) {
		return limiterScript.Do(conn, realKey, lm.capacity, int(lm.period/time.Second), lm.quota, lm.quantity)
	}))
}

func (lm *Limiter) acquireWithNative(key Key) ([]int64, error) {
	lm.Lock()
	defer lm.Unlock()

	capacity, quota, quantity := lm.capacity, lm.quota, lm.quantity
	timestamp := time.Now().Unix()

	if exist, err := lm.client.Exists(key); err != nil || !exist {
		// 不存在，新建令牌桶
		lt := LimiterToken{Tokens: float64(capacity), Timestamp: timestamp}
		if err = lm.client.JSON.Val.SetWithExpiry(key, lt, lm.period); err != nil {
			return nil, err
		}
	} else {
		// 计算从上一次更新到现在这段时间内应该要生成的令牌数
		var last LimiterToken
		if err = lm.client.JSON.Val.Get(key, &last); nil != err {
			return nil, err
		}
		supply := (float64(timestamp) - float64(last.Timestamp)) / float64(lm.period/time.Second) * float64(quota)
		if supply > 0 {
			// 重置令牌数
			tokens := math.Min(last.Tokens+supply, float64(capacity))
			lt := LimiterToken{Tokens: tokens, Timestamp: timestamp}
			if err = lm.client.JSON.Val.SetWithExpiry(key, lt, lm.period); err != nil {
				return nil, err
			}
		}
	}

	var lt LimiterToken
	if err := lm.client.JSON.Val.Get(key, &lt); nil != err {
		return nil, err
	}

	if lt.Tokens < float64(quantity) {
		// 令牌不够
		return []int64{0, int64(lt.Tokens)}, nil
	}

	// 令牌充足
	// 重置剩余令牌数
	tokens := lt.Tokens - float64(quantity)
	lt = LimiterToken{Tokens: tokens, Timestamp: timestamp}
	if err := lm.client.JSON.Val.SetWithExpiry(key, lt, lm.period); err != nil {
		return nil, err
	}

	// 返回当前所需要的令牌数量，同时返回剩余令牌数
	return []int64{int64(quantity), int64(tokens)}, nil
}

// SetLimiterCapacity 设置桶容量
func SetLimiterCapacity(capacity uint) LimiterOption {
	return limiterOption(func(limiter *Limiter) {
		limiter.capacity = capacity
	})
}

// SetLimiterQuantity 设置每次需要取得的令牌数
func SetLimiterQuantity(quantity uint) LimiterOption {
	return limiterOption(func(limiter *Limiter) {
		limiter.quantity = quantity
	})
}

// LimiterToken 令牌
type LimiterToken struct {
	Tokens    float64 `json:"tokens"`
	Timestamp int64   `json:"timestamp"`
}
