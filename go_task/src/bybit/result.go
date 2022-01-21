package bybit

import (
	"encoding/json"
	sjson "encoding/json"
	"strings"
	"time"
)

type BaseResult struct {
	RetCode         int         `json:"ret_code"`
	RetMsg          string      `json:"ret_msg"`
	ExtCode         string      `json:"ext_code"`
	Result          interface{} `json:"result"`
	TimeNow         string      `json:"time_now"`
	RateLimitStatus int         `json:"rate_limit_status"`
}

type Item struct {
	Price float64 `json:"price,string"`
	Size  float64 `json:"size"`
}

type OrderBook struct {
	Asks []Item    `json:"asks"`
	Bids []Item    `json:"bids"`
	Time time.Time `json:"time"`
}

type RawItem struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
	Size   float64 `json:"size"`
	Side   string  `json:"side"` // Buy/Sell
}

type GetOrderBookResult struct {
	RetCode int       `json:"ret_code"`
	RetMsg  string    `json:"ret_msg"`
	ExtCode string    `json:"ext_code"`
	ExtInfo string    `json:"ext_info"`
	Result  []RawItem `json:"result"`
	TimeNow string    `json:"time_now"`
}

type OHLC struct {
	Symbol   string  `json:"symbol"`
	Interval string  `json:"interval"`
	OpenTime int64   `json:"open_time"`
	Open     float64 `json:"open,string"`
	High     float64 `json:"high,string"`
	Low      float64 `json:"low,string"`
	Close    float64 `json:"close,string"`
	Volume   float64 `json:"volume,string"`
	Turnover float64 `json:"turnover,string"`
}

