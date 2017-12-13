package utils

import (
	"encoding/base64"
	"testing"
)
//
//var privateKey = []byte(`
//-----BEGIN RSA PRIVATE KEY-----
//MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
//7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
//Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
//AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
//ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
//XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
///jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
//IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
//4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
//DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
//9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
//DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
//AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
//-----END RSA PRIVATE KEY-----
//`)
//
//var publicKey = []byte(`
//-----BEGIN PUBLIC KEY-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
//ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
//wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
//AUeJ6PeW+DAkmJWF6QIDAQAB
//-----END PUBLIC KEY-----
//`)
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAu9n07e4eE7Biacx9R0LEiQ3rZgcwz9LJjyKol7EdhlVHJ98J
qhYXfotJdeXGLQvxGADmsbDY5ldr0Wk1UsjNwxgTCXIXSabXKc7rOBlX1PaiPiAF
3qjc5rv660s0K9oWckZ+08i4NNmQFsZ6a0HgsyfSLG2hfegWh77Orgcz2BcnRoBL
9MJCEuFENqf+xdueVTfZz2rCM7dugMcZTyoWFcP74c76yw4vHDQAPjpbmrh0WuMI
M5cJ4W+u/856kEaD60mxyPEL+1W2NJQaJ9usSNDYvXKr3Vonu0I0BQt0LfuYHzWK
3iBwq3SzcamR/60heziJgZ0bJBofQyHj+Cn9lwIDAQABAoIBACqS4izOa6igsB00
SCxPWIWLTw9nj8t6BU5YV4dRj9RzHVZO+TzAFwEKBlMfCUQKUdDT23ToFLBXncrv
IjOp6OBPY3kfj2GU22zRRYQIUlykrO6RiWMGOFJexiZve9p4ad/qVDIhaoYnzL0s
rHAElS1lV//TtOb5I6oON38/iKNvcEKJlhxELvqzLDOnGnAKzM3l93J5dmrkpmOZ
8rZVyeXNtrHisrrJ0AoDmj7MXr4mmmiNs7ZJb2yv55B+fq0JyR/icuI0jfcgtd0V
Boh2gSULZObfgD0j9IMYlwkiR4DmkeOuldewNcz4l44tCy9m+FS7nBa6kYo1J/7P
4zWg8zECgYEAzCmi8ODfdKZ7q4am/1zqI9dEfA6dGEIQQyREGlS65blzOlFLq+hG
2QMtzL4MenRUpP/QRLzUnkiJqupds2iWrrfte6Kec3uo4rItrSnnaXcLWH7dNOeD
WBSeeKcIfkBHpPRDkuiC8okbhf8Ng4sqRT5+oUM73N/7uOjRkAthNV8CgYEA64wb
ELEEgtAIeWzZTXekILv8RWREYmSYjHrBG4zKSsouhIXjl7t0eHxXfzWtLBBkLfWJ
X8pvhawt74g2EhYMAYBx7q4CwfBDcrNisTgfoep19a7d3pHM7IUHSEXXhBoIFJDi
q+24bGmyZaPBUEUlJ7MeTrh3ILX/2tk9E0haqskCgYBKEITC++E0sTzGIggtNajf
LbXzh12oMjcyFFL8dmaC9j7+FgXsrEwfaA7SatOeDNu0K/WDKjm73jbLIVCyyCt5
4NGve3QeEutWqir12fDQitY72XIoQiCc8IX44SesnWcgSVjGT8FJeUHZ34gog3Dn
Q9+uYvSxkTQBhbyYk/hE4wKBgHJeQ+H14XfWpNa4aEZ5+gI+5H2Y8q9Hot5K2CqV
UL/BrZaBIAHTbfj2ftFwcZX8m3fJSZtuQnoIIQG2BHMBq3CrOiam7QXXsBgoS5o6
4vkOS5ov/uCLsJGDAgcwijVFInlB5B2QvkQ9ifZZ7YoZGLJPAT89x/HlDMbpRgNv
1T4pAoGBAIvPw0EyjjmzyQrGGoZTcrpS5u4qmM0dKGkBF21QFnQlnrgRNPeMgwLe
h78Pfag1//hY3qm3hE9t23CUlyxxrsBmri6yYUMT2kspIhoAGCeIHBqF35dBmj8Q
ypHxVH9XYbZRUgR7LWtpXOD+9ySUG3VoK0OJTpUN+sLDfVl5T59g
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu9n07e4eE7Biacx9R0LE
iQ3rZgcwz9LJjyKol7EdhlVHJ98JqhYXfotJdeXGLQvxGADmsbDY5ldr0Wk1UsjN
wxgTCXIXSabXKc7rOBlX1PaiPiAF3qjc5rv660s0K9oWckZ+08i4NNmQFsZ6a0Hg
syfSLG2hfegWh77Orgcz2BcnRoBL9MJCEuFENqf+xdueVTfZz2rCM7dugMcZTyoW
FcP74c76yw4vHDQAPjpbmrh0WuMIM5cJ4W+u/856kEaD60mxyPEL+1W2NJQaJ9us
SNDYvXKr3Vonu0I0BQt0LfuYHzWK3iBwq3SzcamR/60heziJgZ0bJBofQyHj+Cn9
lwIDAQAB
-----END PUBLIC KEY-----
`)


func TestRsa(t *testing.T) {
	data, err := RsaEncrypt(S2B("2#95588#13683515842#陈贞#123456#"), publicKey)
	if err != nil {
		panic(err)
	}
	//t.Log(B2S(data))
	tt:= base64.StdEncoding.EncodeToString(data)
	t.Log(tt)
	tt2,err:=base64.StdEncoding.DecodeString(tt)
	origData, err := RsaDecrypt(tt2, privateKey)
	if err != nil {
		panic(err)
	}
	t.Log(B2S(origData))
//	1#95588#13683515838#陈贞#123456#411303199108182816
}

func BenchmarkRas(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		data, err := RsaEncrypt(S2B("test123"), publicKey)
		if err != nil {
			panic(err)
		}
		origData, err := RsaDecrypt(data, privateKey)
		if err != nil {
			panic(err)
		}
		b.Log(B2S(origData))
	}
}
