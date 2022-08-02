package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var ConfigFile = ConfigStruct{}

type ConfigStruct struct {
	Binance Binance
}

func InitConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = viper.Unmarshal(&ConfigFile)
	if err != nil {
		fmt.Println(err.Error())
	}
}
