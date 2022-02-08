package bybit

import "github.com/markcheno/go-talib"

//ボリンジャーバンド
func BBand(closes []float64, n int, k float64) (up, mid, down []float64) {
	var bbandUP []float64
	var bbandMID []float64
	var bbandDOWN []float64
	if n <= len(closes) {
		up, mid, down := talib.BBands(closes, n, k, k, 0)
		bbandUP = up
		bbandMID = mid
		bbandDOWN = down
	}
	return bbandUP, bbandMID, bbandDOWN
}

func Macd(close []float64, f, s, sign int) (macd, macdSignal, macdHist []float64) {
	if s <= len(close) {
		m, ms, mh := talib.Macd(close, f, s, sign)
		macd = m
		macdSignal = ms
		macdHist = mh
	}
	return macd, macdSignal, macdHist
}

func RSI(close []float64, period int) (rsi []float64) {
	if period <= len(close) {
		rsi = talib.Rsi(close, period)
	}
	return
}

func EMA(close []float64, period int) (ema []float64) {
	if period <= len(close) {
		ema = talib.Ema(close, period)
	}
	return
}
