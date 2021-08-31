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
// desc   : POST 方法

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/yhyzgn/gog"
)

// PostForm 普通 POST form-urlencoded 请求
func PostForm(url string, params map[string]interface{}) ([]byte, error) {
	return PostFormWithHeader(url, nil, params)
}

// PostClientForm 普通 POST form-urlencoded 请求
func PostClientForm(client *http.Client, url string, params map[string]interface{}) ([]byte, error) {
	return PostClientFormWithHeader(client, url, nil, params)
}

// PostFormWithHeader 带请求头的 POST form-urlencoded 请求
func PostFormWithHeader(url string, headers map[string]interface{}, params map[string]interface{}) ([]byte, error) {
	return PostClientFormWithHeader(http.DefaultClient, url, headers, params)
}

// PostClientFormWithHeader 带请求头的 POST form-urlencoded 请求
func PostClientFormWithHeader(client *http.Client, url string, headers map[string]interface{}, params map[string]interface{}) ([]byte, error) {
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

	var payload io.Reader
	if len(params) > 0 {
		payload = strings.NewReader(BuildRawQuery(params))
	}

	if headers == nil {
		headers = make(map[string]interface{})
	}
	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	bs, err := RequestClient(client, http.MethodPost, url, headers, payload)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// PostJSON 普通 POST raw 请求
func PostJSON(url string, value interface{}) ([]byte, error) {
	return PostJSONWithHeader(url, nil, value)
}

// PostClientJSON 普通 POST raw 请求
func PostClientJSON(client *http.Client, url string, value interface{}) ([]byte, error) {
	return PostClientJSONWithHeader(client, url, nil, value)
}

// PostJSONWithHeader 带请求头的 POST raw 请求
func PostJSONWithHeader(url string, headers map[string]interface{}, value interface{}) ([]byte, error) {
	return PostClientJSONWithHeader(http.DefaultClient, url, headers, value)
}

// PostClientJSONWithHeader 带请求头的 POST raw 请求
func PostClientJSONWithHeader(client *http.Client, url string, headers map[string]interface{}, value interface{}) ([]byte, error) {
	if url == "" {
		return nil, errors.New("invalid empty url")
	}

	bs, err := json.Marshal(value)
	if err != nil {
		gog.Error(err)
		return nil, nil
	}

	// 这一步只为得到 path 参数
	params := make(map[string]interface{})
	if err = json.Unmarshal(bs, &params); err != nil {
		gog.Error(err)
		return nil, nil
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
		}
	}

	if headers == nil {
		headers = make(map[string]interface{})
	}
	if _, ok := headers["Content-Type"]; !ok {
		headers["Content-Type"] = "application/json"
	}

	bs, err = RequestClient(client, http.MethodPost, url, headers, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	return bs, nil
}
