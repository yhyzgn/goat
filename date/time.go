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
// time   : 2020-11-19 10:46
// version: 1.0.0
// desc   : 时间

package date

import (
	"strconv"
	"time"
)

// Pattern 默认日期格式化模板
const (
	// 日期格式化模板
	Pattern = "2006-01-02 15:04:05"
)

// Format 格式化日期
func Format(date time.Time) string {
	return date.Format(Pattern)
}

// Now 当前时刻，东八区
func Now() time.Time {
	// 设置时区为：东八区
	return time.Now().In(time.FixedZone("CST", 8*3600))
}

// NowMillis 当前时刻时间戳，东八区，单位ms
func NowMillis() int64 {
	return Now().UnixNano() / 1e6
}

// NowUnix 当前时刻时间戳，东八区，单位s
func NowUnix() int64 {
	return Now().Unix()
}

// NowMillisStr 当前时刻时间戳，东八区，单位ms
func NowMillisStr() string {
	return strconv.FormatInt(NowMillis(), 10)
}

// NowUnixStr 当前时刻时间戳，东八区，单位s
func NowUnixStr() string {
	return strconv.FormatInt(NowUnix(), 10)
}
