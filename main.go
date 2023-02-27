package main

import (
	"fmt"
	"log"

	// "log"

	"net/http"

	"example.com/m/tools"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	_ "gorm.io/gorm"
)

type UserProfile struct {
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

func main() {
	// gin.DefaultWriter = colorable.NewColorableStdout()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Serve static files
	r.Static("/static", "./static")

	// push //
	// tools.Test()

	bot, err := linebot.New("07b57cd44f406d7426002ddde26f8510", "ckF2YVNKOFIHq/z83vtF+CspJaiEFTO/mAeBzN+FvY2LUb6OzlDy56nOQUv8GfZrZLiik9YZxHnCEhCbq/PHM8JY5pekcYGDcB2wN2h6oCGUub5Pv5ijC1CK+osVCpgvi1SavlLseym2rxC6vlHQ3QdB04t89/1O/w1cDnyilFU=")
	if _, err := bot.PushMessage("Udc51e4d955763ca7cae7e5c9ed5e8bde", linebot.NewTextMessage("https://ourcargo-platform.com/orderpage")).Do(); err != nil {
		fmt.Println(err.Error())
	}

	// OrderPage //
	r.GET("/orderpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "order.html", gin.H{})
	})

	r.POST("/callback", func(c *gin.Context) {
		if err != nil {
			fmt.Println(err.Error())
			// c.HTML(http.StatusOK, gin.H{"error": "invalid argument"})
		}

		req := c.Request
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.Writer.WriteHeader(400)
			} else {
				c.Writer.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {

				case *linebot.TextMessage:
					userID := event.Source.UserID

					if message.Text == "accept" {
						tools.Accept(req, bot, event)
						fmt.Println(userID)

					} else if message.Text == "decline" {
						tools.Decline(req, bot, event)
						fmt.Println(userID)

					} else {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("お確かめの上、もう一度ご回答ください。")).Do(); err != nil {
							log.Print(err)
							fmt.Println(userID)
						}
					}
				}
			}
		}

	})
	// http.HandleFunc("/callback/user", func(w http.ResponseWriter, req *http.Request) {
	// 	....
	// })

	// })

	fmt.Println("server is running at port 443")

	// srv := &http.Server{
	// 	Addr:         "0.0.0.0:" + os.Getenv("PORT"),
	// 	Handler:      router,
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	// client := gorequest.New()

	if err := r.RunTLS(
		":443",
		"/etc/letsencrypt/live/ourcargo-platform.com/fullchain.pem",
		"/etc/letsencrypt/live/ourcargo-platform.com/privkey.pem",
	); err != nil {
		fmt.Println(err.Error())
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	// http.ListenAndServe(":80", nil)
}

// func getProfile(userID string) (*http.Response, []error) {
// 	request := client.Get("https://api.line.me/v2/bot/profile").
// 		Query("user_id="+userID).
// 		Set("Authorization", "Bearer "+"ckF2YVNKOFIHq/z83vtF+CspJaiEFTO/mAeBzN+FvY2LUb6OzlDy56nOQUv8GfZrZLiik9YZxHnCEhCbq/PHM8JY5pekcYGDcB2wN2h6oCGUub5Pv5ijC1CK+osVCpgvi1SavlLseym2rxC6vlHQ3QdB04t89/1O/w1cDnyilFU=")
// 	response, body, errs := request.End()

// 	return response, errs
// }

// func parseProfile(response *http.Response) (*UserProfile, error) {
// 	userProfile := &UserProfile{}
// 	err := json.NewDecoder(response.Body).Decode(userProfile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return userProfile, nil
// }
