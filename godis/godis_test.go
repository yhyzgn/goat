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
// time   : 2020-11-22 10:13
// version: 1.0.0
// desc   :

package godis

import (
	"fmt"
	"testing"
	"time"
)

// Instance ...
var (
	Instance *Client
)

func init() {
	cfg := Config{
		Host:      "localhost",
		Port:      6379,
		Password:  "root",
		DB:        5,
		MaxIdle:   20,
		MaxActive: 0,
		KeyPrefix: "goat-test",
	}
	Instance = Init(cfg)
	if err := Instance.Ping(); err != nil {
		fmt.Println(err)
	}
}

func Test(t *testing.T) {
	fmt.Println(Instance)

	jsnVal := Instance.JSON.Val
	fmt.Println(jsnVal.SetWithExpiry("tttt", "123", 3*time.Minute))

	var jVal []string
	fmt.Println(jsnVal.Get("json-val-test", &jVal))
	fmt.Println(jVal)

	jsnSet := Instance.JSON.Set
	fmt.Println(jsnSet.Add("json-set-test", "111111"))
	fmt.Println(jsnSet.Add("json-set-test", "222222"))
	fmt.Println(jsnSet.Add("json-set-test", "222222"))

	var jSet []string
	fmt.Println(jsnSet.All("json-set-test", &jSet))
	fmt.Println(jSet)

	strVal := Instance.String.Val
	fmt.Println(strVal.SetWithExpiry("test", "123", 5*time.Minute))
	//fmt.Println(strVal.SetWithExpiry("test-exp", "123", 3*time.Minute))
	var val string
	fmt.Println(strVal.Get("test", &val))
	fmt.Println(val)
	//fmt.Println(strVal.Len("test"))

	strList := Instance.String.List
	fmt.Println(strList.RPush("string-list", "a", "b"))
	fmt.Println(strList.RPushWithExpiry(3*time.Minute, "string-list", "c", "d"))
	fmt.Println(strList.LPush("string-list", "a", "b"))
	fmt.Println(strList.Len("string-list"))
	//fmt.Println(strList.Del("string-list", 1))
	var lst []string
	fmt.Println(strList.All("string-list", &lst))
	fmt.Println(lst)

	var pop string
	fmt.Println(strList.LPop("string-list", &pop))
	fmt.Println(pop)
}
