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
// time   : 2020-11-19 10:58
// version: 1.0.0
// desc   :

package client

import (
	"testing"

	"github.com/yhyzgn/goat/json"
	"github.com/yhyzgn/gog"
)

func TestPostForm(t *testing.T) {
	a := make(map[string]interface{})
	//a["tt"] = 234
	_, ok := a["tt"]
	gog.Info(ok)

	bs, err := PostForm("https://t-suite-login.ybsjyyn.com/api/health", map[string]interface{}{
		"id":   23,
		"Name": "李四",
		"age":  44,
	})
	if err != nil {
		return
	}

	std := &Student{}
	json.Decode(string(bs), std)
	gog.Info(std)
}

func TestPostJSON(t *testing.T) {
	bs, err := PostJSON("https://t-suite-login.ybsjyyn.com/api/health", map[string]interface{}{
		"name": "李四",
		"age":  44,
	})
	if err != nil {
		return
	}

	user := &User{}
	json.Decode(string(bs), user)
	gog.Info(user)
}

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Token string `json:"token"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
