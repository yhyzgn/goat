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
// time   : 2020-11-19 10:57
// version: 1.0.0
// desc   :

package client

import (
	"testing"

	"github.com/yhyzgn/gog"
)

func TestGet(t *testing.T) {
	//bs, err := Get("http://localhost:8080/gox/remote/param/param", map[string]interface{}{
	//	"name": "啊哈哈",
	//	"age":  34,
	//})
	bs, err := Get("https://t-suite-login.ybsjyyn.com/api/health", map[string]interface{}{
		"name": "啊哈哈",
		"age":  34,
	})
	if err != nil {
		return
	}

	gog.Trace(string(bs))
}
