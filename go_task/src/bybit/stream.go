package bybit

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

type WsKlineResult struct {
	Topic string `json:"topic"`
	Data  []struct {
		Start     float64 `json:"start"`
		End       float64 `json:"end"`
		Open      float64 `json:"open"`
		Close     float64 `json:"close"`
		High      float64 `json:"high"`
		Low       float64 `json:"low"`
		Volume    float64 `json:"volume"`
		Turnover  float64 `json:"turnover"`
		Confirm   bool    `json:"confirm"`
		CrossSeq  float64 `json:"cross_seq"`
		Timestamp int64   `json:"timestamp"`
	} `json:"data"`
	TimestampE6 int64 `json:"timestamp_e6"`
}

func RealTimeKline(op string, argsParam ...string) {
	u := url.URL{Scheme: "wss", Host: "stream.bybit.com", Path: "/realtime"}
	log.Printf("Connecting to....%s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println(err)
	}
	defer c.Close()
	param := map[string]interface{}{}
	//ws.send('{"op":"subscribe","args":["klineV2.1.BTCUSD"]}')

	var args []string
	topic := "klineV2.5."
	coinpair := strings.Join(argsParam, "|")
	args = append(args, topic+coinpair)
	param["op"] = op
	param["args"] = args
	log.Println(param)
	byteParam, err := json.Marshal(param)
	c.WriteMessage(websocket.TextMessage, byteParam)

	for {
		_, messsage, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		// log.Println(string(messsage))
		var wsKlines WsKlineResult
		if err := json.Unmarshal(messsage, &wsKlines); err != nil {
			log.Println(err)
		}
		log.Println(wsKlines)

	}
}

//"request":{"op":"subscribe","args":["candle.5.BTCUSDT"]}}

//"request":{"op":"subscribe","args":["klineV2.1.BTCUSD"]}}
