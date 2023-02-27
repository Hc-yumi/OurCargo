package tools

import (
	"log"
	"net/http"

	_ "github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// func Accept(w http.ResponseWriter, req *http.Request, bot *linebot.Client, event *linebot.Event) {
func Accept(req *http.Request, bot *linebot.Client, event *linebot.Event) {

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("御引き受けありがとうございます。")).Do(); err != nil {
		log.Print(err)
	}

	// if _, err := bot.PushMessage("Udc51e4d955763ca7cae7e5c9ed5e8bde", linebot.NewTextMessage("accept or decline?")).Do(); err != nil {
	// 	fmt.Println(err.Error())
	// }

}
