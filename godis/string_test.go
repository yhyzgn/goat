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
// time   : 2020-11-21 19:54
// version: 1.0.0
// desc   :

package godis

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	//tst := "abc"
	//tp := reflect.TypeOf(tst)
	//fmt.Println(tp)
	//fmt.Println(tp.Kind())
	//fmt.Println(tp.Kind() == reflect.String)
	//
	//tp = reflect.TypeOf(&tst)
	//fmt.Println(tp)
	//fmt.Println(tp.Elem())
	//fmt.Println(tp.Elem().Kind())

	//tst := []string{"abc"}
	//tp := reflect.ValueOf(tst)
	//fmt.Println(tp.Kind())
	//fmt.Println(tp.Kind() == reflect.Slice)

	//fmt.Println(tp.Type().Elem())
	//fmt.Println(tp.Type().Elem().Kind())
	//fmt.Println(tp.Type().Elem().Kind() == reflect.String)
	//
	//fmt.Println(reflect.SliceOf(tp.Type().Elem()))

	tst := []string{"abc"}
	tp := reflect.TypeOf(tst)
	//fmt.Println(tp.Kind())
	//fmt.Println(tp.Kind() == reflect.Slice)
	//
	//fmt.Println(tp.Elem())
	//fmt.Println(tp.Elem().Kind())
	//fmt.Println(tp.Elem().Kind() == reflect.String)

	vl := reflect.MakeSlice(reflect.SliceOf(tp.Elem()), 0, 0)
	fmt.Println(vl)
	fmt.Println(vl.Type())
}
