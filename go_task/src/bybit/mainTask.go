package bybit

import (
	"log"
	"sync"
	"time"
)

func (b *Bybit) CryptTaskV1() (upSignals, downSignals []Signal) {

	// bbandSignalCh := make(chan Signal)
	macdSignalCh := make(chan Signal)
	rsiSignalCh := make(chan Signal)
	coins := b.GetAllCoin("USDT")
	wg := &sync.WaitGroup{}

	var upList []Signal
	var downList []Signal

	// go b.BBandSignalTask(coins, bbandSignalCh)
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for signal := range bbandSignalCh {
	// 		if signal.Sign == 0 {
	// 			downList = append(downList, signal)
	// 		}
	// 		if signal.Sign == 1 {
	// 			upList = append(upList, signal)
	// 		}
	// 	}
	// }()

	go b.MacdSignalTask(coins, macdSignalCh)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for signal := range macdSignalCh {
			if signal.Sign == 0 {
				downList = append(downList, signal)
			}
			if signal.Sign == 1 {
				upList = append(upList, signal)
			}
		}
	}()

	go b.RsiSignalTask(coins, rsiSignalCh)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for signal := range rsiSignalCh {
			if signal.Sign == 0 {
				downList = append(downList, signal)
			}
			if signal.Sign == 1 {
				upList = append(upList, signal)
			}
		}
	}()
	wg.Wait()

	upSignals = DuplicateCheckTask(upList)
	downSignals = DuplicateCheckTask(downList)

	before200time := time.Now().Add(-200 * time.Minute).Unix()
	candles, _ := b.GetLinearKline("JASMYUSDT", "5", before200time, 0)

	SumCandle(candles[len(candles)-1], candles)
	//b.StrategyV2(candles[len(candles)-1], candles, 0)
	// if len(upSignals) > 0 {
	// 	for i, up := range upSignals {
	// 		ticker, err := b.Ticker(up.CoinPair)
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 		for _, tick := range ticker {
	// 			upSignals[i].Price = tick.LastPrice
	// 		}
	// 	}
	// }
	// if len(downSignals) > 0 {
	// 	for i, up := range downSignals {
	// 		ticker, err := b.Ticker(up.CoinPair)
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 		for _, tick := range ticker {
	// 			downSignals[i].Price = tick.LastPrice
	// 		}
	// 	}
	// }
	// lastSignal := b.StrategyV1(upSignals)
	// lastSignal2 := b.StrategyV1(downSignals)

	log.Println("タスク終了")
	return upSignals, downSignals
}

func (b *Bybit) CryptTaskV2() (upSignals, downSignals []Signal) {
	var upSignal []Signal
	var downSignal []Signal
	allCoins := b.GetAllCoin("USDT")
	for i := 0; i < len(allCoins); i++ {
		before200time := time.Now().Add(-1000 * time.Minute).Unix()
		candles, _ := b.GetLinearKline(allCoins[i], "5", before200time, 0)
		signal := SumCandle(candles[len(candles)-1], candles)
		var latestCandle = candles[0]
		var closes []float64
		for i := 0; i < len(candles); i++ {
			closes = append(closes, candles[i].Close)
		}
		//上昇予測の場合
		if signal == "up" {
			ema := EMA(closes, 9)
			_, _, down := BBand(closes, 20, 2)
			log.Println("[UP]", latestCandle.Symbol, ":", latestCandle.Close)
			if latestCandle.Close >= down[20] && ema[20] >= ema[21] && latestCandle.Close >= ema[20] {
				log.Println("[---->UP 1]")
				log.Println("通貨情報-->", latestCandle.Symbol, ":", latestCandle.Close)
				upSignal = append(upSignal,
					Signal{
						Type:     "判定Ver.2",
						CoinPair: allCoins[i],
						Sign:     1,
						Price:    latestCandle.Close})
			}
		}
		//下降予測の場合
		if signal == "down" {
			ema := EMA(closes, 9)
			up, _, _ := BBand(closes, 20, 2)
			log.Println("[DOWN]", latestCandle.Symbol, ":", latestCandle.Close)
			if latestCandle.Close <= up[20] && ema[20] <= ema[21] && latestCandle.Close <= ema[20] {
				log.Println("[---->DOWN]")
				log.Println("通貨情報-->", latestCandle.Symbol, ":", latestCandle.Close)
				downSignal = append(downSignal,
					Signal{
						Type:     "判定Ver.2",
						CoinPair: allCoins[i],
						Sign:     0,
						Price:    latestCandle.Close})
			}

		}
		upSignals = upSignal
		downSignals = downSignal
	}
	log.Println("CryptTaskV2--->Done")
	return upSignals, downSignals
}

func (b *Bybit) CryptTaskV3(allCoins []string, signalCh chan<- Signal) {
	defer close(signalCh)
	// allCoins := b.GetAllCoin("USDT")
	log.Println(len(allCoins))
	for i := 0; i < len(allCoins); i++ {
		before200time := time.Now().Add(-1000 * time.Minute).Unix()
		candles, _ := b.GetLinearKline(allCoins[i], "5", before200time, 0)
		log.Println(len(candles))
		signal := SumCandle(candles[len(candles)-1], candles)
		var latestCandle = candles[0]
		var closes []float64
		for i := 0; i < len(candles); i++ {
			closes = append(closes, candles[i].Close)
		}
		//上昇予測の場合
		if signal == "up" {
			ema := EMA(closes, 9)
			_, _, down := BBand(closes, 20, 2)

			if latestCandle.Close >= down[20] && ema[20] >= ema[21] && latestCandle.Close >= ema[20] {
				log.Println("[---->UP 1]")
				log.Println("通貨情報-->", latestCandle.Symbol, ":", latestCandle.Close)
				signalCh <- Signal{
					Type:     "判定Ver.2",
					CoinPair: allCoins[i],
					Sign:     1,
					Price:    latestCandle.Close}
			}
		}
		//下降予測の場合
		if signal == "down" {
			ema := EMA(closes, 9)
			up, _, _ := BBand(closes, 20, 2)

			if latestCandle.Close <= up[20] && ema[20] <= ema[21] && latestCandle.Close <= ema[20] {
				log.Println("[---->DOWN]")
				log.Println("通貨情報-->", latestCandle.Symbol, ":", latestCandle.Close)
				signalCh <- Signal{
					Type:     "判定Ver.2",
					CoinPair: allCoins[i],
					Sign:     0,
					Price:    latestCandle.Close}
			}

		}
	}
	log.Println("CryptTaskV2--->Done")
}

func (b *Bybit) CryptTaskV4() (upSignals, downSignals []Signal) {
	var upSignal []Signal
	var downSignal []Signal
	allCoins := b.GetAllCoin("USDT")
	for i := 0; i < len(allCoins); i++ {
		before200time := time.Now().Add(-1000 * time.Minute).Unix()
		candles, _ := b.GetKline(allCoins[i], "5", before200time, 0)
		log.Println(candles)
		signal := SumCandle2(candles[len(candles)-1], candles)
		var latestCandle = candles[0]
		var closes []float64
		for i := 0; i < len(candles); i++ {
			closes = append(closes, candles[i].Close)
		}
		//上昇予測の場合
		if signal == "up" {
			ema := EMA(closes, 9)
			_, _, down := BBand(closes, 20, 2)
			log.Println("[UP]", latestCandle.Symbol, ":", latestCandle.Close)
			if latestCandle.Close >= down[20] && ema[20] >= ema[21] && latestCandle.Close >= ema[20] {
				log.Println("[---->UP 1]")
				log.Println("通貨情報-->", latestCandle.Symbol, ":", latestCandle.Close)
				upSignal = append(upSignal,
					Signal{
						Type:     "判定Ver.2",
						CoinPair: allCoins[i],
						Sign:     1,
						Price:    latestCandle.Close})
			}
		}
		//下降予測の場合
		if signal == "down" {
			ema := EMA(closes, 9)
			up, _, _ := BBand(closes, 20, 2)
			log.Println("[DOWN]", latestCandle.Symbol, ":", latestCandle.Close)
			if latestCandle.Close <= up[20] && ema[20] <= ema[21] && latestCandle.Close <= ema[20] {
				log.Println("[---->DOWN]")
				log.Println("通貨情報-->", latestCandle.Symbol, ":", latestCandle.Close)
				downSignal = append(downSignal,
					Signal{
						Type:     "判定Ver.2",
						CoinPair: allCoins[i],
						Sign:     0,
						Price:    latestCandle.Close})
			}

		}
		upSignals = upSignal
		downSignals = downSignal
	}
	log.Println("CryptTaskV2--->Done")
	return upSignals, downSignals
}
