package libmiddleware

import "github.com/spf13/viper"

var internalPrivateKey string
var internalPublicKey string

func Init() {
	internalPrivateKey = viper.GetString("internal.private.key")
	internalPublicKey = viper.GetString("internal.public.key")
	if internalPrivateKey == "" || internalPublicKey == "" {
		panic("Internal keys are not set in the configuration")
	}
}
