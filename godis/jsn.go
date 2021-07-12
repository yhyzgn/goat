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
// time   : 2020-11-21 20:12
// version: 1.0.0
// desc   : json

package godis

import (
	"strings"
	"time"
)

type jsn struct {
	Val  tpVal
	List tpList
	Set  tpSet
	ZSet tpZSet
	Hash tpHash
}

func newJSON(pool *RedisPool) *jsn {
	return &jsn{
		Val:  jsnVal{strVal: stringVal{pool: pool}},
		List: jsnList{strList: stringList{pool: pool}},
		Set:  jsnSet{strSet: stringSet{pool: pool}},
		ZSet: jsnZSet{strZSet: stringZSet{pool: pool}},
		Hash: jsnHash{strHash: stringHash{pool: pool}},
	}
}

//
// --
/************************************************************************** JSON Value **************************************************************************/
//

// JSON Value
type jsnVal struct {
	strVal stringVal
}

// Set ...
func (jv jsnVal) Set(key Key, value interface{}) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jv.strVal.Set(key, jsn)
}

// SetNX ...
func (jv jsnVal) SetNX(key Key, value interface{}) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jv.strVal.SetNX(key, jsn)
}

// SetWithExpiry ...
func (jv jsnVal) SetWithExpiry(key Key, value interface{}, expiry time.Duration) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jv.strVal.SetWithExpiry(key, jsn, expiry)
}

// SetNXWithExpiry ...
func (jv jsnVal) SetNXWithExpiry(key Key, value interface{}, expiry time.Duration) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jv.strVal.SetNXWithExpiry(key, jsn, expiry)
}

// Get ...
func (jv jsnVal) Get(key Key, value interface{}) error {
	var res string
	if err := jv.strVal.Get(key, &res); err != nil {
		return err
	}
	if err := fromJSON(res, value); err != nil {
		return err
	}
	return nil
}

// GetSet ...
func (jv jsnVal) GetSet(key Key, value interface{}) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	if err := jv.strVal.GetSet(key, &jsn); err != nil {
		return err
	}
	if err := fromJSON(jsn, value); err != nil {
		return err
	}
	return nil
}

// Len ...
func (jv jsnVal) Len(key Key) int {
	// json 类型只判断是否存在
	if exist, err := jv.strVal.pool.Exists(key); err == nil && exist {
		return 1
	}
	return 0
}

//
// --
/************************************************************************** JSON List **************************************************************************/
//

// JSON List
type jsnList struct {
	strList stringList
}

// RPush ...
func (jl jsnList) RPush(key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return jl.strList.RPush(key, values...)
}

// RPushWithExpiry ...
func (jl jsnList) RPushWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return jl.strList.RPushWithExpiry(expiry, key, values...)
}

// RPushX ...
func (jl jsnList) RPushX(key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return jl.strList.RPushX(key, values...)
}

// RPushXWithExpiry ...
func (jl jsnList) RPushXWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return jl.strList.RPushXWithExpiry(expiry, key, values...)
}

// RPop ...
func (jl jsnList) RPop(key Key, value interface{}) error {
	var res string
	if err := jl.strList.RPop(key, &res); err != nil {
		return err
	}
	if err := fromJSON(res, value); err != nil {
		return err
	}
	return nil
}

// BRPop ...
func (jl jsnList) BRPop(key Key, value interface{}) error {
	var res string
	if err := jl.strList.BRPop(key, &res); err != nil {
		return err
	}
	if err := fromJSON(res, value); err != nil {
		return err
	}
	return nil
}

// LPush ...
func (jl jsnList) LPush(key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return jl.strList.LPush(key, values...)
}

// LPushWithExpiry ...
func (jl jsnList) LPushWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return jl.strList.LPushWithExpiry(expiry, key, values...)
}

// LPushX ...
func (jl jsnList) LPushX(key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return jl.strList.LPushX(key, values...)
}

// LPushXWithExpiry ...
func (jl jsnList) LPushXWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return jl.strList.LPushXWithExpiry(expiry, key, values...)
}

// LRange ...
func (jl jsnList) LRange(key Key, value interface{}, start, stop int) error {
	var res []string
	if err := jl.strList.LRange(key, &res, start, stop); err != nil {
		return err
	}
	val := "[" + strings.Join(res, ",") + "]"
	if err := fromJSON(val, value); err != nil {
		return err
	}
	return nil
}

// LPop ...
func (jl jsnList) LPop(key Key, value interface{}) error {
	var res string
	if err := jl.strList.LPop(key, &res); err != nil {
		return err
	}
	if err := fromJSON(res, value); err != nil {
		return err
	}
	return nil
}

// BLPop ...
func (jl jsnList) BLPop(key Key, value interface{}) error {
	var res string
	if err := jl.strList.BLPop(key, &res); err != nil {
		return err
	}
	if err := fromJSON(res, value); err != nil {
		return err
	}
	return nil
}