type LinearOHLC struct {
	ID       int64   `json:"id"`
	Period   string  `json:"period"`
	Start_at int64   `json:"start_at"`
	Symbol   string  `json:"symbol"`
	Interval string  `json:"interval"`
	OpenTime int64   `json:"open_time"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Turnover float64 `json:"turnover"`
}

type GetLinearKlineResult struct {
	RetCode int          `json:"ret_code"`
	RetMsg  string       `json:"ret_msg"`
	ExtCode string       `json:"ext_code"`
	ExtInfo string       `json:"ext_info"`
	Result  []LinearOHLC `json:"result"`
	TimeNow string       `json:"time_now"`
}

type GetKlineResult struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  []OHLC `json:"result"`
	TimeNow string `json:"time_now"`
}

type Ticker struct {
	Symbol               string    `json:"symbol"`
	BidPrice             float64   `json:"bid_price,string"`
	AskPrice             float64   `json:"ask_price,string"`
	LastPrice            float64   `json:"last_price,string"`
	LastTickDirection    string    `json:"last_tick_direction"`
	PrevPrice24H         float64   `json:"prev_price_24h,string"`
	Price24HPcnt         float64   `json:"price_24h_pcnt,string"`
	HighPrice24H         float64   `json:"high_price_24h,string"`
	LowPrice24H          float64   `json:"low_price_24h,string"`
	PrevPrice1H          float64   `json:"prev_price_1h,string"`
	Price1HPcnt          float64   `json:"price_1h_pcnt,string"`
	MarkPrice            float64   `json:"mark_price,string"`
	IndexPrice           float64   `json:"index_price,string"`
	OpenInterest         float64   `json:"open_interest"`
	OpenValue            float64   `json:"open_value,string"`
	TotalTurnover        float64   `json:"total_turnover,string"`
	Turnover24H          float64   `json:"turnover_24h,string"`
	TotalVolume          float64   `json:"total_volume"`
	Volume24H            float64   `json:"volume_24h"`
	FundingRate          float64   `json:"funding_rate,string"`
	PredictedFundingRate float64   `json:"predicted_funding_rate,string"`
	NextFundingTime      time.Time `json:"next_funding_time"`
	CountdownHour        int       `json:"countdown_hour"`
}

type GetTickersResult struct {
	RetCode int      `json:"ret_code"`
	RetMsg  string   `json:"ret_msg"`
	ExtCode string   `json:"ext_code"`
	ExtInfo string   `json:"ext_info"`
	Result  []Ticker `json:"result"`
	TimeNow string   `json:"time_now"`
}

type TradingRecord struct {
	ID     int       `json:"id"`
	Symbol string    `json:"symbol"`
	Price  float64   `json:"price"`
	Qty    int       `json:"qty"`
	Side   string    `json:"side"`
	Time   time.Time `json:"time"`
}

type GetTradingRecordsResult struct {
	RetCode int             `json:"ret_code"`
	RetMsg  string          `json:"ret_msg"`
	ExtCode string          `json:"ext_code"`
	ExtInfo string          `json:"ext_info"`
	Result  []TradingRecord `json:"result"`
	TimeNow string          `json:"time_now"`
}

type LeverageFilter struct {
	MinLeverage  int     `json:"min_leverage"`
	MaxLeverage  int     `json:"max_leverage"`
	LeverageStep float64 `json:"leverage_step,string"`
}

type PriceFilter struct {
	MinPrice float64 `json:"min_price,string"`
	MaxPrice float64 `json:"max_price,string"`
	TickSize float64 `json:"tick_size,string"`
}

type LotSizeFilter struct {
	MaxTradingQty int `json:"max_trading_qty"`
	MinTradingQty int `json:"min_trading_qty"`
	QtyStep       int `json:"qty_step"`
}

type SymbolInfo struct {
	Name           string         `json:"name"`
	BaseCurrency   string         `json:"base_currency"`
	QuoteCurrency  string         `json:"quote_currency"`
	PriceScale     int            `json:"price_scale"`
	TakerFee       float64        `json:"taker_fee,string"`
	MakerFee       float64        `json:"maker_fee,string"`
	LeverageFilter LeverageFilter `json:"leverage_filter"`
	PriceFilter    PriceFilter    `json:"price_filter"`
	LotSizeFilter  LotSizeFilter  `json:"lot_size_filter"`
}

type GetSymbolsResult struct {
	RetCode int          `json:"ret_code"`
	RetMsg  string       `json:"ret_msg"`
	ExtCode string       `json:"ext_code"`
	ExtInfo string       `json:"ext_info"`
	Result  []SymbolInfo `json:"result"`
	TimeNow string       `json:"time_now"`
}

type Balance struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

type GetBalanceResult struct {
	RetCode          int                  `json:"ret_code"`
	RetMsg           string               `json:"ret_msg"`
	ExtCode          string               `json:"ext_code"`
	ExtInfo          string               `json:"ext_info"`
	Result           GetBalanceResultData `json:"result"`
	TimeNow          string               `json:"time_now"`
	RateLimitStatus  int                  `json:"rate_limit_status"`
	RateLimitResetMs int64                `json:"rate_limit_reset_ms"`
	RateLimit        int                  `json:"rate_limit"`
}

type GetBalanceResultData struct {
	BTC  Balance `json:"BTC"`
	ETH  Balance `json:"ETH"`
	EOS  Balance `json:"EOS"`
	XRP  Balance `json:"XRP"`
	USDT Balance `json:"USDT"`
}

type CreateOrderResult struct {
	RetCode         int    `json:"ret_code"`
	RetMsg          string `json:"ret_msg"`
	ExtCode         string `json:"ext_code"`
	Result          Order  `json:"result"`
	TimeNow         string `json:"time_now"`
	RateLimitStatus int    `json:"rate_limit_status"`
}

type OrderLite struct {
	OrderID string `json:"order_id"`
}

type ReplaceOrderResult struct {
	RetCode         int       `json:"ret_code"`
	RetMsg          string    `json:"ret_msg"`
	ExtCode         string    `json:"ext_code"`
	Result          OrderLite `json:"result"`
	TimeNow         string    `json:"time_now"`
	RateLimitStatus int       `json:"rate_limit_status"`
}

type CancelOrderResult struct {
	RetCode         int    `json:"ret_code"`
	RetMsg          string `json:"ret_msg"`
	ExtCode         string `json:"ext_code"`
	Result          Order  `json:"result"`
	TimeNow         string `json:"time_now"`
	RateLimitStatus int    `json:"rate_limit_status"`
}

type OrderListResultData struct {
	Data        []OrderV3 `json:"data"`
	CurrentPage int       `json:"current_page"`
	LastPage    int       `json:"last_page"`
}

type OrderListResult struct {
	RetCode         int                 `json:"ret_code"`
	RetMsg          string              `json:"ret_msg"`
	ExtCode         string              `json:"ext_code"`
	Result          OrderListResultData `json:"result"`
	TimeNow         string              `json:"time_now"`
	RateLimitStatus int                 `json:"rate_limit_status"`
}

// Order ...
type Order struct {
	OrderID       string     `json:"order_id"`
	StopOrderID   string     `json:"stop_order_id"`
	UserID        int        `json:"user_id"`
	Symbol        string     `json:"symbol"`
	Side          string     `json:"side"`
	OrderType     string     `json:"order_type"`
	Price         string     `json:"price"`
	Qty           string     `json:"qty"`
	TimeInForce   string     `json:"time_in_force"`
	OrderStatus   string     `json:"order_status"`
	LastExecTime  string     `json:"last_exec_time"`
	LastExecPrice float64    `json:"last_exec_price"`
	LeavesQty     string     `json:"leaves_qty"`
	CumExecQty    string     `json:"cum_exec_qty"`
	CumExecValue  string     `json:"cum_exec_value"`
	CumExecFee    string     `json:"cum_exec_fee"`
	RejectReason  string     `json:"reject_reason"`
	OrderLinkID   string     `json:"order_link_id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	ExtFields     *ExtFields `json:"ext_fields,omitempty"`
}

