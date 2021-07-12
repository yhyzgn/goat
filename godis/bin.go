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
// time   : 2020-11-21 20:20
// version: 1.0.0
// desc   : binary

package godis

import (
	"reflect"
	"time"

	"github.com/pkg/errors"
)

type bin struct {
	Val  tpVal
	List tpList
	Set  tpSet
	ZSet tpZSet
	Hash tpHash
}

func newBin(pool *RedisPool) *bin {
	return &bin{
		Val:  binVal{strVal: stringVal{pool: pool}},
		List: binList{strList: stringList{pool: pool}},
		Set:  binSet{strSet: stringSet{pool: pool}},
		ZSet: binZSet{strZSet: stringZSet{pool: pool}},
		Hash: binHash{strHash: stringHash{pool: pool}},
	}
}

//
// --
/************************************************************************** Bin Value **************************************************************************/
//

// Bin Value
type binVal struct {
	strVal stringVal
}

// Set ...
func (bv binVal) Set(key Key, value interface{}) error {
	gb, err := toGob(value)
	if err != nil {
		return err
	}
	return bv.strVal.Set(key, gb)
}

// SetNX ...
func (bv binVal) SetNX(key Key, value interface{}) error {
	gb, err := toGob(value)
	if err != nil {
		return err
	}
	return bv.strVal.SetNX(key, gb)
}

// SetWithExpiry ...
func (bv binVal) SetWithExpiry(key Key, value interface{}, expiry time.Duration) error {
	gb, err := toGob(value)
	if err != nil {
		return err
	}
	return bv.strVal.SetWithExpiry(key, gb, expiry)
}

// SetNXWithExpiry ...
func (bv binVal) SetNXWithExpiry(key Key, value interface{}, expiry time.Duration) error {
	gb, err := toGob(value)
	if err != nil {
		return err
	}
	return bv.strVal.SetNXWithExpiry(key, gb, expiry)
}

// Get ...
func (bv binVal) Get(key Key, value interface{}) error {
	var res string
	if err := bv.strVal.Get(key, &res); err != nil {
		return err
	}
	if err := fromGob(res, value); err != nil {
		return err
	}
	return nil
}

// GetSet ...
func (bv binVal) GetSet(key Key, value interface{}) error {
	jsn, err := toJSON(value)
	if err != nil {
		return err
	}
	if err := bv.strVal.GetSet(key, &jsn); err != nil {
		return err
	}
	if err := fromGob(jsn, value); err != nil {
		return err
	}
	return nil
}

// Len ...
func (bv binVal) Len(key Key) int {
	// json 类型只判断是否存在
	if exist, err := bv.strVal.pool.Exists(key); err == nil && exist {
		return 1
	}
	return 0
}

//
// --
/************************************************************************** Bin List **************************************************************************/
//

// Bin List
type binList struct {
	strList stringList
}

// RPush ...
func (bl binList) RPush(key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bl.strList.RPush(key, values...)
}

// RPushWithExpiry ...
func (bl binList) RPushWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bl.strList.RPushWithExpiry(expiry, key, values...)
}

// RPushX ...
func (bl binList) RPushX(key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bl.strList.RPushX(key, values...)
}

// RPushXWithExpiry ...
func (bl binList) RPushXWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bl.strList.RPushXWithExpiry(expiry, key, values...)
}

// RPop ...
func (bl binList) RPop(key Key, value interface{}) error {
	var res string
	if err := bl.strList.RPop(key, &res); err != nil {
		return err
	}
	if err := fromGob(res, value); err != nil {
		return err
	}
	return nil
}

// BRPop ...
func (bl binList) BRPop(key Key, value interface{}) error {
	var res string
	if err := bl.strList.BRPop(key, &res); err != nil {
		return err
	}
	if err := fromGob(res, value); err != nil {
		return err
	}
	return nil
}

// LPush ...
func (bl binList) LPush(key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bl.strList.LPush(key, values...)
}

// LPushWithExpiry ...
func (bl binList) LPushWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bl.strList.LPushWithExpiry(expiry, key, values...)
}

// LPushX ...
func (bl binList) LPushX(key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bl.strList.LPushX(key, values...)
}

// LPushXWithExpiry ...
func (bl binList) LPushXWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bl.strList.LPushXWithExpiry(expiry, key, values...)
}

// LRange ...
func (bl binList) LRange(key Key, value interface{}, start, stop int) error {
	var res []string
	if err := bl.strList.LRange(key, &res, start, stop); err != nil {
		return err
	}
	return gobStringToInterfaceSlice(value, res...)
}

// LPop ...
func (bl binList) LPop(key Key, value interface{}) error {
	var res string
	if err := bl.strList.LPop(key, &res); err != nil {
		return err
	}
	if err := fromGob(res, value); err != nil {
		return err
	}
	return nil
}

// BLPop ...
func (bl binList) BLPop(key Key, value interface{}) error {
	var res string
	if err := bl.strList.BLPop(key, &res); err != nil {
		return err
	}
	if err := fromGob(res, value); err != nil {
		return err
	}
	return nil
}

// Set ...
func (bl binList) Set(key Key, index int, value interface{}) error {
	gb, err := toGob(value)
	if err != nil {
		return err
	}
	return bl.strList.Set(key, index, gb)
}

// Rem ...
func (bl binList) Rem(key Key, count int, value interface{}) error {
	gb, err := toGob(value)
	if err != nil {
		return err
	}
	return bl.strList.Rem(key, count, gb)
}

// Del ...
func (bl binList) Del(key Key, index int) error {
	return bl.strList.Del(key, index)
}

