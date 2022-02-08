package bybit

import (
	"log"
	"math"
	"sort"
	"time"
)

/*Strategy : 作戦 [テクニカル分析を利用したシグナルを判定するモジュール郡 ]*/

func (b *Bybit) StrategyV1(signals []Signal) []Signal {
	var signalList []Signal
	log.Println("strategyTask")
	for _, signal := range signals {
		candles, err := b.GetBefore200LinearKline(signal.CoinPair, "5", time.Minute)
		if err != nil {
			log.Println(err)
		}
		log.Println(signal)
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
		log.Println("signalsPrice:", signal.Price)
		log.Println("candlePrice:", closes[len(closes)-2])
		if signal.Price < closes[len(closes)-2] {
			//現在価格が１個前より低く、１個前より２個前が低い状況 もしくは上線より現在価格が上の場合
			//下降している
			signal.Sign = 0
			signalList = append(signalList, signal)
		}
		// if signals[i].Price > closes[len(closes)-1] && closes[len(closes)-1] > closes[len(closes)-2] {
		// 	signals[i].Sign = 1
		// 	signalList = append(signalList, signals[i])
		// }
		if signal.Price > closes[len(closes)-2] {
			signal.Sign = 1
			signalList = append(signalList, signal)
		}
	}
	return signalList
}

//ドテンくん
func (b *Bybit) StrategyV2(candle LinearOHLC, candles []LinearOHLC, signal int) {
	var sum_range_Buy float64
	var sum_range_Sell float64
	log.Println(candles[0])
	sort.Slice(candles, func(i, j int) bool {
		return candles[i].OpenTime > candles[j].OpenTime
	})
	log.Println(candles[0])
	if signal == 0 {
		for i := 0; i < 5; i++ {
			high := math.Round(candles[i].High * 100000 / 100000)
			low := math.Round(candles[i].Low * 100000 / 100000)
			sum_range_Buy += high - low
			log.Println(high, low)
		}
	}
	if signal == 1 {
		for i := 0; i < -5; i-- {
			sum_range_Sell += candles[i].High - candles[i].Low
		}
	}

	ave_range_Buy := sum_range_Buy / 5
	ave_range_Sell := sum_range_Sell / 5
	log.Println(ave_range_Buy, ave_range_Sell)
	if candle.High-candle.Open > ave_range_Buy*1.6 {
		log.Println(candle.High - candle.Open)
		log.Println("ドテン君上昇")
		log.Println(ave_range_Buy)
	}
	if candle.Open-candle.Low > ave_range_Sell*1.6 {
		log.Println(candle.Open - candle.Low)
		log.Println("ドテン君下降")
		log.Println(ave_range_Sell)
	}
}
func SumCandle2(candle OHLC, candles []OHLC) string {
	var sum_range_Buy float64

	sort.Slice(candles, func(i, j int) bool {
		return candles[i].OpenTime > candles[j].OpenTime
	})
	for i := 0; i < 5; i++ {
		sum_range_Buy += candles[i].High - candles[i].Low
	}

	ave_range_Buy := ((sum_range_Buy * 1000) / 5) / 1000
	// log.Println(ave_range_Buy)
	// log.Println(ave_range_Buy * 1.6)
	target_range := ave_range_Buy * 1.6
	// if (candle.High*1000 - candle.Low*1000) / 1000 > ave_range_Buy {
	// 	log.Println("ドテンUP！")
	// }
	//log.Println((candle.High*1000 - candle.Low*1000) / 1000)
	next_side_Buy := candle.High - candle.Open
	next_side_Sell := candle.Open - candle.Low
	if next_side_Buy > target_range {
		//log.Println("UP")
		return "up"
	}
	if next_side_Sell > target_range {
		//log.Println("DOWN")
		return "down"
	}
	// log.Println("Open", candle.Open)
	// log.Println("High", candle.High)
	// log.Println("Low", candle.Low)
	return ""
}

func SumCandle(candle LinearOHLC, candles []LinearOHLC) string {
	var sum_range_Buy float64

	sort.Slice(candles, func(i, j int) bool {
		return candles[i].OpenTime > candles[j].OpenTime
	})
	for i := 0; i < 5; i++ {
		sum_range_Buy += candles[i].High - candles[i].Low
	}

	ave_range_Buy := ((sum_range_Buy * 1000) / 5) / 1000
	// log.Println(ave_range_Buy)
	// log.Println(ave_range_Buy * 1.6)
	target_range := ave_range_Buy * 1.6
	// if (candle.High*1000 - candle.Low*1000) / 1000 > ave_range_Buy {
	// 	log.Println("ドテンUP！")
	// }
	//log.Println((candle.High*1000 - candle.Low*1000) / 1000)
	next_side_Buy := candle.High - candle.Open
	next_side_Sell := candle.Open - candle.Low
	if next_side_Buy > target_range {
		//log.Println("UP")
		return "up"
	}
	if next_side_Sell > target_range {
		//log.Println("DOWN")
		return "down"
	}
	// log.Println("Open", candle.Open)
	// log.Println("High", candle.High)
	// log.Println("Low", candle.Low)
	return ""
}

func Round(f float64) float64 {
	return math.Floor(f + .8)
}
