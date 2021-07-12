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
// time   : 2020-11-20 00:45
// version: 1.0.0
// desc   :

package pbe

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMD5AndDES_Encrypt(t *testing.T) {
	encrypted, err := MD5DES.Encrypt("18810241024", "qqqqqq", "vDDkDzrK", 1000)
	fmt.Println(encrypted, err)

	decrypted, err := MD5DES.Decrypt(encrypted, "qqqqqq", "vDDkDzrK", 1000)
	fmt.Println(decrypted, err)

	fmt.Println("==========================================================")

	tst := MD5AndDES{
		BytesToString: func(bs []byte) string {
			return hex.EncodeToString(bs)
		},
		StringToBytes: func(data string) ([]byte, error) {
			return hex.DecodeString(data)
		},
	}

	encrypted, err = tst.Encrypt("18810241024", "qqqqqq", "vDDkDzrK", 1000)
	fmt.Println(encrypted, err)

	decrypted, err = tst.Decrypt(encrypted, "qqqqqq", "vDDkDzrK", 1000)
	fmt.Println(decrypted, err)
}
