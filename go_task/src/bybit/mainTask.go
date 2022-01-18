package bybit

import "log"

func (b *Bybit) CryptTaskV1(symbol string) {
	var tickerChannel = make(chan string)
	go b.BBandSignalTask(symbol, tickerChannel)
	for chanel := range tickerChannel {
		log.Println(chanel)
	}
}
