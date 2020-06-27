package main

import (
	"fmt"
	"github.com/beego-dev/beemod/pkg/jwt"
)

func main() {
	testHeader := jwt.Header{
		Algorithm: "RS256",
		Type:      "JWT",
	}
	testPayload := jwt.Payload{
		Issuer:     "guan",
		Subject:    "hello",
		Audience:   "brower",
		Expiration: 111,
		NotBefore:  222,
		IssuedAt:   333,
		JwtId:      "001",
	}
	//fmt.Println(testHeader.Algorithm)
	token, err := jwt.NewToken(&testHeader, &testPayload)
	if err != nil {
		fmt.Println("err：", err)
		return
	}
	fmt.Println("token:", token)
	verify := jwt.VerifyToken(token)
	fmt.Println("verify:", verify)
}
