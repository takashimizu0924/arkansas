package bybit

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func (b *Bybit) GetOrderBook(symbol string) (result OrderBook, err error) {
	var ret GetOrderBookResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	_, err = b.PublicRequest(http.MethodGet, "/v2/public/orderBook/L2", params, &ret)
	if err != nil {
		log.Println("action=GetOrderBook ==>", err.Error())
		return
	}
	for _, v := range ret.Result {
		if v.Side == "Sell" {
			result.Asks = append(result.Asks, Item{
				Price: v.Price,
				Size:  v.Size,
			})
		} else if v.Side == "Buy" {
			result.Bids = append(result.Bids, Item{
				Price: v.Price,
				Size:  v.Size,
			})
		}
	}
	sort.Slice(result.Asks, func(i, j int) bool {
		return result.Asks[i].Price < result.Asks[j].Price
	})

	sort.Slice(result.Bids, func(i, j int) bool {
		return result.Bids[i].Price > result.Bids[j].Price
	})

	var timeNow float64
	timeNow, err = strconv.ParseFloat(ret.TimeNow, 64) // 1582011750.433202
	if err != nil {
		return
	}
	result.Time = time.Unix(0, int64(timeNow*1e9))
	return
}

func (b *Bybit) GetKline(symbol, interval string, from int64, limit int) (result []OHLC, err error) {
	var ret GetKlineResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["interval"] = interval
	params["from"] = from
	if limit > 0 {
		params["limit"] = limit
	}
	_, err = b.PublicRequest(http.MethodGet, "/v2/public/kline/list", params, &ret)
	if err != nil {
		log.Println("action=GetKline ==>", err.Error())
		return
	}
	result = ret.Result
	return
}

func (b *Bybit) Ticker(symbol string) (result []Ticker, err error) {
	var ret GetTickersResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	_, err = b.PublicRequest(http.MethodGet, "/v2/public/tickers", params, &ret)

	result = ret.Result
	return
}

func (b *Bybit) GetTradingRecords(symbol string, from int64, limit int) (result []TradingRecord, err error) {
	var ret GetTradingRecordsResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if from > 0 {
		params["from"] = from
	}
	if limit > 0 {
		params["limit"] = limit
	}

	_, err = b.PublicRequest(http.MethodGet, "/v2/public/trading-records", params, &ret)
	if err != nil {
		log.Println("action=GetTradingRecords ==>", err.Error())
		return
	}

	result = ret.Result
	return
}

func (b *Bybit) GetSymbols() (result []SymbolInfo, err error) {
	var ret GetSymbolsResult
	params := map[string]interface{}{}
	_, err = b.PublicRequest(http.MethodGet, "/v2/public/symbols", params, &ret)

	result = ret.Result
	return
}
