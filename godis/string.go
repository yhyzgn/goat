// Copyright 2020 yhyzgn goat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-11-21 19:43
// version: 1.0.0
// desc   : string

package godis

import (
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"github.com/yhyzgn/goat/date"
	"github.com/yhyzgn/goat/rand"
)

type str struct {
	Val  tpVal
	List tpList
	Set  tpSet
	ZSet tpZSet
	Hash tpHash
}

func newString(pool *RedisPool) *str {
	return &str{
		Val:  stringVal{pool: pool},
		List: stringList{pool: pool},
		Set:  stringSet{pool: pool},
		ZSet: stringZSet{pool: pool},
		Hash: stringHash{pool: pool},
	}
}

//
// --
/************************************************************************** String Value **************************************************************************/
//

// String Value
type stringVal struct {
	pool *RedisPool
}

// Set ...
func (sv stringVal) Set(key Key, value interface{}) error {
	return sv.SetWithExpiry(key, value, 0)
}

// SetNX ...
func (sv stringVal) SetNX(key Key, value interface{}) error {
	return sv.SetNXWithExpiry(key, value, 0)
}

// SetWithExpiry ...
func (sv stringVal) SetWithExpiry(key Key, value interface{}, expiry time.Duration) error {
	if kind := reflect.TypeOf(value).Kind(); kind != reflect.String {
		panic("Unexpected type of value. Expects 'string', but received '" + kind.String() + "'.")
	}
	_, err := sv.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		if expiry > 0 {
			_, err := conn.Do("SET", realKey, value, "PX", expiry.Milliseconds())
			return nil, err
		}
		_, err := conn.Do("SET", realKey, value)
		return nil, err

	})
	return err
}

// SetNXWithExpiry ...
func (sv stringVal) SetNXWithExpiry(key Key, value interface{}, expiry time.Duration) error {
	if kind := reflect.TypeOf(value).Kind(); kind != reflect.String {
		panic("Unexpected type of value. Expects 'string', but received '" + kind.String() + "'.")
	}
	_, err := sv.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		if expiry > 0 {
			_, err := conn.Do("SETNX", realKey, value, "PX", expiry.Milliseconds())
			return nil, err
		}
		_, err := conn.Do("SETNX", realKey, value)
		return nil, err

	})
	return err
}

// Get ...
func (sv stringVal) Get(key Key, value interface{}) error {
	res, err := sv.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.String(conn.Do("GET", realKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	if err == nil && res != nil {
		err = setValue(res, value)
	}
	return err
}

// GetSet ...
func (sv stringVal) GetSet(key Key, value interface{}) error {
	res, err := sv.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.String(conn.Do("GETSET", realKey, value))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	if err == nil && res != nil {
		err = setValue(res, value)
	}
	return err
}

// Len ...
func (sv stringVal) Len(key Key) int {
	res, err := sv.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Int(conn.Do("STRLEN", realKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	if err != nil {
		return 0
	}
	return res.(int)
}

//
// --
/************************************************************************** String List **************************************************************************/
//

// String List
type stringList struct {
	pool *RedisPool
}

// RPush ...
func (sl stringList) RPush(key Key, value ...interface{}) error {
	return sl.RPushWithExpiry(0, key, value...)
}

// RPushWithExpiry ...
func (sl stringList) RPushWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	return sl.push("RPUSH", expiry, key, value...)
}

// RPushX ...
func (sl stringList) RPushX(key Key, value ...interface{}) error {
	return sl.RPushXWithExpiry(0, key, value...)
}

// RPushXWithExpiry ...
func (sl stringList) RPushXWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	return sl.push("RPUSHX", expiry, key, value...)
}

// RPop ...
func (sl stringList) RPop(key Key, value interface{}) error {
	return sl.pop("RPOP", key, value)
}

// BRPop ...
func (sl stringList) BRPop(key Key, value interface{}) error {
	return sl.pop("BRPOP", key, value)
}

// LPush ...
func (sl stringList) LPush(key Key, value ...interface{}) error {
	return sl.LPushWithExpiry(0, key, value...)
}

// LPushWithExpiry ...
func (sl stringList) LPushWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	sort.SliceStable(value, func(i, j int) bool { return true })
	return sl.push("LPUSH", expiry, key, value...)
}

// LPushX ...
func (sl stringList) LPushX(key Key, value ...interface{}) error {
	return sl.LPushXWithExpiry(0, key, value...)
}

// LPushXWithExpiry ...
func (sl stringList) LPushXWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	sort.SliceStable(value, func(i, j int) bool { return true })
	return sl.push("LPUSHX", expiry, key, value...)
}

