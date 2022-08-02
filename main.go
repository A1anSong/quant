package main

import (
	"fmt"
	"quant/config"
)

func main() {
	config.InitConfig()
	fmt.Println("API Key: " + config.ConfigFile.Binance.ApiKey)
	fmt.Println("Secret Key: " + config.ConfigFile.Binance.SecretKey)
}
