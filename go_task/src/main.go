package main

import (
	"arkansas/bybit"
	"arkansas/config"

	//"arkansas/line"
	"log"
)

func main() {
	config.NewConfig()
	log.Println("[ Start Arkansas ]")
	//slackBot := slack.NewSlack()

	client := bybit.NewBybit(config.Config.ApiKey, config.Config.ApiSecret, true)
	// for {
	// 	app.BybitAutoTrade(*client)
	// }
	// lineBot, err := line.NewLine()
	// if err != nil {
	// 	log.Println(err)
	client.SpotSignalTaskV1()
	//client.RealtimeGet()

	// }
	// go client.RealtimeGet()
	// 1/24 00:48時点のこのコードでは処理時間が約１分２０秒かかっている
	// for {
	// 	upSignal, downSignal := client.CryptTaskV1()
	// 	if len(upSignal) > 0 {
	// 		line.PushCryptTrend(upSignal, lineBot)
	// 		log.Println(upSignal)
	// 	}
	// 	if len(downSignal) > 0 {
	// 		line.PushCryptTrend(downSignal, lineBot)
	// 		log.Println(downSignal)
	// 	}
	// }
	// for {
	// 	upSignal, downSignal := client.CryptTaskV4()
	// 	if len(upSignal) > 0 {
	// 		//line.PushCryptTrend(upSignal, lineBot)
	// 		slack.PostTextMessage(slackBot, upSignal)
	// 		// log.Println(upSignal)
	// 	}
	// 	if len(downSignal) > 0 {
	// 		//	line.PushCryptTrend(downSignal, lineBot)
	// 		slack.PostTextMessage(slackBot, downSignal)
	// 		//log.Println(downSignal)
	// 	}
	// }
	log.Println("[ Finished Arkansas ]")
}
