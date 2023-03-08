package main

import (
	"fmt"
	"log"
	"time"

	// "log"

	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/ourcargo/constant"
	"github.com/ourcargo/dao"
	"github.com/ourcargo/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type OrderContent struct {
	UserID          int    `json:"UserID"`
	TruckSizeID     int    `json:"TruckSize"`
	PickupDatetime  string `json:"PickupDatetime"`
	ArrivalDatetime string `json:"ArrivalDatetime"`
	PickupLocation  string `json:"PickupLocation"`
	ArrivalLocation string `json:"ArrivalLocation"`
	OrderDatetime   string `json:"OrderDatetime"`
	Price           string `json:"Price"`
}

var db *gorm.DB
var err error

func main() {
	// gin.DefaultWriter = colorable.NewColorableStdout()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	dsn := "host=myrds.c8eoe8ahfumy.ap-northeast-1.rds.amazonaws.com user=postgres password=Hach8686 dbname=test port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	bot, err := linebot.New("07b57cd44f406d7426002ddde26f8510", "ckF2YVNKOFIHq/z83vtF+CspJaiEFTO/mAeBzN+FvY2LUb6OzlDy56nOQUv8GfZrZLiik9YZxHnCEhCbq/PHM8JY5pekcYGDcB2wN2h6oCGUub5Pv5ijC1CK+osVCpgvi1SavlLseym2rxC6vlHQ3QdB04t89/1O/w1cDnyilFU=")
	// if _, err := bot.PushMessage("Udc51e4d955763ca7cae7e5c9ed5e8bde", linebot.NewTextMessage("https://ourcargo-platform.com/orderpage")).Do(); err != nil {
	// 	fmt.Println(err.Error())
	// }

	// OrderPage //
	r.GET("/orderpage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "order.html", gin.H{
			"TruckSizeSmall":   constant.TruckSizeSmall,
			"TruckSizeTowTon":  constant.TruckSizeTowTon,
			"TruckSizeFourTon": constant.TruckSizeFourTon,
			"TruckSizeTenTon":  constant.TruckSizeTenTon,
		})
	})

	// form を送る
	r.POST("/orderpage", func(c *gin.Context) {
		var order OrderContent
		if err := c.ShouldBind(&order); err != nil {
			fmt.Print(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument"})
			return
		}

		Layout := "2006-01-02T15:04:00.000Z"
		Layout_day := "2006-01-02"
		Pt, err := time.Parse(Layout, order.PickupDatetime)
		At, err := time.Parse(Layout, order.ArrivalDatetime)
		Ot, err := time.Parse(Layout_day, order.OrderDatetime)

		newRecord := dao.Order{
			UserID:          order.UserID,
			PickupDatetime:  Pt,
			ArrivalDatetime: At,
			OrderDatetime:   Ot,
			PickupLocation:  order.PickupLocation,
			ArrivalLocation: order.ArrivalLocation,
			TruckSizeID:     order.TruckSizeID,
			Price:           order.Price,
		}

		if err != nil {
			fmt.Println(err)
		}

		dbc := db.Create(&newRecord)
		if dbc.Error != nil {
			fmt.Print(dbc.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// DBへの保存が成功したら結果を表示するページに戻るために/showpageのAPIを内部で読んでそちらでページの表示を行う。
		location := url.URL{Path: "/"}
		c.Redirect(http.StatusFound, location.RequestURI())
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

	fmt.Println("server is running at port 443")

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

// 	return userProfile, nil
// }
