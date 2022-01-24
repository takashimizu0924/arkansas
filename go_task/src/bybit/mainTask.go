package bybit

import (
	"log"
	"sync"
)

func (b *Bybit) CryptTaskV1() (upSignals, downSignals []Signal) {

	bbandSignalCh := make(chan Signal)
	macdSignalCh := make(chan Signal)
	rsiSignalCh := make(chan Signal)
	coins := b.GetAllCoin("USDT")
	wg := &sync.WaitGroup{}

	var upList []Signal
	var downList []Signal

	go b.BBandSignalTask(coins, bbandSignalCh)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for signal := range bbandSignalCh {
			if signal.Sign == 0 {
				downList = append(downList, signal)
			}
			if signal.Sign == 1 {
				upList = append(upList, signal)
			}
		}
	}()

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
	if len(upSignals) > 0 {
		for i, up := range upSignals {
			ticker, err := b.Ticker(up.CoinPair)
			if err != nil {
				log.Println(err)
			}
			for _, tick := range ticker {
				upSignals[i].Price = tick.LastPrice
			}
		}
	}
	if len(downSignals) > 0 {
		for i, up := range downSignals {
			ticker, err := b.Ticker(up.CoinPair)
			if err != nil {
				log.Println(err)
			}
			for _, tick := range ticker {
				downSignals[i].Price = tick.LastPrice
			}
		}
	}

	log.Println("タスク終了")
	return upSignals, downSignals
}
