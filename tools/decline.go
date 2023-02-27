package tools

import (
	"log"
	"net/http"

	_ "github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Decline(req *http.Request, bot *linebot.Client, event *linebot.Event) {

	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ご検討ありがとうございました。またよろしくお願いします。")).Do(); err != nil {
		log.Print(err)
	}

}
