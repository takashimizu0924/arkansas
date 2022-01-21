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
	go b.BBandSignalTask(coins, signalCh)
	go b.MacdSignalTask(coins, signalCh)
	go b.RsiSignalTask(coins, signalCh)
	for signal := range signalCh {
		signals = append(signals, signal)
	}
	log.Println(signals)
	log.Println("タスク終了")
}
