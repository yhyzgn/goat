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
// time   : 2020-11-19 10:53
// version: 1.0.0
// desc   : CONNECT 方法

package client

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/yhyzgn/gog"
)

// Connect 协议
func Connect(url string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodConnect, url, nil)
	if err != nil {
		gog.Error(err)
		return false
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		gog.Error(err)
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			gog.Error(err)
		}
	}(res.Body)

	return res.StatusCode == http.StatusOK
}
