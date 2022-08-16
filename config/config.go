package config

import "github.com/spf13/viper"

// InitConfig reads config file
func InitConfig()  {

	viper.SetConfigFile("../config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
