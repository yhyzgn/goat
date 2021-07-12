package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetFileNamesInFolder 获取文件夹内各文件的文件名
func GetFileNamesInFolder(path string) (fileNames []string, err error) {
	dirList, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, v := range dirList {
		fileNames = append(fileNames, v.Name())
	}
	return
}

// Directory 当前可执行文件目录
func Directory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

// ExecFilePath 当前可执行文件路径
func ExecFilePath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file, _ = filepath.Abs(file)
	}
	return file
}

// RunningRoot 当前运行根目录
func RunningRoot() string {
	path := filepath.Dir(os.Args[0])
	root, err := filepath.Abs(path)
	if err != nil {
		root = path
	}
	return root
}

// Write 写文件
func Write(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	return writeBytes(file, data)
}

// Append 写文件追加
func Append(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	return writeBytes(file, data)
}

// WriteString 写字符串到文件
func WriteString(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	return writeString(file, data)
}

// AppendString 追加字符串到文件
func AppendString(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	return writeString(file, data)
}

// Read 读文件
func Read(filename string) (data []byte, err error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(file)
	return
}

func writeBytes(file *os.File, data []byte) error {
	_, err := file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func writeString(file *os.File, data string) error {
	_, err := file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

// Copy 复制文件
func Copy(src, dest string) error {
	if !Exists(src) {
		return fmt.Errorf("source file '%s' not found", src)
	}
	// 读取源文件
	bs, err := Read(src)
	if err != nil {
		return err
	}
	// 写入目标文件
	return Write(dest, bs)
}

// Exists 检测文件路径是否存在
//
// 如果返回的错误为nil,说明文件或文件夹存在
// 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
// 如果返回的错误为其它类型,则不确定是否在存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
