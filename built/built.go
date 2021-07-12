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
// time   : 2020-11-19 10:37
// version: 1.0.0
// desc   : 构建信息

package built

import (
	"fmt"
	"time"

	"github.com/yhyzgn/goat/date"
)

// 构建信息
var (
	Name    = "goat"  // 名称
	Version = "1.0.0" // 版本
	BuiltAt = ""      // 构建时间
)

// 启动信息
var (
	FullName            string    // 全名
	FullNameWithBuildAt string    // 全名和构建时间
	StartedAt           time.Time // 启动时间
)

// 项目启动时，生成系统信息
func init() {
	FullName = fmt.Sprintf("%s-%s", Name, Version)
	FullNameWithBuildAt = fmt.Sprintf("%s (%s)", FullName, BuiltAt)
	StartedAt = date.Now()
}
