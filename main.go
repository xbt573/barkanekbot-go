package main

import (
	"log"
	"os"
	"time"

	"github.com/xbt573/barkanekbot-go/handlers"
	"gopkg.in/telebot.v3"
)

func main() {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	handlers.Router(bot)

	bot.Start()
}
