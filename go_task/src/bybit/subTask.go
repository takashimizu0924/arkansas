package bybit

import (
	"log"
	"time"
)

func (b *Bybit) BBandSignalTask(symbol string, ch chan<- string) {
	for {
		before200Min := time.Now().Add(-200 * time.Hour).Unix()
		candles, err := b.GetKline(symbol, "60", before200Min, 0)
		if err != nil {
			log.Println(err)
		}
		var closes []float64
		for _, v := range candles {
			closes = append(closes, v.Close)
		}
		//ボリンジャーバンドを作成し、現在の価格がどの位置にあるかを調べる
		//今のところ出来ている 1/8 pm10:32
		up, mid, down := BBand(closes, 20, 2)

		ticker, err := b.Ticker(symbol)
		for _, tick := range ticker {
			if tick.LastPrice >= mid[len(mid)-1] {
				sign := downTrend(up[len(up)-1], tick.LastPrice)
				if sign {
					ch <- "下降サインが出ています"
				} else {
					ch <- "平均線より上にありますが、サインは出ていません"
				}
			}
			if tick.LastPrice <= mid[len(mid)-1] {
				sign := upTrend(down[len(down)-1], tick.LastPrice)
				if sign {
					ch <- "上昇サインが出ています"
				} else {
					ch <- "平均線より下にありますが、サインは出ていません"
				}
			}
		}
	}
}

func upTrend(downBand, price float64) bool {
	if downBand >= price {
		return true
	}
	if downBand <= price && downBand/price*100 > 99.8 {
		return true
	}
	return false
}

func downTrend(upBand, price float64) bool {
	if upBand <= price {
		return true
	}
	if upBand >= price && price/upBand*100 > 99.8 {
		return true
	}
	return false
}
