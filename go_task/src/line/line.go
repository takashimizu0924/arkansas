package line

import (
	"arkansas/config"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

const timeLayout = "2006-01-02 15:04:05"

func NewLine() (*linebot.Client, error) {
	bot, err := linebot.New(config.Config.LineSecret, config.Config.LineToken)
	if err != nil {
		log.Printf("action=NewLine err=%s", err)
	}
	return bot, err
}

func PostTextMessage(date time.Time, profit float64, bot *linebot.Client) error {
	var sb strings.Builder
	strTime := timeToString(date)
	strProfit := strconv.FormatFloat(profit, 'f', 2, 64)
	sb.WriteString("【bitflyer 自動売買 収益】\n")
	sb.WriteString("date / profit /\n")
	sb.WriteString(strTime + "/" + strProfit + "/")
	message := linebot.NewTextMessage(sb.String())
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Println("action=PostTextMessage err=:", err)
		return err
	}
	return nil
}

func timeToString(t time.Time) string {
	str := t.Format(timeLayout)
	return str
}
