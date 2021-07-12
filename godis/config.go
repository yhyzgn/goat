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
// time   : 2020-11-19 11:49
// version: 1.0.0
// desc   : redis 配置

package godis

// Config 配置信息
type Config struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Password  string `yaml:"password"`
	DB        int    `yaml:"db"`
	MaxIdle   int    `yaml:"max-idle"`
	MaxActive int    `yaml:"max-active"`
	KeyPrefix string `yaml:"key-prefix"`
}
