package app

import (
	"arkansas/bybit"
	"arkansas/slack"
)

// var b bybit.Bybit

func BybitAutoTrade(b bybit.Bybit) {
	signalCh := make(chan bybit.Signal)
	tradeCh := make(chan bybit.Signal)
	bot := slack.NewSlack()
	allCoins := b.GetAllCoin("USDT")
	go b.CryptTaskV3(allCoins, signalCh)
	go b.CryptTrading(signalCh, tradeCh)
	go slack.TradingPush(bot, tradeCh)
}
