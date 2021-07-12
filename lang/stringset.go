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
// time   : 2020-11-19 14:55
// version: 1.0.0
// desc   : string set

package lang

// StringSet 字符串队列
type StringSet map[string]bool

// NewStringSet 新队列
func NewStringSet() *StringSet {
	return &StringSet{}
}

// Add 添加
func (ss *StringSet) Add(item string) bool {
	if ss.Has(item) {
		return false
	}
	(*ss)[item] = true
	return true
}

// Has 是否存在
func (ss *StringSet) Has(item string) bool {
	return (*ss)[item]
}
