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
// time   : 2020-11-19 11:05
// version: 1.0.0
// desc   : 随机字符串

package rand

// credits: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

import (
	"math/rand"
	"strconv"

	"github.com/yhyzgn/goat/date"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func init() {
	rand.Seed(date.NowMillis())
}

// String 生成指定长度的随机字符串
func String(length int) string {
	b := make([]byte, length)
	// A src.Int63() generates 63 id bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// Code Generate validation code
func Code(length int) string {
	var code int
	switch length {
	case 4:
		code = rand.Intn(8999) + 1000
	case 5:
		code = rand.Intn(89999) + 10000
	case 6:
		code = rand.Intn(899999) + 100000
	default:
		code = rand.Intn(8999) + 1000
	}
	return strconv.Itoa(code)
}
