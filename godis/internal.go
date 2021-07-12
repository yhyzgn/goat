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
// time   : 2020-11-19 11:00
// version: 1.0.0
// desc   : golang redis client

package godis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func exists(conn redis.Conn, key string) (bool, error) {
	return redis.Bool(conn.Do("EXISTS", key))
}

func expire(conn redis.Conn, key string, expiry time.Duration) (valid bool, err error) {
	return redis.Bool(conn.Do("PEXPIRE", key, expiry.Milliseconds(), 10))
}

func del(conn redis.Conn, key ...string) error {
	keys := make([]interface{}, len(key))
	for i, k := range key {
		keys[i] = k
	}
	_, err := conn.Do("DEL", keys...)
	return err
}

func bys(conn redis.Conn, key string) ([]byte, error) {
	bs, err := redis.Bytes(conn.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		return nil, fmt.Errorf("[redis][getInt] error in del redis data [name: %s] %s", key, err.Error())
	}
	return bs, nil
}

func write(conn redis.Conn, key string, data []byte, expiry time.Duration) (err error) {
	// 超时时间大于0时才起效
	if expiry > 0 {
		_, err = conn.Do("SET", key, data, "PX", expiry.Milliseconds(), 10)
	} else {
		_, err = conn.Do("SET", key, data)
	}
	return err
}
