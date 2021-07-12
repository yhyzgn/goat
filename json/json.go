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
// time   : 2020-11-19 14:46
// version: 1.0.0
// desc   : json

package json

import "encoding/json"

// Encode json编码
func Encode(bean interface{}) string {
	bs, err := json.Marshal(bean)
	if err != nil {
		panic(err)
	}
	return string(bs)
}

// Decode json解码
func Decode(jsonStr string, bean interface{}) {
	err := json.Unmarshal([]byte(jsonStr), bean)
	if err != nil {
		panic(err)
	}
}

// Copy json实现深拷贝
func Copy(src, dest interface{}) error {
	bs, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, dest)
}
