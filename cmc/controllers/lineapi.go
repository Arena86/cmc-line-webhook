package controllers

import (
	"cmc/cmc/databases/conn"
	libs "cmc/cmc/helpers"
	"cmc/cmc/models"
	"fmt"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/gin-gonic/gin"
)

func WebHook(c *gin.Context) {
	defer libs.RecoverError(c)
	var (
		status = 200
		msg    string
	)
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err == nil {
		events, err := bot.ParseRequest(c.Request)
		if err == nil {
			for _, event := range events {
				var lineEvent models.LineEvent
				if event.Type == linebot.EventTypeMessage {
					switch message := event.Message.(type) {
					case *linebot.TextMessage:
						{
							lineEvent.Event = string(linebot.MessageTypeText)
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()
							if err != nil {
								fmt.Println("err: ", err)
							}
						}
					case *linebot.StickerMessage:
						{
							lineEvent.Event = string(linebot.MessageTypeSticker)
							replyMessage := fmt.Sprintf(
								"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
							_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
							if err != nil {
								fmt.Println("err: ", err)
							}
						}
					}
				}
				if lineEvent.Event != "" {
					db := conn.Connect()
					if db != nil {
						resultCreate := db.Create(&lineEvent)
						if resultCreate.Error != nil {
							fmt.Println("resultCreate > Error: ", resultCreate.Error)
						}
					}
				}
			}
		} else {
			status = 403
			msg = err.Error()
		}
	} else {
		status = 403
		msg = err.Error()
	}
	if status == 200 {
		msg = "Success"
	}
	responseData := gin.H{
		"status": status,
		"msg":    msg,
	}
	libs.APIResponseData(c, status, responseData)
	return
}
