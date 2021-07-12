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
// time   : 2020-11-20 00:40
// version: 1.0.0
// desc   : PBEWithMD5AndDES

package pbe

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"strings"
)

// MD5AndDES MD5AndDES
type MD5AndDES struct {
	BytesToString func(bs []byte) string            // 加密时 bytes 转为 string 的方法
	StringToBytes func(data string) ([]byte, error) // 解密时 string 转为 bytes 的方法
}

// MD5DES 默认实例
var (
	MD5DES MD5AndDES // 默认实例
)

func init() {
	// 初始化，默认 base64 编码
	MD5DES = MD5AndDES{
		BytesToString: func(bs []byte) string {
			return base64.StdEncoding.EncodeToString(bs)
		},
		StringToBytes: func(data string) ([]byte, error) {
			return base64.StdEncoding.DecodeString(data)
		},
	}
}

func getKey(password string, salt []byte, count int) ([]byte, []byte) {
	key := md5.Sum([]byte(password + string(salt)))
	for i := 0; i < count-1; i++ {
		key = md5.Sum(key[:])
	}
	return key[:8], key[8:]
}

// Encrypt 加密
func (mad MD5AndDES) Encrypt(plaintext, password, salt string, iterations int) (string, error) {
	padNum := byte(8 - len(plaintext)%8)
	for i := byte(0); i < padNum; i++ {
		plaintext += string(padNum)
	}

	dk, iv := getKey(password, []byte(salt), iterations)

	block, err := des.NewCipher(dk)
	if err != nil {
		return "", err
	}

	encryptor := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(plaintext))
	encryptor.CryptBlocks(encrypted, []byte(plaintext))

	if mad.BytesToString == nil {
		panic("Must implement method 'BytesToString' of MD5AndDES.")
	}
	return mad.BytesToString(encrypted), nil
}

// Decrypt 解密
func (mad MD5AndDES) Decrypt(ciphertext, password, salt string, iterations int) (string, error) {
	if mad.StringToBytes == nil {
		panic("Must implement method 'StringToBytes' of MD5AndDES.")
	}
	bs, err := mad.StringToBytes(ciphertext)
	if err != nil {
		return "", err
	}

	dk, iv := getKey(password, []byte(salt), iterations)

	block, err := des.NewCipher(dk)
	if err != nil {
		return "", err
	}

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(bs))
	decrypter.CryptBlocks(decrypted, bs)

	decryptedString := strings.TrimRight(string(decrypted), "\x01\x02\x03\x04\x05\x06\x07\x08")
	return decryptedString, nil
}
