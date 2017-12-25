package utils

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 通用文件保存
func CommonSaveFile(fileBuffer io.Reader, root, subPath, fileName string) (uri string, err error) {
	dir := root + subPath
	if d, err := os.Stat(dir); err != nil || !d.IsDir() {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	fileByte, err := ioutil.ReadAll(fileBuffer)
	if err != nil {
		return "", err
	}
	filePath := dir + fileName
	err = SaveByte(filePath, fileByte)
	if err != nil {
		return
	}
	uri = strings.TrimPrefix(filePath, root)
	return
}
func SaveByte(path string, fileByte []byte) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write(fileByte)
	if err != nil {
		return
	}
	return
}