type OrderV3 struct {
	UserID       int       `json:"user_id"`
	OrderStatus  string    `json:"order_status"`
	Symbol       string    `json:"symbol"`
	Side         string    `json:"side"`
	OrderType    string    `json:"order_type"`
	Price        string    `json:"price"`
	Qty          string    `json:"qty"`
	TimeInForce  string    `json:"time_in_force"`
	OrderLinkID  string    `json:"order_link_id"`
	OrderID      string    `json:"order_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	LeavesQty    string    `json:"leaves_qty"`
	LeavesValue  string    `json:"leaves_value"`
	CumExecQty   string    `json:"cum_exec_qty"`
	CumExecValue string    `json:"cum_exec_value"`
	CumExecFee   string    `json:"cum_exec_fee"`
	RejectReason string    `json:"reject_reason"`
}

type ExtFields struct {
	ReduceOnly  bool   `json:"reduce_only"`
	OpFrom      string `json:"op_from"`
	Remark      string `json:"remark"`
	OReqNum     int64  `json:"o_req_num"`
	XreqType    string `json:"xreq_type"`
	CrossStatus string `json:"cross_status,omitempty"`
}

type InExtFields struct {
	ReduceOnly  bool   `json:"reduce_only"`
	OpFrom      string `json:"op_from"`
	Remark      string `json:"remark"`
	OReqNum     int64  `json:"o_req_num"`
	XreqType    string `json:"xreq_type"`
	CrossStatus string `json:"cross_status,omitempty"`
}

func (e *ExtFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ExtFields) UnmarshalJSON(b []byte) error {
	s := string(b)
	if strings.HasPrefix(s, "[") {
		return nil
	}
	o := InExtFields{}
	if err := json.Unmarshal(b, &o); err == nil {
		e.ReduceOnly = o.ReduceOnly
		e.OpFrom = o.OpFrom
		e.Remark = o.Remark
		e.OReqNum = o.OReqNum
		e.XreqType = o.XreqType
		e.CrossStatus = o.CrossStatus
		return nil
	} else {
		return err
	}
}

type GetLeverageResult struct {
	RetCode         int                     `json:"ret_code"`
	RetMsg          string                  `json:"ret_msg"`
	ExtCode         string                  `json:"ext_code"`
	Result          map[string]LeverageItem `json:"result"`
	TimeNow         string                  `json:"time_now"`
	RateLimitStatus int                     `json:"rate_limit_status"`
}

type LeverageItem struct {
	Leverage int `json:"leverage"`
}

type PositionV1 struct {
	ID                  int                `json:"id"`
	UserID              int                `json:"user_id"`
	RiskID              int                `json:"risk_id"`
	Symbol              string             `json:"symbol"`
	Size                float64            `json:"size"`
	Side                string             `json:"side"`
	EntryPrice          float64            `json:"entry_price"`
	LiqPrice            float64            `json:"liq_price"`
	BustPrice           float64            `json:"bust_price"`
	TakeProfit          float64            `json:"take_profit"`
	StopLoss            float64            `json:"stop_loss"`
	TrailingStop        float64            `json:"trailing_stop"`
	PositionValue       float64            `json:"position_value"`
	Leverage            float64            `json:"leverage"`
	PositionStatus      string             `json:"position_status"`
	AutoAddMargin       float64            `json:"auto_add_margin"`
	OrderMargin         float64            `json:"order_margin"`
	PositionMargin      float64            `json:"position_margin"`
	OccClosingFee       float64            `json:"occ_closing_fee"`
	OccFundingFee       float64            `json:"occ_funding_fee"`
	ExtFields           *PositionExtFields `json:"ext_fields"`
	WalletBalance       float64            `json:"wallet_balance"`
	CumRealisedPnl      float64            `json:"cum_realised_pnl"`
	CumCommission       float64            `json:"cum_commission"`
	RealisedPnl         float64            `json:"realised_pnl"`
	DeleverageIndicator float64            `json:"deleverage_indicator"`
	OcCalcData          string             `json:"oc_calc_data"`
	CrossSeq            float64            `json:"cross_seq"`
	PositionSeq         float64            `json:"position_seq"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
	UnrealisedPnl       float64            `json:"unrealised_pnl"`
}

