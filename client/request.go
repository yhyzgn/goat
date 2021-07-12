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
// time   : 2020-11-19 10:49
// version: 1.0.0
// desc   : 请求处理

package client

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/yhyzgn/goat/built"
	"github.com/yhyzgn/gog"
)

var (
	timeout       = 6 * time.Second // 默认超时时间
	logTag        = built.Name + "-client-request-"
	headerTraceID = "Trace-Id"
)

const (
	// 匹配 pathVariable 格式的 pattern
	pathVariable = `{([\w\-_]+)}`
)

// BuildRawQuery GET 参数组装
func BuildRawQuery(params map[string]interface{}) string {
	if params == nil {
		return ""
	}
	values := url.Values{}
	for key, val := range params {
		values[key] = []string{fmt.Sprintf("%v", val)}
	}
	return values.Encode()
}

// Request 请求处理
func Request(method, url string, headers map[string]interface{}, requestBody io.Reader) ([]byte, error) {
	var (
		logReq, logRes []byte
	)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer func() {
		cancel()
	}()

	tID, ok := headers[headerTraceID]
	if !ok {
		tID = base64.URLEncoding.EncodeToString([]byte(url))
	}
	traceID := tID.(string)
	tag := logTag + traceID

	req, err := http.NewRequestWithContext(ctx, method, url, requestBody)
	if err != nil {
		gog.ErrorTag(tag, err)
		return nil, err
	}

	defer func() {
		var sb strings.Builder

		sb.WriteString(fmt.Sprintf("\n================= 【%v】Start =================\n", traceID))
		sb.WriteString("======>> Request <<======\n")
		if nil != logReq {
			sb.Write(logReq)
		}
		sb.WriteString("\n======>> Response <<======\n")
		if nil != logRes {
			sb.Write(logRes)
		}
		sb.WriteString(fmt.Sprintf("\n================= 【%v】End =================\n", traceID))
		gog.InfoTag(tag, sb.String())
	}()

	req.Header.Set("User-Agent", built.Name+"/"+built.Version)
	req.Header.Set("Connection", "Keep-Alive")

	// 自定义 header
	if headers != nil {
		for key, val := range headers {
			req.Header.Set(key, fmt.Sprintf("%v", val))
		}
	}

	// request log
	logReq, _ = httputil.DumpRequest(req, true)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		gog.Error(err)
		return nil, err
	}

	// response log
	logRes, _ = httputil.DumpResponse(res, true)

	return response(res.Body)
}

// 响应处理
func response(responseBody io.ReadCloser) ([]byte, error) {
	defer responseBody.Close()
	bs, err := ioutil.ReadAll(responseBody)
	if err != nil {
		gog.Error(err)
		return nil, err
	}
	return bs, nil
}

// 获取 URL 中的 PathVariable 参数列表
func findPathVariables(url string) map[string]string {
	reg := regexp.MustCompile(pathVariable)
	group := reg.FindAllStringSubmatch(url, -1)
	// url: /api/{id}/book/{code}
	// group: [[{id} id] [{code} code]]
	res := make(map[string]string)
	for _, item := range group {
		res[item[0]] = item[1]
	}
	return res
}
