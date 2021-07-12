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
// time   : 2020-11-19 14:53
// version: 1.0.0
// desc   : lang

package lang

import (
	"strconv"
	"strings"
)

// If 模拟三目运算符
func If(condition bool, positive, negative interface{}) interface{} {
	if condition {
		return positive
	}
	return negative
}

// ParseBool 字符串转布尔
func ParseBool(str string) (bool, error) {
	if str = strings.TrimSpace(str); str == "" {
		return false, nil
	}
	return strconv.ParseBool(str)
}

// BtoI ...
func BtoI(value bool) int {
	return If(value, 1, 0).(int)
}