// LRange ...
func (sl stringList) LRange(key Key, value interface{}, start, stop int) error {
	res, err := sl.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		// 全部取出
		values, err := redis.Values(conn.Do("LRANGE", realKey, start, stop))
		if err != nil {
			return nil, err
		}
		// 构造返回数据
		data := make([]string, 0)
		for _, item := range values {
			switch item.(type) {
			case string:
				data = append(data, item.(string))
			case []uint8:
				var bs []byte
				for _, b := range item.([]uint8) {
					bs = append(bs, b)
				}
				data = append(data, string(bs))
			}
		}
		return data, nil
	})
	if err == nil && res != nil {
		err = setValue(res, value)
	}
	return err
}

// LPop ...
func (sl stringList) LPop(key Key, value interface{}) error {
	return sl.pop("LPOP", key, value)
}

// BLPop ...
func (sl stringList) BLPop(key Key, value interface{}) error {
	return sl.pop("BLPOP", key, value)
}

// Set ...
func (sl stringList) Set(key Key, index int, value interface{}) error {
	if tp := reflect.TypeOf(value); !(tp.Kind() == reflect.String || tp.Kind() == reflect.Ptr && tp.Elem().Kind() == reflect.String) {
		panic("Unexpected type of value. Expects 'string' or '*string', but received '" + tp.String() + "'.")
	}
	_, err := sl.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Int(conn.Do("LSET", realKey, index, value))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	return err
}

// Rem ...
func (sl stringList) Rem(key Key, count int, value interface{}) error {
	if tp := reflect.TypeOf(value); !(tp.Kind() == reflect.String || tp.Kind() == reflect.Ptr && tp.Elem().Kind() == reflect.String) {
		panic("Unexpected type of value. Expects 'string' or '*string', but received '" + tp.String() + "'.")
	}
	_, err := sl.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Int(conn.Do("LREM", realKey, count, value))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	return err
}

// Del ...
func (sl stringList) Del(key Key, index int) error {
	_, err := sl.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		delValue := "del-goat-godis-list-" + rand.String(12) + "-" + date.NowUnixStr()
		if err := conn.Send("LSET", realKey, index, delValue); err != nil {
			return nil, err
		}
		if err := conn.Send("LREM", realKey, 0, delValue); err != nil {
			return nil, err
		}
		return nil, conn.Flush()
	})
	return err
}

// All ...
func (sl stringList) All(key Key, value interface{}) error {
	return sl.LRange(key, value, 0, -1)
}

// Len ...
func (sl stringList) Len(key Key) int {
	res, err := sl.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Int(conn.Do("LLen", realKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	if err == nil && res != nil {
		return res.(int)
	}
	// 如果列表 key 不存在，则 key 被解释为一个空列表，返回 0
	return 0
}

// Index ...
func (sl stringList) Index(key Key, index int, value interface{}) error {
	res, err := sl.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.String(conn.Do("LINDEX", realKey, index))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] or index [%d] not exists", realKey, index)
		}
		return value, err
	})
	if err == nil && res != nil {
		err = setValue(res, value)
	}
	return err
}

func (sl stringList) push(cmd string, expiry time.Duration, key Key, value ...interface{}) error {
	if value == nil {
		return errors.New("argument value of " + cmd + " can not be nil")
	}
	_, err := sl.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		args := append([]interface{}{realKey}, value...)
		if err := conn.Send(cmd, args...); err != nil {
			return nil, err
		}
		// 超时时间大于0时才起效
		if expiry > 0 {
			if err := conn.Send("PEXPIRE", realKey, expiry.Milliseconds()); err != nil {
				return nil, err
			}
		}
		return nil, conn.Flush()
	})
	return err
}

