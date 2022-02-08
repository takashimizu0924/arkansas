package slack

import (
	"arkansas/bybit"
	"arkansas/config"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/slack-go/slack"
)

func NewSlack() *slack.Client {
	bot := slack.New(config.Config.SlackToken)
	return bot
}

func PostTextMessage(bot *slack.Client, signals []bybit.Signal) {
	var sb strings.Builder
	sb.WriteString("【 シグナルでました Ver1.1 】\n ")
	for i := 0; i < len(signals); i++ {
		sb.WriteString(" [ ")
		sb.WriteString(signals[i].CoinPair)
		sb.WriteString(" ] :")
		sb.WriteString(signals[i].Type)
		if signals[i].Sign == 0 {
			sb.WriteString("での下降予想です\n")
			sb.WriteString("現在価格-->")
			sb.WriteString(strconv.FormatFloat(signals[i].Price, 'f', 5, 64))
			sb.WriteString("\n")
		}
		if signals[i].Sign == 1 {
			sb.WriteString("での上昇予想です\n")
			sb.WriteString("現在価格-->")
			sb.WriteString(strconv.FormatFloat(signals[i].Price, 'f', 5, 64))
			sb.WriteString("\n")
		}
	}
	_, _, err := bot.PostMessage(config.Config.SlackChannel, slack.MsgOptionText(sb.String(), true))
	if err != nil {
		log.Println(err)
	}
}

//チャンネルを監視し、それぞれのチャンネルから受信した内容を
func slackBotV1(bot *slack.Client, signals <-chan []bybit.Signal) {
	rtm := bot.NewRTM()

	go rtm.ManageConnection()
	for msg := range rtm.IncomingEvents {
		switch event := msg.Data.(type) {
		case *slack.MessageEvent:
			if fmt.Sprintf("%v", event.ClientMsgID) == "" {
				break
			}
		}
	}
}

/*-----------------------------------------------------------------------------*/
//メッセージの中身を確認
func validateMessageEvent(event *slack.Message) {

}

/*-----------------------------------------------------------------------------*/
//Slackで擬似的に売買を行う
//シグナルに従って、現在の価格で購入（実際は現在価格＊1.0025）を行い、Slackで通知する
//配列に一時保管する。配列に入れる際は、同じ通貨があるかを確認し、ある場合は保管しない。
//配列に入っている価格は常時監視する。売買価格,もしくは損切り価格（購入価格 / 1.004）になったらSlackで通知する
func TradingPush(bot *slack.Client, tradeCh <-chan bybit.Signal) {
	var sb strings.Builder
	for t := range tradeCh {
		if t.Sign == 0 {
			sb.WriteString("売り決済完了\n")
			sb.WriteString("[ ")
			sb.WriteString(t.CoinPair)
			sb.WriteString(" ]\n")
			sb.WriteString("利益は:")
			sb.WriteString(strconv.FormatFloat(t.Price, 'f', 5, 64))
		}
		if t.Sign == 1 {
			sb.WriteString("買い決済完了\n")
			sb.WriteString("[ ")
			sb.WriteString(t.CoinPair)
			sb.WriteString(" ]\n")
			sb.WriteString("利益は:")
			sb.WriteString(strconv.FormatFloat(t.Price, 'f', 5, 64))
		}
		_, _, err := bot.PostMessage(config.Config.SlackChannel, slack.MsgOptionText(sb.String(), true))
		if err != nil {
			log.Println(err)
		}
	}
}
