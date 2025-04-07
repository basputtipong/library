package libenv

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func InitEnv() {
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Cannot read config.yaml with error: ", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("Configfile ConfigFileNotFoundError: ", err)
		} else {
			log.Println("Configfile format : ", err)
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, key := range viper.AllKeys() {
		val := viper.Get(key)
		viper.Set(key, val)
	}
}