func (sl stringList) pop(cmd string, key Key, value interface{}) error {
	res, err := sl.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.String(conn.Do(cmd, realKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	if err == nil && res != nil {
		err = setValue(res, value)
	}
	return err
}

//
// --
/************************************************************************** String Set **************************************************************************/
//

// String Set
type stringSet struct {
	pool *RedisPool
}

// Add ...
func (ss stringSet) Add(key Key, value ...interface{}) error {
	return ss.AddWithExpiry(0, key, value...)
}

// AddWithExpiry ...
func (ss stringSet) AddWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	if value == nil {
		return errors.New("argument value of Add can not be nil")
	}
	_, err := ss.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		args := append([]interface{}{realKey}, value...)
		if err := conn.Send("SADD", args...); err != nil {
			return nil, err
		}
		// 超时时间大于0时才起效
		if expiry > 0 {
			if err := conn.Send("PEXPIRE", realKey, expiry.Milliseconds()); err != nil {
				return nil, err
			}
		}
		return nil, conn.Flush()
	})
	return err
}

// All ...
func (ss stringSet) All(key Key, value interface{}) error {
	if tp := reflect.TypeOf(value); tp.Kind() != reflect.Ptr || tp.Elem().Kind() != reflect.Slice || tp.Elem().Elem().Kind() != reflect.String {
		panic("Unexpected type of value. Expects '[]string', but received '" + tp.String() + "'.")
	}

	res, err := ss.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		// 全部取出
		values, err := redis.Values(conn.Do("SMEMBERS", realKey))
		if err != nil {
			return nil, err
		}
		// 构造返回数据
		data := make([]string, 0)
		for _, item := range values {
			switch item.(type) {
			case string:
				data = append(data, item.(string))
			case []uint8:
				var bs []byte
				for _, b := range item.([]uint8) {
					bs = append(bs, b)
				}
				data = append(data, string(bs))
			}
		}
		return data, nil
	})
	if err == nil && res != nil {
		err = setValue(res, value)
	}
	return err
}

// Len ...
func (ss stringSet) Len(key Key) int {
	res, err := ss.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Int(conn.Do("SCARD", realKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	if err == nil && res != nil {
		return res.(int)
	}
	// 如果列表 key 不存在，返回 0
	return 0
}

// Has ...
func (ss stringSet) Has(key Key, value interface{}) bool {
	res, err := ss.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Bool(conn.Do("SISMEMBER", realKey, value))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] or value not exists", realKey)
		}
		return value, err
	})
	if err == nil && res != nil {
		return res.(bool)
	}
	return false
}

// Rem ...
func (ss stringSet) Rem(key Key, value ...interface{}) error {
	if value == nil {
		return errors.New("argument value of SREM can not be nil")
	}
	_, err := ss.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		args := []interface{}{realKey, value[:]}
		value, err := redis.Int(conn.Do("SREM", args...))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	return err
}

//
// --
/************************************************************************** String ZSet **************************************************************************/
//

// String ZSet
type stringZSet struct {
	pool *RedisPool
}

// Add ...
func (sz stringZSet) Add(key Key, score int, value interface{}) error {
	return sz.AddWithExpiry(key, score, value, 0)
}

// AddWithExpiry ...
func (sz stringZSet) AddWithExpiry(key Key, score int, value interface{}, expiry time.Duration) error {
	if tp := reflect.TypeOf(value); !(tp.Kind() == reflect.String || tp.Kind() == reflect.Ptr && tp.Elem().Kind() == reflect.String) {
		panic("Unexpected type of value. Expects 'string' or '*string', but received '" + tp.String() + "'.")
	}
	_, err := sz.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		// 超时时间大于0时才起效
		if err := conn.Send("ZADD", realKey, score, value); err != nil {
			return nil, err
		}
		if expiry > 0 {
			if err := conn.Send("PEXPIRE", realKey, expiry.Milliseconds(), 10); err != nil {
				return nil, err
			}
		}
		return nil, conn.Flush()
	})
	return err
}

