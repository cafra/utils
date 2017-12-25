package utils

import (
	"bytes"
	"testing"
)

func TestSaveByte(t *testing.T) {
	t.Log(SaveByte("/Users/cz/Downloads/test.txt", ([]byte)("hello")))
}
func TestCommonSaveFile(t *testing.T) {
	t.Log(CommonSaveFile(bytes.NewBufferString("test3"), "/Users/cz/Downloads/testdir", "/file2/", "test3.txt"))
}
