package libmiddleware

import (
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var internalPrivateKey string
var internalPublicKey string
var rsaPrivateKey *rsa.PrivateKey

func Init() {
	internalPrivateKey = viper.GetString("internal.private.key")
	internalPublicKey = viper.GetString("internal.public.key")
	if internalPrivateKey == "" || internalPublicKey == "" {
		panic("Internal keys are not set in the configuration")
	}

	privateKeyData := []byte(internalPrivateKey)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		panic(err)
	}

	rsaPrivateKey = privateKey
}