type Position struct {
	ID                  int                `json:"id"`
	UserID              int                `json:"user_id"`
	RiskID              int                `json:"risk_id"`
	Symbol              string             `json:"symbol"`
	Size                float64            `json:"size"`
	Side                string             `json:"side"`
	EntryPrice          float64            `json:"entry_price,string"`
	LiqPrice            float64            `json:"liq_price,string"`
	BustPrice           float64            `json:"bust_price,string"`
	TakeProfit          float64            `json:"take_profit,string"`
	StopLoss            float64            `json:"stop_loss,string"`
	TrailingStop        float64            `json:"trailing_stop,string"`
	PositionValue       float64            `json:"position_value,string"`
	Leverage            float64            `json:"leverage,string"`
	PositionStatus      string             `json:"position_status"`
	AutoAddMargin       float64            `json:"auto_add_margin"`
	OrderMargin         float64            `json:"order_margin,string"`
	PositionMargin      float64            `json:"position_margin,string"`
	OccClosingFee       float64            `json:"occ_closing_fee,string"`
	OccFundingFee       float64            `json:"occ_funding_fee,string"`
	ExtFields           *PositionExtFields `json:"ext_fields"`
	WalletBalance       float64            `json:"wallet_balance,string"`
	CumRealisedPnl      float64            `json:"cum_realised_pnl,string"`
	CumCommission       float64            `json:"cum_commission,string"`
	RealisedPnl         float64            `json:"realised_pnl,string"`
	DeleverageIndicator float64            `json:"deleverage_indicator"`
	OcCalcData          string             `json:"oc_calc_data"`
	CrossSeq            float64            `json:"cross_seq"`
	PositionSeq         float64            `json:"position_seq"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
	UnrealisedPnl       float64            `json:"unrealised_pnl"`
}

type PositionExtFields struct {
	Remark string `json:"_remark"`
}

type PositionListResult struct {
	BaseResult
	ExtInfo        interface{} `json:"ext_info"`
	Result         PositionV3  `json:"result"`
	RateLimitReset int         `json:"rate_limit_reset"`
	RateLimit      int         `json:"rate_limit"`
}
type PositionListResultData struct {
	//IsValid bool       `json:"is_valid"`
	Data PositionV3 `json:"data"`
}

type GetPositionResult struct {
	BaseResult
	ExtInfo        interface{} `json:"ext_info"`
	Result         Position    `json:"result"`
	RateLimitReset int         `json:"rate_limit_reset"`
	RateLimit      int         `json:"rate_limit"`
}

type OrderV2 struct {
	UserID        int          `json:"user_id"`
	OrderID       string       `json:"order_id"`
	Symbol        string       `json:"symbol"`
	Side          string       `json:"side"`
	OrderType     string       `json:"order_type"`
	Price         sjson.Number `json:"price"`
	Qty           float64      `json:"qty"`
	TimeInForce   string       `json:"time_in_force"`
	OrderStatus   string       `json:"order_status"`
	LastExecTime  sjson.Number `json:"last_exec_time"`
	LastExecPrice sjson.Number `json:"last_exec_price"`
	LeavesQty     float64      `json:"leaves_qty"`
	CumExecQty    float64      `json:"cum_exec_qty"`
	CumExecValue  sjson.Number `json:"cum_exec_value"`
	CumExecFee    sjson.Number `json:"cum_exec_fee"`
	RejectReason  string       `json:"reject_reason"`
	OrderLinkID   string       `json:"order_link_id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type CreateOrderV2Result struct {
	RetCode          int     `json:"ret_code"`
	RetMsg           string  `json:"ret_msg"`
	ExtCode          string  `json:"ext_code"`
	ExtInfo          string  `json:"ext_info"`
	Result           OrderV2 `json:"result"`
	TimeNow          string  `json:"time_now"`
	RateLimitStatus  int     `json:"rate_limit_status"`
	RateLimitResetMs int64   `json:"rate_limit_reset_ms"`
	RateLimit        int     `json:"rate_limit"`
}

