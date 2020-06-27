package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/astaxie/beego/utils"
	"os"
)

// generate privatekey and publickey by RSA
func GenerateRsaKey(keySize int) {
	/***************** privatekey **********************/
	// 1 generate privatekey
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	// 2 marshal privatekey to DER encoded string of ASN.1 by x509-PKCS1
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3 create a pem.Block with privatekey
	block := pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}
	// 4 encode to pem
	file, err := os.Create("private.pem") // make file
	if err != nil {
		fmt.Println("err：", err)
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()

	/***************** publickey **********************/
	// 1 get publickey from privatekey struct
	publicKey := privateKey.PublicKey
	// 2 marshal publickey by x509
	derpText, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// 3 create a pem.Block
	block = pem.Block{
		Type:  "rsa public key",
		Bytes: derpText,
	}
	// 4 encode to pem
	file, err = os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	pem.Encode(file, &block)
	file.Close()
}

// sign by RSA privatekey
func SignatureRSA256(plainText []byte, fileName string) []byte {
	// 1 open privatekey file
	bl := utils.FileExists(fileName)
	if !bl {
		GenerateRsaKey(2048)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	buf := make([]byte, fileInfo.Size())
	// 2 read the privatekey content
	file.Read(buf)
	file.Close()
	// 3 get a blok by decoding with pem
	block, _ := pem.Decode(buf)
	// 4 get privatekey by parsing with x509
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 5 create a hash
	myhash := sha256.New()
	// 6 wirte data
	myhash.Write(plainText)
	// 7 sum the hash
	hashText := myhash.Sum(nil)
	// 8 sign
	signText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashText)
	if err != nil {
		panic(err)
	}
	return signText
}

// verify RSA signature
func VerifyRSA256(plainText, signText []byte, fileName string) bool {
	// 1 open publickey file and read content
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	// 2 get block by decoding pem
	block, _ := pem.Decode(buf)
	// 3 get interface by parsing with x509
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 4 get publickey
	publicKey := pubInterface.(*rsa.PublicKey)
	// 5 get hashText by sum plainText
	hashText := sha256.Sum256(plainText)
	// 6 verify signature
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashText[:], signText)
	if err == nil {
		return true
	}
	return false
}