// Range ...
func (sz stringZSet) Range(key Key, value interface{}, start, stop int) error {
	tp := reflect.TypeOf(value)
	if tp.Kind() != reflect.Slice || tp.Elem().Kind() != reflect.String {
		panic("Unexpected type of value. Expects '[]string', but received '" + tp.String() + "'.")
	}

	res, err := sz.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		// 全部取出
		values, err := redis.Values(conn.Do("ZRANGE", realKey, start, stop))
		if err != nil {
			return nil, err
		}
		// 构造返回数据
		data := make([]string, 0)
		for _, item := range values {
			switch item.(type) {
			case string:
				data = append(data, item.(string))
			case []uint8:
				var bs []byte
				for _, b := range item.([]uint8) {
					bs = append(bs, b)
				}
				data = append(data, string(bs))
			}
		}
		return data, nil
	})
	if err == nil && res != nil {
		err = setValue(res, value)
	}
	return err
}

// All ...
func (sz stringZSet) All(key Key, value interface{}) error {
	return sz.Range(key, value, 0, -1)
}

// Len ...
func (sz stringZSet) Len(key Key) int {
	res, err := sz.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Int(conn.Do("ZCARD", realKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	if err == nil && res != nil {
		return res.(int)
	}
	// 如果列表 key 不存在，返回 0
	return 0
}

// Rem ...
func (sz stringZSet) Rem(key Key, value ...interface{}) error {
	if value == nil {
		return errors.New("argument value of ZREM can not be nil")
	}
	tp := reflect.TypeOf(value)
	if kind := tp.Elem().Kind(); kind != reflect.String {
		panic("Unexpected type of value. Expects '[]string', but received '" + tp.String() + "'.")
	}
	_, err := sz.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		args := []interface{}{realKey, value[:]}
		value, err := redis.Int(conn.Do("ZREM", args...))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	return err
}

//
// --
/************************************************************************** String Hash **************************************************************************/
//

// String Hash
type stringHash struct {
	pool *RedisPool
}

// Set ...
func (sh stringHash) Set(key, hashKey Key, value interface{}) error {
	if tp := reflect.TypeOf(value); !(tp.Kind() == reflect.String || tp.Kind() == reflect.Ptr && tp.Elem().Kind() == reflect.String) {
		panic("Unexpected type of value. Expects 'string' or '*string', but received '" + tp.String() + "'.")
	}
	_, err := sh.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		// 超时时间大于0时才起效
		if err := conn.Send("HSET", realKey, hashKey, value); err != nil {
			return nil, err
		}
		return nil, conn.Flush()
	})
	return err
}

// SetNX ...
func (sh stringHash) SetNX(key, hashKey Key, value interface{}) error {
	if tp := reflect.TypeOf(value); !(tp.Kind() == reflect.String || tp.Kind() == reflect.Ptr && tp.Elem().Kind() == reflect.String) {
		panic("Unexpected type of value. Expects 'string' or '*string', but received '" + tp.String() + "'.")
	}
	_, err := sh.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		// 超时时间大于0时才起效
		if err := conn.Send("HSETNX", realKey, hashKey, value); err != nil {
			return nil, err
		}
		return nil, conn.Flush()
	})
	return err
}

// Get ...
func (sh stringHash) Get(key, hashKey Key, value interface{}) error {
	res, err := sh.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.String(conn.Do("HGET", realKey, hashKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	if err == nil && res != nil {
		err = setValue(res, value)
	}
	return err
}

// Exists ...
func (sh stringHash) Exists(key, hashKey Key) bool {
	res, err := sh.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Bool(conn.Do("HEXISTS", realKey, hashKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	return err == nil && res.(bool)
}

// Del ...
func (sh stringHash) Del(key Key, hashKey ...Key) error {
	_, err := sh.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Int(conn.Do("HDEL", realKey, hashKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	return err
}

// Len ...
func (sh stringHash) Len(key Key) int {
	res, err := sh.pool.call(key, func(conn redis.Conn, realKey string) (interface{}, error) {
		value, err := redis.Int(conn.Do("HLEN", realKey))
		if err == redis.ErrNil {
			err = fmt.Errorf("the key [%s] not exists", realKey)
		}
		return value, err
	})
	if err != nil {
		return 0
	}
	return res.(int)
}