// All ...
func (bl binList) All(key Key, value interface{}) error {
	return bl.LRange(key, value, 0, -1)
}

// Len ...
func (bl binList) Len(key Key) int {
	return bl.strList.Len(key)
}

// Index ...
func (bl binList) Index(key Key, index int, value interface{}) error {
	var res string
	if err := bl.strList.Index(key, index, &res); err != nil {
		return err
	}
	if err := fromGob(res, value); err != nil {
		return err
	}
	return nil
}

//
// --
/************************************************************************** Bin Set **************************************************************************/
//

// Bin Set
type binSet struct {
	strSet stringSet
}

// Add ...
func (bs binSet) Add(key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bs.strSet.Add(key, values...)
}

// AddWithExpiry ...
func (bs binSet) AddWithExpiry(expiry time.Duration, key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bs.strSet.AddWithExpiry(expiry, key, values...)
}

// All ...
func (bs binSet) All(key Key, value interface{}) error {
	var res []string
	if err := bs.strSet.All(key, &res); err != nil {
		return err
	}
	return gobStringToInterfaceSlice(value, res...)
}

// Len ...
func (bs binSet) Len(key Key) int {
	return bs.strSet.Len(key)
}

// Has ...
func (bs binSet) Has(key Key, value interface{}) bool {
	gb, err := toJSON(value)
	if err != nil {
		return false
	}
	return bs.strSet.Has(key, gb)
}

// Rem ...
func (bs binSet) Rem(key Key, value ...interface{}) error {
	values, err := jsonInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bs.strSet.Rem(key, values...)
}

//
// --
/************************************************************************** Bin ZSet **************************************************************************/
//

// Bin ZSet
type binZSet struct {
	strZSet stringZSet
}

// Add ...
func (bz binZSet) Add(key Key, score int, value interface{}) error {
	gb, err := toGob(value)
	if err != nil {
		return err
	}
	return bz.strZSet.Add(key, score, gb)
}

// AddWithExpiry ...
func (bz binZSet) AddWithExpiry(key Key, score int, value interface{}, expiry time.Duration) error {
	gb, err := toGob(value)
	if err != nil {
		return err
	}
	return bz.strZSet.AddWithExpiry(key, score, gb, expiry)
}

// Range ...
func (bz binZSet) Range(key Key, value interface{}, start, stop int) error {
	var res []string
	if err := bz.strZSet.Range(key, &res, start, stop); err != nil {
		return err
	}
	return gobStringToInterfaceSlice(value, res...)
}

// All ...
func (bz binZSet) All(key Key, value interface{}) error {
	return bz.Range(key, value, 0, -1)
}

// Len ...
func (bz binZSet) Len(key Key) int {
	return bz.strZSet.Len(key)
}

// Rem ...
func (bz binZSet) Rem(key Key, value ...interface{}) error {
	values, err := gobInterfaceSlice(value...)
	if err != nil {
		return err
	}
	return bz.strZSet.Rem(key, values...)
}

//
// --
/************************************************************************** Bin Hash **************************************************************************/
//

// Bin Hash
type binHash struct {
	strHash stringHash
}

// Set ...
func (bh binHash) Set(key, hashKey Key, value interface{}) error {
	gb, err := toJSON(value)
	if err != nil {
		return err
	}
	return bh.strHash.Set(key, hashKey, gb)
}

// SetNX ...
func (bh binHash) SetNX(key, hashKey Key, value interface{}) error {
	gb, err := toJSON(value)
	if err != nil {
		return err
	}
	return bh.strHash.SetNX(key, hashKey, gb)
}

// Get ...
func (bh binHash) Get(key, hashKey Key, value interface{}) error {
	var res string
	if err := bh.strHash.Get(key, hashKey, &res); err != nil {
		return err
	}
	if err := fromGob(res, value); err != nil {
		return err
	}
	return nil
}

// Exists ...
func (bh binHash) Exists(key, hashKey Key) bool {
	return bh.strHash.Exists(key, hashKey)
}

// Del ...
func (bh binHash) Del(key Key, hashKey ...Key) error {
	return bh.strHash.Del(key, hashKey...)
}

// Len ...
func (bh binHash) Len(key Key) int {
	return bh.strHash.Len(key)
}

// 普通对象列表转换为gob列表
func gobInterfaceSlice(value ...interface{}) ([]interface{}, error) {
	values := make([]interface{}, len(value))
	for idx, val := range value {
		jsn, err := toGob(val)
		if err != nil {
			return nil, err
		}
		values[idx] = jsn
	}
	return values, nil
}

// 将gob字符串数组转换为常规对象数组
func gobStringToInterfaceSlice(value interface{}, values ...string) (err error) {
	tp := reflect.TypeOf(value)
	if tp.Kind() != reflect.Ptr {
		err = errors.New("The argument 'value' must be ptr")
		return
	}
	vo := reflect.ValueOf(value)
	// 取得数组中元素的类型
	elmType := tp.Elem().Elem()
	// 数组的值
	elmValue := vo.Elem()
	// 创建一个新数组并把元素的值追加进去
	newArr := make([]reflect.Value, 0)

	for _, val := range values {
		// new一个数组中的元素对象
		obj := reflect.New(elmType)
		if err = fromGobReflect(val, obj); err != nil {
			return
		}
		newArr = append(newArr, obj.Elem())
	}

	// 把原数组的值和新的数组合并
	resArr := reflect.Append(elmValue, newArr...)

	// 最终结果给原数组
	elmValue.Set(resArr)
	return nil
}
