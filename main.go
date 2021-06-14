package main

import (
	"github.com/Valeviy/fiziklbot/telegram"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/telegram-bot-api.v4"

	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
)

var (
	bot *tgbotapi.BotAPI
	err error
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT must be set")
	}

	// gin router
	router := gin.New()
	router.Use(gin.Logger())

	// telegram
	bot, err = telegram.InitTelegram()
	if err != nil {
		log.Fatal("Bot not initialised")
	}

	router.POST("/"+bot.Token, telegram.WebhookHandler)

	err = router.Run(":" + port)

	if err != nil {
		log.Println(err)
	}
}
