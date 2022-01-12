// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-28 14:11
// version: 1.0.0
// desc   : 分布式限流器 - 令牌桶方案

package godis

import (
	"github.com/yhyzgn/gog"
	"math"
	"strconv"
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
	capacity float64       // 令牌桶容量，默认值 = 单位时间内生成的令牌数
	period   time.Duration // 规定一定数量令牌的单位时间
	quota    float64       // 单位时间内生成令牌的数量
	quantity uint64        // 每次需要的令牌数，默认为 1
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
func NewLimiter(client *Client, key Key, quota uint64, period time.Duration, options ...LimiterOption) *Limiter {
	limiter := &Limiter{
		client:   client,
		key:      key,
		capacity: float64(quota), // 默认值 = 单位时间内生成的令牌数
		period:   period,
		quota:    float64(quota),
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
	return redis.Int64s(lm.client.Pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		return limiterScript.Do(conn, realKey, lm.capacity, int(lm.period/time.Second), lm.quota, lm.quantity)
	}))
}

func (lm *Limiter) acquireWithNative(key Key) ([]int64, error) {
	// 1、先判断令牌桶是否存在
	// - 不存在：初始化令牌桶，并设置过期时间
	// -   存在：计算从上一次更新到现在这段时间内应该要生成的令牌数，更新令牌数量，并设置过期时间
	// 2、查询目前令牌数
	// - 满足所需令牌数：减掉所需令牌数后更新令牌数量，并设置过期时间
	// -        不满足：已超过限流频率
	lm.Lock()
	defer lm.Unlock()

	capacity, quota, quantity := lm.capacity, lm.quota, lm.quantity
	now := time.Now().Unix()

	var (
		temp   string
		tokens float64
		last   int64
	)

	// 1、判断令牌桶是否存在
	if exist, err := lm.client.Exists(key); err != nil || !exist {
		// 不存在，初始化
		// 保存并设置过期时间
		if err = lm.client.String.Hash.Set(key, "tokens", floatToString(capacity)); err != nil {
			return nil, err
		}
		if err = lm.client.String.Hash.Set(key, "timestamp", int64ToString(now)); err != nil {
			return nil, err
		}
		if err = lm.client.Expire(key, lm.period); err != nil {
			return nil, err
		}
	} else {
		// 存在
		// 计算从上一次更新到现在这段时间内应该要生成的令牌数
		if err = lm.client.String.Hash.Get(key, "tokens", &temp); err != nil {
			return nil, err
		}
		tokens = stringToFloat(temp)
		if err = lm.client.String.Hash.Get(key, "timestamp", &temp); err != nil {
			return nil, err
		}
		last = stringToInt64(temp)

		supply := float64(now-last) / lm.period.Seconds() * quota
		if supply > 0 {
			// 重置令牌数量
			tokens = math.Min(tokens+supply, capacity)
			if err = lm.client.String.Hash.Set(key, "tokens", floatToString(tokens)); err != nil {
				return nil, err
			}
			if err = lm.client.String.Hash.Set(key, "timestamp", int64ToString(now)); err != nil {
				return nil, err
			}
			if err = lm.client.Expire(key, lm.period); err != nil {
				return nil, err
			}
		}
	}

	if err := lm.client.String.Hash.Get(key, "tokens", &temp); err != nil {
		return nil, err
	}
	tokens = stringToFloat(temp)
	if tokens < float64(quantity) {
		// 令牌数量不足，返回0表示已超过限流，同时返回剩余令牌数
		return []int64{0, int64(tokens)}, nil
	}
	// 令牌充足
	// 重置剩余令牌数
	tokens -= float64(quantity)
	if err := lm.client.String.Hash.Set(key, "tokens", floatToString(tokens)); err != nil {
		return nil, err
	}
	if err := lm.client.String.Hash.Set(key, "timestamp", int64ToString(now)); err != nil {
		return nil, err
	}
	if err := lm.client.Expire(key, lm.period); err != nil {
		return nil, err
	}
	// 返回所需令牌数和剩余令牌数
	return []int64{int64(quantity), int64(tokens)}, nil
}

// SetLimiterCapacity 设置桶容量
func SetLimiterCapacity(capacity float64) LimiterOption {
	return limiterOption(func(limiter *Limiter) {
		limiter.capacity = capacity
	})
}

// SetLimiterQuantity 设置每次需要取得的令牌数
func SetLimiterQuantity(quantity uint64) LimiterOption {
	return limiterOption(func(limiter *Limiter) {
		limiter.quantity = quantity
	})
}

func floatToString(src float64) string {
	return strconv.FormatFloat(src, 'f', -1, 64)
}

func int64ToString(src int64) string {
	return strconv.FormatInt(src, 10)
}

func stringToFloat(src string) float64 {
	if src == "" {
		return 0
	}
	res, err := strconv.ParseFloat(src, 64)
	if err != nil {
		return 0
	}
	return res
}

func stringToInt64(src string) int64 {
	if src == "" {
		return 0
	}
	res, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		return 0
	}
	return res
}
