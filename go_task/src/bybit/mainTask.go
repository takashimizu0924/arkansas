package bybit

import (
	"log"
)

func (b *Bybit) CryptTaskV1() {
	signalCh := make(chan Signal)
	var signals []Signal
	coins := b.GetAllCoin("USDT")
	// for i := 0; i < len(coins); i++ {
	// 	go b.BBandSignalTask(coins[i], signalCh)
	// 	// go b.MacdSignalTask(coins[i])
	// 	// go b.RsiSignalTask(coins[i])
	// 	// result := <-signalCh
	// 	// signals = append(signals, result)
	// 	// log.Println(signals)
	// }
	var upList []Signal
	var downList []Signal
	var targetCoin []Signal
	go b.BBandSignalTask(coins, signalCh)
	go b.MacdSignalTask(coins, signalCh)
	go b.RsiSignalTask(coins, signalCh)
	for signal := range signalCh {
		signals = append(signals, signal)
		if signal.Sign == 0 {
			downList = append(downList, signal)
			if len(downList) > 0 {
				for i := 0; i < len(downList); i++ {
					if downList[i].CoinPair == signal.CoinPair {
						targetCoin = append(targetCoin, signal)
					}
				}
			}
		}
		if signal.Sign == 1 {
			upList = append(upList, signal)
			if len(upList) > 0 {
				for i := 0; i < len(upList); i++ {
					if upList[i].CoinPair == signal.CoinPair {
						log.Println(targetCoin)
						targetCoin = append(targetCoin, signal)
					}
				}
			}
		}
	}

	log.Println("Signals:", signals)
	log.Println("Up:", upList)
	log.Println("Down:", downList)
	log.Println(targetCoin)
	log.Println("タスク終了")
}
