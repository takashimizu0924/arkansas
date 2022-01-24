package bybit

import (
	"log"
	"strings"
	"time"
)

type Signal struct {
	Type     string
	CoinPair string
	Sign     int
	Price    float64
}

func (b *Bybit) GetBefore200LinearKline(symbol, interval string, from time.Duration) (candles []LinearOHLC, err error) {
	before200time := time.Now().Add(-200 * from).Unix()
	candles, err = b.GetLinearKline(symbol, interval, before200time, 0)
	if err != nil {
		return nil, err
	}
	return candles, nil
}

/*-------------------------------ボリンジャーバンド----------------------------------------------*/

func (b *Bybit) BBandSignalTask(symbols []string, signal chan<- Signal) {
	defer close(signal)
	for _, symbol := range symbols {
		candles, err := b.GetBefore200LinearKline(symbol, "5", time.Minute)
		if err != nil {
			log.Println(err)
		}
		var closes []float64
		for i := 0; i < len(candles); i++ {
			closes = append(closes, candles[i].Close)
		}
		up, mid, down := BBand(closes, 20, 2)
		ticker, err := b.Ticker(symbol)
		if err != nil {
			log.Println(err)
		}
		if len(up) == 0 || len(mid) == 0 || len(down) == 0 {
			return
		}
		for _, tick := range ticker {
			if tick.LastPrice >= up[len(up)-1] {
				signal <- Signal{
					Type:     "BBand",
					CoinPair: symbol,
					Sign:     0,
				}
			}
			if tick.LastPrice <= down[len(down)-1] {
				signal <- Signal{
					Type:     "BBand",
					CoinPair: symbol,
					Sign:     1,
				}
			}
		}
	}
}

/*-------------------------------------MACD-------------------------------------------------*/

func (b *Bybit) MacdSignalTask(symbols []string, signal chan<- Signal) {
	defer close(signal)
	for _, symbol := range symbols {
		candles, err := b.GetBefore200LinearKline(symbol, "5", time.Minute)
		if err != nil {
			log.Println(err)
		}
		var closes []float64
		for i := 0; i < len(candles); i++ {
			closes = append(closes, candles[i].Close)
		}
		macd, macdSignal, macdHist := Macd(closes, 12, 26, 9)
		if len(macd) == 0 || len(macdSignal) == 0 || len(macdHist) == 0 {
			return
		}
		if macd[len(macd)-1] > 0 && macdHist[len(macdHist)-1] > 0 && macd[len(macd)-1] > macdSignal[len(macdSignal)-1] {
			signal <- Signal{
				Type:     "MACD",
				CoinPair: symbol,
				Sign:     3,
			} //"MACDにより強気サインが出ています"
		}
		if macd[len(macd)-1] < 0 && macdHist[len(macdHist)-1] < 0 && macd[len(macd)-1] < macdSignal[len(macdSignal)-1] {
			signal <- Signal{
				Type:     "MACD",
				CoinPair: symbol,
				Sign:     2,
			} //"MACDにより弱気サインが出ています"
		}
		if macd[len(macd)-1] < 0 && macdHist[len(macdHist)-1] > 0 && macd[len(macd)-1] > macdSignal[len(macdSignal)-1] {
			signal <- Signal{
				Type:     "MACD",
				CoinPair: symbol,
				Sign:     1,
			} //"上昇の可能性があるためこの通貨を監視する"
		}
		if macd[len(macd)-1] > 0 && macdHist[len(macdHist)-1] < 0 && macd[len(macd)-1] < macdSignal[len(macdSignal)-1] {

			signal <- Signal{
				Type:     "MACD",
				CoinPair: symbol,
				Sign:     0,
			} //"下降の可能性があるためこの通貨を監視する"
		}
	}
}

/*------------------------------------------RSI---------------------------------------------------------*/

func (b *Bybit) RsiSignalTask(symbols []string, signal chan<- Signal) {
	defer close(signal)
	for _, symbol := range symbols {
		candles, err := b.GetBefore200LinearKline(symbol, "5", time.Minute)
		if err != nil {
			log.Println(err)
		}
		var closes []float64
		for i := 0; i < len(candles); i++ {
			closes = append(closes, candles[i].Close)
		}
		rsi := RSI(closes, 9)
		if rsi[len(rsi)-1] >= 70 {
			signal <- Signal{
				Type:     "RSI",
				CoinPair: symbol,
				Sign:     0,
			}
		}
		if rsi[len(rsi)-1] <= 30 {
			signal <- Signal{
				Type:     "RSI",
				CoinPair: symbol,
				Sign:     1,
			}
		}
	}
}

/*----------------------------USDTの通貨ペアのコインを全部取得する------------------------------------------------*/
func (b *Bybit) GetAllCoin(currency string) (coins []string) {
	allcoins, _ := b.GetSymbols()
	for _, coin := range allcoins {
		b := strings.Contains(coin.Name, "USDT")
		if b {
			coins = append(coins, coin.Name)
		}
	}
	return
}

/*--------------------------重複チェックタスク---------------------------------------*/
func DuplicateCheckTask(signals []Signal) []Signal {
	result := make([]Signal, 0, len(signals))
	encountered := map[string]bool{}
	//{BTCUSDT:false}
	for i := 0; i < len(signals); i++ {
		if !encountered[signals[i].CoinPair] {
			encountered[signals[i].CoinPair] = true
			//result = append(result, signals[i])
		} else {
			result = append(result, signals[i])
			log.Println(encountered[signals[i].CoinPair])
		}
	}
	return result
}

//TODO
/*
ある一定時間の間に平均線を超えなかったら損切りを行う処理を実装
多分２〜４時間ぐらい
そのプッシュ通知も実装する
*/
