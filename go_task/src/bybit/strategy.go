package bybit

import (
	"log"
	"time"
)

/*Strategy : 作戦 [テクニカル分析を利用したシグナルを判定するモジュール郡 ]*/

func (b *Bybit) StrategyV1(signals []Signal) []Signal {
	var signalList []Signal
	for i := 0; i > len(signals); i++ {
		candles, err := b.GetBefore200LinearKline(signals[i].CoinPair, "30", time.Minute)
		if err != nil {
			log.Println(err)
		}
		var closes []float64
		for i := 0; i < len(candles); i++ {
			closes = append(closes, candles[i].Close)
		}
		// up, mid, down := BBand(closes, 20, 2)
		// if len(up) == 0 || len(mid) == 0 || len(down) == 0 {
		// 	return nil
		// }
		// if signals[i].Price < closes[len(closes)-1] && closes[len(closes)-1] < closes[len(closes)-2] {
		// 	//現在価格が１個前より低く、１個前より２個前が低い状況 もしくは上線より現在価格が上の場合
		// 	//下降している
		// 	signals[i].Sign = 0
		// 	signalList = append(signalList, signals[i])
		// }
		if signals[i].Price < closes[len(closes)-1] {
			//現在価格が１個前より低く、１個前より２個前が低い状況 もしくは上線より現在価格が上の場合
			//下降している
			signals[i].Sign = 0
			signalList = append(signalList, signals[i])
		}
		// if signals[i].Price > closes[len(closes)-1] && closes[len(closes)-1] > closes[len(closes)-2] {
		// 	signals[i].Sign = 1
		// 	signalList = append(signalList, signals[i])
		// }
		if signals[i].Price > closes[len(closes)-1] {
			signals[i].Sign = 1
			signalList = append(signalList, signals[i])
		}
	}
	return signalList
}