type CancelOrderV2Result struct {
	RetCode          int     `json:"ret_code"`
	RetMsg           string  `json:"ret_msg"`
	ExtCode          string  `json:"ext_code"`
	ExtInfo          string  `json:"ext_info"`
	Result           OrderV2 `json:"result"`
	TimeNow          string  `json:"time_now"`
	RateLimitStatus  int     `json:"rate_limit_status"`
	RateLimitResetMs int64   `json:"rate_limit_reset_ms"`
	RateLimit        int     `json:"rate_limit"`
}

type CancelAllOrderV2Result struct {
	RetCode          int       `json:"ret_code"`
	RetMsg           string    `json:"ret_msg"`
	ExtCode          string    `json:"ext_code"`
	ExtInfo          string    `json:"ext_info"`
	Result           []OrderV2 `json:"result"`
	TimeNow          string    `json:"time_now"`
	RateLimitStatus  int       `json:"rate_limit_status"`
	RateLimitResetMs int64     `json:"rate_limit_reset_ms"`
	RateLimit        int       `json:"rate_limit"`
}

type QueryOrderResult struct {
	RetCode          int     `json:"ret_code"`
	RetMsg           string  `json:"ret_msg"`
	ExtCode          string  `json:"ext_code"`
	ExtInfo          string  `json:"ext_info"`
	Result           OrderV2 `json:"result"`
	TimeNow          string  `json:"time_now"`
	RateLimitStatus  int     `json:"rate_limit_status"`
	RateLimitResetMs int64   `json:"rate_limit_reset_ms"`
	RateLimit        int     `json:"rate_limit"`
}

