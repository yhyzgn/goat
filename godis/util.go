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
// time   : 2020-12-06 15:26
// version: 1.0.0
// desc   : 一些工具

package godis

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
	"github.com/yhyzgn/goat/gob"
)

// 反射设置值
func setValue(src, dest interface{}) (err error) {
	if dest == nil {
		err = errors.New("dest can not be nil")
		return
	}
	value := reflect.ValueOf(dest)
	if kind := value.Kind(); kind != reflect.Ptr && kind != reflect.Slice {
		err = errors.New("setValue just support pointer and slice")
		return
	}
	value = value.Elem()
	if !value.CanSet() {
		err = errors.New("dest can not set value")
		return
	}
	value.Set(reflect.ValueOf(src))
	return
}

func toJSON(value interface{}) (string, error) {
	bs, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func fromJSON(jsn string, value interface{}) error {
	return json.Unmarshal([]byte(jsn), value)
}

func toGob(value interface{}) (string, error) {
	bs, err := gob.Encode(value)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func fromGob(gb string, value interface{}) error {
	return gob.Decode([]byte(gb), value)
}

func fromGobReflect(gb string, value reflect.Value) error {
	return gob.DecodeValue([]byte(gb), value)
}
