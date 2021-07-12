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
// time   : 2020-11-22 02:54
// version: 1.0.0
// desc   : gob encoding

package gob

import (
	"bytes"
	"encoding/gob"
	"reflect"
)

// Encode gob编码
func Encode(value interface{}) ([]byte, error) {
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(value)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// Decode gob解码
func Decode(data []byte, value interface{}) error {
	encoded := bytes.NewReader(data)
	decoder := gob.NewDecoder(encoded)
	return decoder.Decode(value)
}

// DecodeValue gob编码，反射赋值
func DecodeValue(data []byte, value reflect.Value) error {
	encoded := bytes.NewReader(data)
	decoder := gob.NewDecoder(encoded)
	return decoder.DecodeValue(value)
}

// Copy gob实现深拷贝
func Copy(src, dest interface{}) error {
	bs, err := Encode(src)
	if err != nil {
		return err
	}
	return Decode(bs, dest)
}
