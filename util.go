package utils

import "unsafe"

func S2B(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
