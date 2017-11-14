package utils

import (
	"testing"
)

func TestS2B(t *testing.T) {
	t.Log(B2S(S2B("1111")))
}

func BenchmarkS2B(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//500000	      2259 ns/op
		b.Log(B2S(S2B("1111")))
		//500000	      2299 ns/op
		//b.Log(string([]byte("1111")))
	}
}
func TestBufPoolGet(t *testing.T) {
	buf := BufPoolGet()
	buf.WriteString("hello")
	buf.String()
	BufPoolFree(buf)
}

func BenchmarkBufPool(b *testing.B) {
	//b.SetBytes(10000)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := BufPoolGet()
		buf.WriteString("hello")
		buf.String()
		BufPoolFree(buf)
	}
}
