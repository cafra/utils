package utils

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

var secret = []byte("1234")

//tps= 156105
func TestGetJwtToken(t *testing.T) {
	t.Log(GetJwtToken(jwt.MapClaims{"uid": 11111111}, secret))
}

func BenchmarkGetJwtToken(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _,err:=GetJwtToken(jwt.MapClaims{"uid": 11111111}, secret);err!=nil{
			b.Error(err)
		}
	}
}

func TestCheckJwtToken(t *testing.T) {
	t.Log(CheckJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjExMTExMTExfQ.EXyv-MogtTUB-aPIHjbzVH6hhAEQLZuc4HlvgkXIbEE", secret))
}
