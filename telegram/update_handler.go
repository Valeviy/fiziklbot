package telegram

import (
	"github.com/Valeviy/fiziklbot/telegram/commands"
	"github.com/Valeviy/fiziklbot/telegram/messages"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"reflect"
)

func UpdateHandler(u *tgbotapi.Update) tgbotapi.MessageConfig {
	var msg = tgbotapi.NewMessage(u.Message.Chat.ID, "")

	if u.Message != nil {
		if u.Message.IsCommand() {
			CommandHandler(u, &msg)
		} else {
			TextHandler(u, &msg)
		}
	} else {
		msg.Text = messages.EMPTY
	}

	return msg
}

func CommandHandler(u *tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	switch u.Message.Command() {
	case commands.START:
		msg.Text = messages.START
	default:
		msg.Text = messages.EMPTY
	}
}

func TextHandler(u *tgbotapi.Update, msg *tgbotapi.MessageConfig) {
	if reflect.TypeOf(u.Message.Text).Kind() == reflect.String && u.Message.Text != "" {
		msg.Text = u.Message.Text
		log.Printf("From: %+v Text: %+v\n", u.Message.From, u.Message.Text)
	} else {
		msg.Text = messages.EMPTY
	}
}
