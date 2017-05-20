package authentication

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/jorjuela33/quality-api/config"
)

type JWTAuthentication struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

const (
	tokenDuration = 72
	expire        = 3600
)

var jwtAuthentication *JWTAuthentication = nil

func Shared() *JWTAuthentication {
	if jwtAuthentication == nil {
		jwtAuthentication = &JWTAuthentication{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return jwtAuthentication
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, error := os.Open(config.Current().PrivateKeyPath)
	if error != nil {
		panic(error)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, error = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKey, error := x509.ParsePKCS1PrivateKey(data.Bytes)

	if error != nil {
		panic(error)
	}

	return privateKey
}

func getPublicKey() *rsa.PublicKey {
	publicKeyFile, error := os.Open(config.Current().PublicKeyPath)
	if error != nil {
		panic(error)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, error = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKey, error := x509.ParsePKIXPublicKey(data.Bytes)

	if error != nil {
		panic(error)
	}

	rsaPub, ok := publicKey.(*rsa.PublicKey)

	if !ok {
		panic(error)
	}

	return rsaPub
}