type StopOrderV2 struct {
	ClOrdID           string       `json:"clOrdID"`
	UserID            int64        `json:"user_id"`
	Symbol            string       `json:"symbol"`
	Side              string       `json:"side"`
	OrderType         string       `json:"order_type"`
	Price             sjson.Number `json:"price"`
	Qty               float64      `json:"qty"`
	TimeInForce       string       `json:"time_in_force"`
	CreateType        string       `json:"create_type"`
	CancelType        string       `json:"cancel_type"`
	OrderStatus       string       `json:"order_status"`
	LeavesQty         float64      `json:"leaves_qty"`
	LeavesValue       string       `json:"leaves_value"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	CrossStatus       string       `json:"cross_status"`
	CrossSeq          float64      `json:"cross_seq"`
	StopOrderType     string       `json:"stop_order_type"`
	TriggerBy         string       `json:"trigger_by"`
	BasePrice         sjson.Number `json:"base_price"`
	ExpectedDirection string       `json:"expected_direction"`
}

type CancelStopOrdersV2Result struct {
	RetCode          int           `json:"ret_code"`
	RetMsg           string        `json:"ret_msg"`
	ExtCode          string        `json:"ext_code"`
	ExtInfo          string        `json:"ext_info"`
	Result           []StopOrderV2 `json:"result"`
	TimeNow          string        `json:"time_now"`
	RateLimitStatus  int           `json:"rate_limit_status"`
	RateLimitResetMs int64         `json:"rate_limit_reset_ms"`
	RateLimit        int           `json:"rate_limit"`
}

type StopOrder struct {
	UserID          int64     `json:"user_id"`
	StopOrderStatus string    `json:"stop_order_status"`
	Symbol          string    `json:"symbol"`
	Side            string    `json:"side"`
	OrderType       string    `json:"order_type"`
	Price           float64   `json:"price"`
	Qty             float64   `json:"qty"`
	TimeInForce     string    `json:"time_in_force"`
	StopOrderType   string    `json:"stop_order_type"`
	TriggerBy       string    `json:"trigger_by"`
	BasePrice       float64   `json:"base_price"`
	OrderLinkID     string    `json:"order_link_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	StopPx          float64   `json:"stop_px"`
	StopOrderID     string    `json:"stop_order_id"`
}

type GetStopOrdersResultData struct {
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	Data        []StopOrder `json:"data"`
}

type GetStopOrdersResult struct {
	RetCode          int                     `json:"ret_code"`
	RetMsg           string                  `json:"ret_msg"`
	ExtCode          string                  `json:"ext_code"`
	Result           GetStopOrdersResultData `json:"result"`
	ExtInfo          interface{}             `json:"ext_info"`
	TimeNow          string                  `json:"time_now"`
	RateLimitStatus  int                     `json:"rate_limit_status"`
	RateLimitResetMs int64                   `json:"rate_limit_reset_ms"`
	RateLimit        int                     `json:"rate_limit"`
}

type OrderBookL2 struct {
	ID     int64   `json:"id"`
	Price  float64 `json:"price,string"`
	Side   string  `json:"side"`
	Size   int64   `json:"size"`
	Symbol string  `json:"symbol"`
}

type OrderBookL2Delta struct {
	Delete []OrderBookL2 `json:"delete"`
	Update []OrderBookL2 `json:"update"`
	Insert []OrderBookL2 `json:"insert"`
}

type KLine struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
	//ID       int64   `json:"id"`        // 563
	//Symbol   string  `json:"symbol"`    // BTCUSD
	//OpenTime int64   `json:"open_time"` // 1539918000
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Turnover float64 `json:"turnover"` // 0.0013844
	//Interval string  `json:"interval"` // 1m
	Confirm   bool  `json:"confirm"`
	Timestamp int64 `json:"timestamp"`
	CrossSeq  int64 `json:"cross_seq"`
}
type KLines struct {
	Data []*KLine `json:"data"`
}

// type KlineV2 struct {
// 	Topic string `json:"topic"`
// 	Data  []struct {
// 		Start     int64   `json:"start"`
// 		End       int64   `json:"end"`
// 		Open      float64 `json:"open"`
// 		Close     float64 `json:"close"`
// 		High      float64 `json:"high"`
// 		Low       float64 `json:"low"`
// 		Volume    float64 `json:"volume"`
// 		Turnover  float64 `json:"turnover"`
// 		Confirm   bool    `json:"confirm"`
// 		CrossSeq  int64   `json:"cross_seq"`
// 		Timestamp int64   `json:"timestamp"`
// 	} `json:"data"`
// 	TimestampE6 int64 `json:"timestamp_e6"`
// }
type KlineV2 struct {
	Topic string `json:"topic"`
	Data  []struct {
		Start     int64   `json:"start"`
		End       int64   `json:"end"`
		Open      float64 `json:"open"`
		Close     float64 `json:"close"`
		High      float64 `json:"high"`
		Low       float64 `json:"low"`
		Volume    float64 `json:"volume"`
		Turnover  float64 `json:"turnover"`
		Confirm   bool    `json:"confirm"`
		CrossSeq  int64   `json:"cross_seq"`
		Timestamp int64   `json:"timestamp"`
	} `json:"data"`
	TimestampE6 int64 `json:"timestamp_e6"`
}

