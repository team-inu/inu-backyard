package config

import (
	"log"

	"github.com/spf13/viper"
)

func SetConfig(class interface{}) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&class)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

// TODO: read from struct not viper
func PrintConfig() {
	for _, field := range viper.AllKeys() {
		if viper.IsSet(field) {
			log.Printf("%s: %v", field, viper.Get(field))
		}
	}
}
