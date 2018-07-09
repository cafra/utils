package utils

import (
	"bytes"
	"io/ioutil"
	"os"
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
	b.Log("1")
	buf := BufPoolGet()
	defer BufPoolFree(buf)
	for i := 0; i < b.N; i++ {
		buf.WriteString(`2018-05-25 14:43:34	4	6DE0755D5FDAF08CCF59859606B6D2DF	72503	gaid_122acc07-6e7e-4a6a-84a5-2d5c4bca0638		1		100022	460	zh_CN	Asia%2FShanghai	116.53042	39.930256	wifi	{"name":"Github","package":"com.seasonfif.github","version":"1.7.1","installtime":"1501554396221"}`)
		buf.WriteString("\n")
	}

	ioutil.WriteFile("./tmp.log", buf.Bytes(), os.ModePerm)
}
func BenchmarkBuf2(b *testing.B) {
	//b.SetBytes(10000)
	b.ReportAllocs()
	b.Log("1")
	buf := bytes.NewBufferString("")
	for i := 0; i < b.N; i++ {
		buf.WriteString(`2018-05-25 14:43:34	4	6DE0755D5FDAF08CCF59859606B6D2DF	72503	gaid_122acc07-6e7e-4a6a-84a5-2d5c4bca0638		1		100022	460	zh_CN	Asia%2FShanghai	116.53042	39.930256	wifi	{"name":"Github","package":"com.seasonfif.github","version":"1.7.1","installtime":"1501554396221"}`)
		buf.WriteString("\n")
	}

	ioutil.WriteFile("./tmp.log", buf.Bytes(), os.ModePerm)
}
