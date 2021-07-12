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
// time   : 2020-11-19 15:52
// version: 1.0.0
// desc   : 配置信息

package kafka

import "github.com/yhyzgn/gog"

// Config kafka配置
type Config struct {
	Enabled bool                    `yaml:"enabled"`
	Hosts   []string                `yaml:"hosts"`
	Topics  map[KfkTopicName]string `yaml:"topics"`
}

// Init 初始化
func (c Config) Init() {
	if c.Enabled {
		gog.Info("Connecting Kafka server...")
		_, err := InitKafka(c.Hosts, c.Topics)
		if err != nil {
			gog.Error("Kafka server connect failed: ", err.Error())
		} else {
			// 连接成功
			gog.Info("Kafka server connected.")
		}
	} else {
		gog.Warn("Kafka disabled.")
	}
}
