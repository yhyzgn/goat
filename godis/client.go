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
// time   : 2020-11-21 23:00
// version: 1.0.0
// desc   : redis 客户端

package godis

import (
	"time"
)

// Client ...
type Client struct {
	Pool       *RedisPool
	String     *str
	JSON       *jsn
	Bin        *bin
	Subscriber func(topic ...Key) (*subscriber, error)
	Publisher  func(topic Key) *publisher
}

// Ping ...
func (c *Client) Ping() (err error) {
	return c.Pool.Ping()
}

// Exists 判断 key 是否存在
func (c *Client) Exists(key Key) (bool, error) {
	return c.Pool.Exists(key)
}

// Expire 设置过期时间
func (c *Client) Expire(key Key, expiry time.Duration) error {
	return c.Pool.Expire(key, expiry)
}

// Keys 按通配符 * 查询一批 key
//
// 查询结果自动去掉已配置的前缀
func (c *Client) Keys(key Key) ([]string, error) {
	return c.Pool.Keys(key)
}

// Del 删除一些 key
func (c *Client) Del(key ...Key) error {
	return c.Pool.Del(key...)
}

// Locker 获取分布式锁
func (c *Client) Locker(key Key, options ...LockerOption) *Locker {
	return c.Pool.Locker(key, options...)
}

// Limiter 获取限流器
func (c *Client) Limiter(key Key, quota uint, period time.Duration, options ...LimiterOption) *Limiter {
	return c.Pool.Limiter(c, key, quota, period, options...)
}

// Close 关闭连接池
func (c *Client) Close() error {
	return c.Pool.Pool.Close()
}
