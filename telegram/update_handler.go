package telegram

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

func UpdateHandler(u *tgbotapi.Update) {
	log.Printf("From: %+v Text: %+v\n", u.Message.From, u.Message.Text)
}
