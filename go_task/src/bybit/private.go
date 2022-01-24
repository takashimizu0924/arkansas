package bybit

import (
	"fmt"
	"log"
	"net/http"
)

func (b *Bybit) CreateOrder(side, symbol, orderType string, qty int, price float64, timeInForce string, closeOnTrigger bool, orderLinkID string, takeProfit float64, stopLoss float64, reduceOnly bool) (result OrderV2, err error) {
	var ret CancelOrderV2Result
	params := map[string]interface{}{}
	params["side"] = side
	params["symbol"] = symbol
	params["order_type"] = orderType
	params["qty"] = qty
	params["time_in_force"] = timeInForce
	if price > 0 {
		params["price"] = price
	}
	if closeOnTrigger {
		params["close_on_trigger"] = true
	}
	if orderLinkID != "" {
		params["order_link_id"] = orderLinkID
	}
	if takeProfit > 0 {
		params["take_profit"] = takeProfit
	}
	if stopLoss > 0 {
		params["stop_loss"] = stopLoss
	}
	if reduceOnly {
		params["reduce_only"] = true
	}

	resp, err := b.SignedRequest(http.MethodPost, "/v2/private/order/create", params, &ret)
	if err != nil {
		log.Println("action=CreateOrder ==>", err.Error())
		return
	}
	if ret.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", ret.RetMsg, string(resp))
		return
	}
	result = ret.Result
	return
}

func (b *Bybit) GetOrder(symbol string, orderStatus string, direction string, limit int) (result []OrderV3, err error) {
	var ret OrderListResult

	if limit == 0 {
		limit = 20
	}
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderStatus != "" {
		params["order_status"] = orderStatus
	}
	if direction != "" {
		params["direction"] = direction
	}
	params["limit"] = limit

	resp, err := b.SignedRequest(http.MethodGet, "/v2/private/order/list", params, &ret)
	if err != nil {
		log.Println("action=GetOrder ==>", err.Error())
		return
	}
	if ret.RetCode != 0 {
		err = fmt.Errorf("action=GetOrder ==>%v : %v", ret.RetMsg, string(resp))
		return
	}
	fmt.Println("resp", string(resp))
	result = ret.Result.Data
	return
}

func (b *Bybit) GetRealTimeOrder(symbol string) (result []RealTimeOrder, err error) {
	var ret RealTimeOrderResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	resp, err := b.SignedRequest(http.MethodGet, "/v2/private/order", params, &ret)
	if err != nil {
		log.Println(err)
		return
	}
	if ret.RetCode != 0 {
		err = fmt.Errorf("action=GetOrder ==>%v : %v", ret.RetMsg, string(resp))
		return
	}
	result = ret.Result
	return
}

func (b *Bybit) CancelOrder(orderID string, orderLinkID string, symbol string) (result OrderV2, err error) {
	var ret CancelOrderV2Result
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderID != "" {
		params["order_id"] = orderID
	}
	if orderLinkID != "" {
		params["order_link_id"] = orderLinkID
	}
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "/v2/private/order/cancel", params, &ret)
	if err != nil {
		return
	}
	if ret.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", ret.RetMsg, string(resp))
		return
	}

	result = ret.Result
	return
}

func (b *Bybit) CancelAllOrder(symbol string) (result []OrderV2, err error) {
	var ret CancelAllOrderV2Result
	params := map[string]interface{}{}
	params["symbol"] = symbol
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "v2/private/order/cancelAll", params, &ret)
	if err != nil {
		return
	}
	if ret.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", ret.RetMsg, string(resp))
		return
	}

	result = ret.Result
	return
}

func (b *Bybit) GetPositions(symbol string) (result PositionV3, err error) {
	var ret PositionListResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	resp, err := b.SignedRequest(http.MethodGet, "/v2/private/position/list", params, &ret)
	if err != nil {
		log.Println(err)
	}
	if ret.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", ret.RetMsg, string(resp))
		return
	}
	result = ret.Result

	return
}

func (b *Bybit) convertPositionV1(position PositionV1) (result Position) {
	result.ID = position.ID
	result.UserID = position.UserID
	result.RiskID = position.RiskID
	result.Symbol = position.Symbol
	result.Size = position.Size
	result.Side = position.Side
	result.EntryPrice = position.EntryPrice
	result.LiqPrice = position.LiqPrice
	result.BustPrice = position.BustPrice
	result.TakeProfit = position.TakeProfit
	result.StopLoss = position.StopLoss
	result.TrailingStop = position.TrailingStop
	result.PositionValue = position.PositionValue
	result.Leverage = position.Leverage
	result.PositionStatus = position.PositionStatus
	result.AutoAddMargin = position.AutoAddMargin
	result.OrderMargin = position.OrderMargin
	result.PositionMargin = position.PositionMargin
	result.OccClosingFee = position.OccClosingFee
	result.OccFundingFee = position.OccFundingFee
	result.ExtFields = position.ExtFields
	result.WalletBalance = position.WalletBalance
	result.CumRealisedPnl = position.CumRealisedPnl
	result.CumCommission = position.CumCommission
	result.RealisedPnl = position.RealisedPnl
	result.DeleverageIndicator = position.DeleverageIndicator
	result.OcCalcData = position.OcCalcData
	result.CrossSeq = position.CrossSeq
	result.PositionSeq = position.PositionSeq
	result.CreatedAt = position.CreatedAt
	result.UpdatedAt = position.UpdatedAt
	result.UnrealisedPnl = position.UnrealisedPnl
	return
}

func (b *Bybit) Profit(symbol string, end int64, from int64) (err error) {
	var ret ProfitResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["start_time"] = from
	params["end_time"] = end
	resp, err := b.SignedRequest(http.MethodGet, "/v2/private/trade/closed-pnl/list", params, &ret)
	if err != nil {
		log.Println(err)
	}
	if ret.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", ret.RetMsg, string(resp))
		return
	}
	profitList := ret.Result.Data
	totalProfit := 0.0
	for _, profit := range profitList {
		fmt.Printf("profit %.8f\n", profit.ClosedPnl)
		totalProfit += profit.ClosedPnl
	}
	fmt.Printf("totalProfit %.8f\n", totalProfit)

	return nil
}
