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


