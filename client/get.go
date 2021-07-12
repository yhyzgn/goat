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
// time   : 2020-11-19 10:54
// version: 1.0.0
// desc   : GET 方法

package client

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Get 普通 GET 请求
func Get(url string, params map[string]interface{}) ([]byte, error) {
	return GetWithHeader(url, nil, params)
}

// GetWithHeader 带请求头的 GET 请求
func GetWithHeader(url string, headers map[string]interface{}, params map[string]interface{}) ([]byte, error) {
	if url == "" {
		return nil, errors.New("invalid empty url")
	}

	// 处理 Path 请求参数
	pvMap := findPathVariables(url)
	if len(pvMap) > 0 {
		for pv, name := range pvMap {
			value, ok := params[name]
			if !ok {
				panic(fmt.Sprintf("Path参数【%s】缺失", name))
				return nil, nil
			}
			url = strings.ReplaceAll(url, pv, fmt.Sprintf("%v", value))
			delete(params, name)
		}
	}

	var b strings.Builder
	b.WriteString(url)

	if len(params) > 0 {
		separator := "?"
		if strings.HasSuffix(url, "?") {
			// ? 在结尾
			separator = ""
		} else if strings.Contains(url, "?") {
			// ? 在中间
			separator = "&"
		}
		b.WriteString(separator)
		b.WriteString(BuildRawQuery(params))
	}

	if headers == nil {
		headers = map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		}
	}
	bs, err := Request(http.MethodGet, b.String(), headers, nil)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
