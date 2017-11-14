package utils

import (
	"bytes"
	"sync"
	"unsafe"
)

func S2B(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

var bufPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

func BufPoolGet() *bytes.Buffer {
	if buf := bufPool.Get(); buf != nil {
		return buf.(*bytes.Buffer)
	} else {
		return &bytes.Buffer{}
	}
}

func put(b *bytes.Buffer) { bufPool.Put(b) }

func BufPoolFree(b *bytes.Buffer) {
	b.Reset()
	put(b)
}