// Set ...
func (jl jsnList) Set(key Key, index int, value interface{}) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jl.strList.Set(key, index, jsn)
}

// Rem ...
func (jl jsnList) Rem(key Key, count int, value interface{}) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jl.strList.Rem(key, count, jsn)
}

// Del ...
func (jl jsnList) Del(key Key, index int) error {
	return jl.strList.Del(key, index)
}

// All ...
func (jl jsnList) All(key Key, value interface{}) error {
	return jl.LRange(key, value, 0, -1)
}

// Len ...
func (jl jsnList) Len(key Key) int {
	return jl.strList.Len(key)
}

// Index ...
func (jl jsnList) Index(key Key, index int, value interface{}) error {
	var res string
	if err := jl.strList.Index(key, index, &res); err != nil {
		return err
	}
	if err := fromJSON(res, value); err != nil {
		return err
	}
	return nil
}

//
// --
/************************************************************************** JSON Set **************************************************************************/
//

// JSON Set
type jsnSet struct {
	strSet stringSet
}

// Add ...
func (js jsnSet) Add(key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return js.strSet.Add(key, values...)
}

// AddWithExpiry ...
func (js jsnSet) AddWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return js.strSet.AddWithExpiry(expiry, key, values...)
}

// All ...
func (js jsnSet) All(key Key, value interface{}) error {
	var res []string
	if err := js.strSet.All(key, &res); err != nil {
		return err
	}
	val := "[" + strings.Join(res, ",") + "]"
	if err := fromJSON(val, value); err != nil {
		return err
	}
	return nil
}

// Len ...
func (js jsnSet) Len(key Key) int {
	return js.strSet.Len(key)
}

// Has ...
func (js jsnSet) Has(key Key, value interface{}) bool {
	jsn, err := toJSON(value)
	if err != nil {
		return false
	}
	return js.strSet.Has(key, jsn)
}

// Rem ...
func (js jsnSet) Rem(key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return js.strSet.Rem(key, values...)
}

//
// --
/************************************************************************** JSON ZSet **************************************************************************/
//

// JSON ZSet
type jsnZSet struct {
	strZSet stringZSet
}

// Add ...
func (jz jsnZSet) Add(key Key, score int, value interface{}) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jz.strZSet.Add(key, score, jsn)
}

// AddWithExpiry ...
func (jz jsnZSet) AddWithExpiry(key Key, score int, value interface{}, expiry time.Duration) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jz.strZSet.AddWithExpiry(key, score, jsn, expiry)
}

// Range ...
func (jz jsnZSet) Range(key Key, value interface{}, start, stop int) error {
	var res []string
	if err := jz.strZSet.Range(key, &res, start, stop); err != nil {
		return err
	}
	val := "[" + strings.Join(res, ",") + "]"
	if err := fromJSON(val, value); err != nil {
		return err
	}
	return nil
}

// All ...
func (jz jsnZSet) All(key Key, value interface{}) error {
	return jz.Range(key, value, 0, -1)
}

// Len ...
func (jz jsnZSet) Len(key Key) int {
	return jz.strZSet.Len(key)
}

// Rem ...
func (jz jsnZSet) Rem(key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return jz.strZSet.Rem(key, values...)
}

//
// --
/************************************************************************** JSON Hash **************************************************************************/
//

// JSON Hash
type jsnHash struct {
	strHash stringHash
}

// Set ...
func (jh jsnHash) Set(key, hashKey Key, value interface{}) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jh.strHash.Set(key, hashKey, jsn)
}

// SetNX ...
func (jh jsnHash) SetNX(key, hashKey Key, value interface{}) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	return jh.strHash.SetNX(key, hashKey, jsn)
}

// Get ...
func (jh jsnHash) Get(key, hashKey Key, value interface{}) error {
	var res string
	if err := jh.strHash.Get(key, hashKey, &res); err != nil {
		return err
	}
	if err := fromJSON(res, value); err != nil {
		return err
	}
	return nil
}

// Exists ...
func (jh jsnHash) Exists(key, hashKey Key) bool {
	return jh.strHash.Exists(key, hashKey)
}

// Del ...
func (jh jsnHash) Del(key Key, hashKey ...Key) error {
	return jh.strHash.Del(key, hashKey...)
}

// Len ...
func (jh jsnHash) Len(key Key) int {
	return jh.strHash.Len(key)
}

// 普通对象列表转换为json列表
func jsonInterfaceSlice(value ...interface{}) ([]interface{}, error) {
	values := make([]interface{}, len(value))
	for idx, val := range value {
		jsn, err := toJSON(val)
		if err != nil {
			return nil, err
		}
		values[idx] = jsn
	}
	return values, nil
}
