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
// time   : 2020-11-19 12:52
// version: 1.0.0
// desc   : 弱密码

package password

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
)

const (
	patternNumber    = ".*?\\d.*"      // 数字
	patternCharacter = ".*?[a-zA-Z].*" // 大小写字母
	patternSymbol    = ".*?[\\W_].*"   // 特殊字符
)

// AtLeast 若密码校验
//
// 规则
// ①、[6-20]位
// ②、字母、数字、特殊字符 至少包含${kind}种类型
//
// go语言不支持正则环视，只能分种类校验
func AtLeast(password string, kind int) (err error) {
	if len(password) < 6 || len(password) > 20 {
		err = errors.New("密码长度必须是6-20位")
	} else {
		count := 0
		bs := []byte(password)
		if ok, e := regexp.Match(patternNumber, bs); e == nil && ok {
			count++
		}
		if ok, e := regexp.Match(patternCharacter, bs); e == nil && ok {
			count++
		}
		if ok, e := regexp.Match(patternSymbol, bs); e == nil && ok {
			count++
		}
		if kind < 1 {
			kind = 1
		}
		if kind > 3 {
			kind = 3
		}
		if count < kind {
			err = fmt.Errorf("密码必须包含字母、数字、特殊字符至少%d种", kind)
		}
	}
	return
}
