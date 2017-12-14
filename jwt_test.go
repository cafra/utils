package utils

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
)

var secret = []byte("123456")

func TestGetJwtToken(t *testing.T) {
	t.Log(GetJwtToken(jwt.MapClaims{"uid": "123456"}, secret))
}

func BenchmarkGetJwtToken(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := GetJwtToken(jwt.MapClaims{"uid": "1111111"}, secret); err != nil {
			b.Error(err)
		}
	}
}

func TestCheckJwtToken(t *testing.T) {
	claims, err := CheckJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxMjM0NTYifQ.Dk0VoNVjTYygBjWwVw1TTRsDyH6vWbPUfuxSySgnzFk", secret)
	t.Log(err, claims,claims["uid"])
}

func BenchmarkCheckJwtToken(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := CheckJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxMjM0NTYifQ.Dk0VoNVjTYygBjWwVw1TTRsDyH6vWbPUfuxSySgnzFk", secret);err!=nil{
			b.Error(err)
		}
	}
}
