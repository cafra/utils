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
