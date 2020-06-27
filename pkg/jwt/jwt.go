package jwt

import (
	"encoding/base64"
	"encoding/json"
	rsa "github.com/beego-dev/beemod/pkg/jwt/encrypt"
	"strings"
)

// generate token
func NewToken(jwtHeader *Header, jwtPayload *Payload) (string, error) {
	headerStr, err := HeaderBase64Str(jwtHeader)
	if err != nil {
		return "", err
	}
	payloadStr, err := PayloadBase64Str(jwtPayload)
	if err != nil {
		return "", err
	}
	// generate signature
	signature := Sign(headerStr, payloadStr)
	// splice token
	token := headerStr + "." + payloadStr + "." + signature
	return token, nil
}

// verify token
func VerifyToken(token string) bool {
	tokenSlice := strings.Split(token, ".")
	headerStr := tokenSlice[0]
	payloadStr := tokenSlice[1]
	signatureStr := tokenSlice[2]
	signature, err := base64.StdEncoding.DecodeString(signatureStr)
	if err != nil {
		return false
	}
	plainText := headerStr + payloadStr
	return rsa.VerifyRSA256([]byte(plainText), signature, "public.pem")
}

// generate signature
func Sign(jwtHeaderStr string, jwtPayloadStr string) string {
	// aplice Header and Payload ,then generate signature
	signature := rsa.SignatureRSA256([]byte(jwtHeaderStr+jwtPayloadStr), "private.pem")
	signStr := base64.StdEncoding.EncodeToString(signature)
	return signStr
}

// switch Header to base64 string
func HeaderBase64Str(jwtHeader *Header) (string, error) {
	byteJwtHeader, err := json.Marshal(jwtHeader)
	if err != nil {
		return "", err
	}
	jwtHeaderStr := base64.StdEncoding.EncodeToString(byteJwtHeader)
	return jwtHeaderStr, nil
}

// switch Payload to base64 string
func PayloadBase64Str(jwtPayload *Payload) (string, error) {
	byteJwtPayload, err := json.Marshal(jwtPayload)
	if err != nil {
		return "", err
	}
	jwtPayloadStr := base64.StdEncoding.EncodeToString(byteJwtPayload)
	return jwtPayloadStr, nil
}
