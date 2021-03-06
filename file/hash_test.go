package file

import (
	"testing"
)

// 测试文件hash

func Test_MD5(t *testing.T) {
	val, err := MD5("./test.txt")
	if err != nil {
		return
	}
	if val != "098f6bcd4621d373cade4e832627b4f6" {
		t.Errorf("file md5值计算错误:%s", val)
	}
}

func Test_SHA1(t *testing.T) {
	val, err := SHA1("./test.txt")
	if err != nil {
		return
	}
	if val != "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3" {
		t.Errorf("file sha1值计算错误:%s", val)
	}
}

func Test_SHA256(t *testing.T) {
	val, err := SHA256("./test.txt")
	if err != nil {
		return
	}
	if val != "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08" {
		t.Errorf("file sha256值计算错误:%s", val)
	}
}

func Test_SHA512(t *testing.T) {
	val, err := SHA512("./test.txt")
	if err != nil {
		return
	}
	if val != "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff" {
		t.Errorf("file sha512值计算错误:%s", val)
	}
}
