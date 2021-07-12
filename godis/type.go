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
// time   : 2020-11-21 19:32
// version: 1.0.0
// desc   : redis 各种数据类型

package godis

import (
	"time"
)

// String
type tpVal interface {
	Set(key Key, value interface{}) error

	SetNX(key Key, value interface{}) error

	SetWithExpiry(key Key, value interface{}, expiry time.Duration) error

	SetNXWithExpiry(key Key, value interface{}, expiry time.Duration) error

	Get(key Key, value interface{}) error

	GetSet(key Key, value interface{}) error

	Len(key Key) int
}

// List
type tpList interface {
	RPush(key Key, value ...interface{}) error

	RPushWithExpiry(expiry time.Duration, key Key, value ...interface{}) error

	RPushX(key Key, value ...interface{}) error

	RPushXWithExpiry(expiry time.Duration, key Key, value ...interface{}) error

	RPop(key Key, value interface{}) error

	BRPop(key Key, value interface{}) error

	LPush(key Key, value ...interface{}) error

	LPushWithExpiry(expiry time.Duration, key Key, value ...interface{}) error

	LPushX(key Key, value ...interface{}) error

	LPushXWithExpiry(expiry time.Duration, key Key, value ...interface{}) error

	LRange(key Key, value interface{}, start, stop int) error

	LPop(key Key, value interface{}) error

	BLPop(key Key, value interface{}) error

	Set(key Key, index int, value interface{}) error

	Rem(key Key, count int, value interface{}) error

	Del(key Key, index int) error

	All(key Key, value interface{}) error

	Len(key Key) int

	Index(key Key, index int, value interface{}) error
}

// Set
type tpSet interface {
	Add(key Key, value ...interface{}) error

	AddWithExpiry(expiry time.Duration, key Key, value ...interface{}) error

	All(key Key, value interface{}) error

	Len(key Key) int

	Has(key Key, value interface{}) bool

	Rem(key Key, value ...interface{}) error
}

// ZSet
type tpZSet interface {
	Add(key Key, score int, value interface{}) error

	AddWithExpiry(key Key, score int, value interface{}, expiry time.Duration) error

	Range(key Key, value interface{}, start, stop int) error

	All(key Key, value interface{}) error

	Len(key Key) int

	Rem(key Key, value ...interface{}) error
}

// Hash
type tpHash interface {
	Set(key, hashKey Key, value interface{}) error

	SetNX(key, hashKey Key, value interface{}) error

	Get(key, hashKey Key, value interface{}) error

	Exists(key, hashKey Key) bool

	Del(key Key, hashKey ...Key) error

	Len(key Key) int
}
