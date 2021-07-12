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
// time   : 2020-11-19 12:30
// version: 1.0.0
// desc   : 初始化

package godis

// DefaultClient 默认实例
var DefaultClient *Client

// Init 初始化一个默认实例
func Init(config Config) *Client {
	DefaultClient = New(config)
	return DefaultClient
}

// New 创建实例
func New(config Config) *Client {
	pool := NewRedisPool(config.Host, config.Port, config.DB, config.MaxIdle, config.MaxActive, config.Password, config.KeyPrefix)
	return &Client{
		Pool:   pool,
		String: newString(pool),
		JSON:   newJSON(pool),
		Bin:    newBin(pool),
		Subscriber: func(topic ...Key) (*subscriber, error) {
			return newSubscriber(pool, topic...)
		},
		Publisher: func(topic Key) *publisher {
			return newPublisher(pool, topic)
		},
	}
}