type PositionV3 struct {
	ID                  int       `json:"id"`
	PositionIdx         int       `json:"position_idx"`
	Mode                float64   `json:"mode"`
	UserID              int       `json:"user_id"`
	RiskID              int       `json:"risk_id"`
	Symbol              string    `json:"symbol"`
	Size                float64   `json:"size"`
	Side                string    `json:"side"`
	EntryPrice          float64   `json:"entry_price,string"`
	LiqPrice            float64   `json:"liq_price,string"`
	BustPrice           float64   `json:"bust_price,string"`
	TakeProfit          float64   `json:"take_profit,string"`
	StopLoss            float64   `json:"stop_loss,string"`
	TrailingStop        float64   `json:"trailing_stop,string"`
	PositionValue       float64   `json:"position_value,string"`
	Leverage            float64   `json:"leverage,string"`
	PositionStatus      string    `json:"position_status"`
	AutoAddMargin       float64   `json:"auto_add_margin"`
	OrderMargin         float64   `json:"order_margin,string"`
	PositionMargin      float64   `json:"position_margin,string"`
	OccClosingFee       float64   `json:"occ_closing_fee,string"`
	OccFundingFee       float64   `json:"occ_funding_fee,string"`
	IsIsolated          bool      `json:"is_isolated"`
	WalletBalance       float64   `json:"wallet_balance,string"`
	CumRealisedPnl      float64   `json:"cum_realised_pnl,string"`
	CumCommission       float64   `json:"cum_commission,string"`
	RealisedPnl         float64   `json:"realised_pnl,string"`
	DeleverageIndicator float64   `json:"deleverage_indicator"`
	OcCalcData          string    `json:"oc_calc_data"`
	CrossSeq            float64   `json:"cross_seq"`
	PositionSeq         float64   `json:"position_seq"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	UnrealisedPnl       float64   `json:"unrealised_pnl"`
	TP_SL_Mode          string    `json:"tp_sl_mode"`
}

type ProfitResult struct {
	RetCode          int              `json:"ret_code"`
	RetMsg           string           `json:"ret_msg"`
	ExtCode          string           `json:"ext_code"`
	ExtInfo          string           `json:"ext_info"`
	Result           ProfitResultData `json:"result"`
	TimeNow          string           `json:"time_now"`
	RateLimitStatus  int              `json:"rate_limit_status"`
	RateLimitResetMs int64            `json:"rate_limit_reset_ms"`
	RateLimit        int              `json:"rate_limit"`
}

type ProfitResultData struct {
	CurrentPage int          `json:"current_page"`
	Data        []ProfitData `json:"data"`
}

type ProfitData struct {
	ID            float64 `json:"id"`
	UserID        float64 `json:"user_id"`
	Symbol        string  `json:"symbol"`
	OrderID       string  `json:"order_id"`
	Side          string  `json:"side"`
	Qty           float64 `json:"qty"`
	OrderPrice    float64 `json:"order_price"`
	OrderType     string  `json:"order_type"`
	ExecType      string  `json:"exec_type"`
	ClosedSize    float64 `json:"closed_size"`
	CumEntryValue float64 `json:"cum_entry_value"`
	AvgEntryPrice float64 `json:"avg_entry_price"`
	CumExitValue  float64 `json:"cum_exit_value"`
	AvgExitPrice  float64 `json:"avg_exit_price"`
	ClosedPnl     float64 `json:"closed_pnl"`
	FillCount     float64 `json:"fill_count"`
	Leverage      float64 `json:"leverage"`
	CreatedAt     float64 `json:"created_at"`
}

type TradeRecordResult struct {
	RetCode          int                   `json:"ret_code"`
	RetMsg           string                `json:"ret_msg"`
	ExtCode          string                `json:"ext_code"`
	ExtInfo          string                `json:"ext_info"`
	Result           TradeRecordResultData `json:"result"`
	TimeNow          string                `json:"time_now"`
	RateLimitStatus  int                   `json:"rate_limit_status"`
	RateLimitResetMs int64                 `json:"rate_limit_reset_ms"`
	RateLimit        int                   `json:"rate_limit"`
}

type TradeRecordResultData struct {
	OrderID   string        `json:"order_id"`
	TradeList []TradeRecord `json:"trade_list"`
}

type TradeRecord struct {
	ClosedSize       int    `json:"closed_size"`
	CrossSeq         int    `json:"cross_seq"`
	ExecFee          string `json:"exec_fee"`
	ExecID           string `json:"exec_id"`
	ExecPrice        string `json:"exec_price"`
	ExecQty          int    `json:"exec_qty"`
	ExecTime         string `json:"exec_time"`
	ExecType         string `json:"exec_type"`
	ExecValue        string `json:"exec_value"`
	FeeRate          string `json:"fee_rate"`
	LastLiquidityInd string `json:"last_liquidity_ind"`
	LeavesQty        int    `json:"leaves_qty"`
	NthFill          int    `json:"nth_fill"`
	OrderID          string `json:"order_id"`
	OrderLinkID      string `json:"order_link_id"`
	OrderPrice       string `json:"order_price"`
	OrderQty         int    `json:"order_qty"`
	OrderType        string `json:"order_type"`
	Side             string `json:"side"`
	Symbol           string `json:"symbol"`
	UserID           int    `json:"user_id"`
	TradeTimeMs      int64  `json:"trade_time_ms"`
}

type AutoGenerated struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		Data []struct {
			UserID       int       `json:"user_id"`
			OrderStatus  string    `json:"order_status"`
			Symbol       string    `json:"symbol"`
			Side         string    `json:"side"`
			OrderType    string    `json:"order_type"`
			Price        string    `json:"price"`
			Qty          string    `json:"qty"`
			TimeInForce  string    `json:"time_in_force"`
			OrderLinkID  string    `json:"order_link_id"`
			OrderID      string    `json:"order_id"`
			CreatedAt    time.Time `json:"created_at"`
			UpdatedAt    time.Time `json:"updated_at"`
			LeavesQty    string    `json:"leaves_qty"`
			LeavesValue  string    `json:"leaves_value"`
			CumExecQty   string    `json:"cum_exec_qty"`
			CumExecValue string    `json:"cum_exec_value"`
			CumExecFee   string    `json:"cum_exec_fee"`
			RejectReason string    `json:"reject_reason"`
		} `json:"data"`
		Cursor string `json:"cursor"`
	} `json:"result"`
	TimeNow          string `json:"time_now"`
	RateLimitStatus  int    `json:"rate_limit_status"`
	RateLimitResetMs int64  `json:"rate_limit_reset_ms"`
	RateLimit        int    `json:"rate_limit"`
}

type RealTimeOrderResult struct {
	RetCode          int             `json:"ret_code"`
	RetMsg           string          `json:"ret_msg"`
	ExtCode          string          `json:"ext_code"`
	ExtInfo          string          `json:"ext_info"`
	Result           []RealTimeOrder `json:"result"`
	TimeNow          string          `json:"time_now"`
	RateLimitStatus  int             `json:"rate_limit_status"`
	RateLimitResetMs int64           `json:"rate_limit_reset_ms"`
	RateLimit        int             `json:"rate_limit"`
}

type RealTimeOrder struct {
	UserID      int    `json:"user_id"`
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	OrderType   string `json:"order_type"`
	Price       string `json:"price"`
	Qty         int    `json:"qty"`
	TimeInForce string `json:"time_in_force"`
	OrderStatus string `json:"order_status"`
	ExtFields   struct {
		OReqNum  int    `json:"o_req_num"`
		XreqType string `json:"xreq_type"`
	} `json:"ext_fields"`
	LastExecTime string    `json:"last_exec_time"`
	LeavesQty    int       `json:"leaves_qty"`
	LeavesValue  string    `json:"leaves_value"`
	CumExecQty   int       `json:"cum_exec_qty"`
	CumExecValue string    `json:"cum_exec_value"`
	CumExecFee   string    `json:"cum_exec_fee"`
	RejectReason string    `json:"reject_reason"`
	CancelType   string    `json:"cancel_type"`
	OrderLinkID  string    `json:"order_link_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	OrderID      string    `json:"order_id"`
}
