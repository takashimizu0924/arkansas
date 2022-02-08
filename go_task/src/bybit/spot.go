package bybit

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

type SpotSymbol struct {
	RetCode int                `json:"ret_code"`
	RetMsg  string             `json:"ret_msg"`
	ExtCode string             `json:"ext_code"`
	ExtInfo string             `json:"ext_info"`
	Result  []SpotSymbolResult `json:"result"`
}

type SpotSymbolResult struct {
	Name              string `json:"name"`
	Alias             string `json:"alias"`
	BaseCurrency      string `json:"baseCurrency"`
	QuoteCurrency     string `json:"quoteCurrency"`
	BasePrecision     string `json:"basePrecision"`
	QuotePrecision    string `json:"quotePrecision"`
	MinTradeQuantity  string `json:"minTradeQuantity"`
	MinTradeAmount    string `json:"minTradeAmount"`
	MinPricePrecision string `json:"minPricePrecision"`
	MaxTradeQuantity  string `json:"maxTradeQuantity"`
	MaxTradeAmount    string `json:"maxTradeAmount"`
	Category          int    `json:"category"`
}

type SpotKline struct {
	RetCode int             `json:"ret_code"`
	RetMsg  string          `json:"ret_msg"`
	Result  [][]interface{} `json:"result"`
	ExtCode string          `json:"ext_code"`
	ExtInfo string          `json:"ext_info"`
}

type SpotKlineOHLC struct {
	// StartTime        float64
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
	// EndTime          float64
	// QuoteAssetVolume string
	// Trades           float64
	// TakerBaseVolume  string
	// TakerQuoteVolume string
}

func (b *Bybit) SpotSignalTaskV1() {
	var symbolList []string
	symbols := b.GetSpotSymbols()
	for _, symbol := range symbols {
		b := strings.Contains(symbol.Name, "USDT")
		if b {
			symbolList = append(symbolList, symbol.Name)
		}
	}
	log.Println(symbolList)
	for _, symbol := range symbolList {
		OHLCList := b.GetSpotKlineOHLC(symbol, "5m")
		result := FluctuationRate(7, OHLCList)
		log.Println(symbol, ":", result)
	}
}

func (b *Bybit) GetSpotSymbols() (result []SpotSymbolResult) {
	var ret SpotSymbol
	params := map[string]interface{}{}

	_, err := b.PublicRequest(http.MethodGet, "/spot/v1/symbols", params, &ret)
	if err != nil {
		log.Println(err)
	}
	//log.Println("ret", ret)
	result = ret.Result
	return
}

func (b *Bybit) GetSpotKlineOHLC(symbpl string, interval string) (ohlcList []SpotKlineOHLC) {
	var ret SpotKline
	params := map[string]interface{}{}
	params["symbol"] = symbpl
	params["interval"] = interval
	params["limit"] = 50

	_, err := b.PublicRequest(http.MethodGet, "/spot/quote/v1/kline", params, &ret)
	if err != nil {
		log.Println(err)
	}

	var ohlc SpotKlineOHLC
	for i := 0; i < len(ret.Result); i++ {
		parseOpen, _ := ret.Result[i][1].(string)
		parseHigh, _ := ret.Result[i][2].(string)
		parseLow, _ := ret.Result[i][3].(string)
		parseClose, _ := ret.Result[i][4].(string)
		parseVolume, _ := ret.Result[i][5].(string)
		open, _ := strconv.ParseFloat(parseOpen, 64)
		high, _ := strconv.ParseFloat(parseHigh, 64)
		low, _ := strconv.ParseFloat(parseLow, 64)
		close, _ := strconv.ParseFloat(parseClose, 64)
		volume, _ := strconv.ParseFloat(parseVolume, 64)
		ohlc = SpotKlineOHLC{
			Open:   open,
			High:   high,
			Low:    low,
			Close:  close,
			Volume: volume,
		}
		ohlcList = append(ohlcList, ohlc)
	}
	return ohlcList
}

func FluctuationRate(limit int, ohlc []SpotKlineOHLC) (result float64) {
	if len(ohlc) == 0 {
		log.Println("no have OHLC...")
	}
	var last = ohlc[len(ohlc)-1]
	var before = ohlc[len(ohlc)-limit]
	result = (last.Close - before.Close) / before.Close * 100
	return
}
