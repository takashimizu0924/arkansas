package main

import (
	"arkansas/bybit"
	"arkansas/config"
	"log"
)

func main() {
	config.NewConfig()
	log.Println("[ Start Arkansas !!!!! ]")

	client := bybit.NewBybit(config.Config.ApiKey, config.Config.ApiSecret, true)

	client.CryptTaskV1("BTCUSD")
}
