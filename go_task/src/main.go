package main

import (
	"arkansas/bybit"
	"arkansas/config"
	"arkansas/line"
	"log"
	"time"
)

func main() {
	config.NewConfig()
	log.Println("[ Start Arkansas !!!!! ]")

	client := bybit.NewBybit(config.Config.ApiKey, config.Config.ApiSecret, true)
	lineBot, err := line.NewLine()
	if err != nil {
		log.Println(err)
	}
	// 1/24 00:48時点のこのコードでは処理時間が約１分２０秒かかっている
	for range time.Tick(5 * time.Minute) {
		upSignal, downSignal := client.CryptTaskV1()
		if len(upSignal) > 0 {
			line.PushCryptTrend(upSignal, lineBot)
			log.Println(upSignal)
		}
		if len(downSignal) > 0 {
			line.PushCryptTrend(downSignal, lineBot)
			log.Println(downSignal)
		}
	}
}
