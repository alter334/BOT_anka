package main

import (
	"bot_anka/handler"
	"os"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

func main() {

	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: os.Getenv("TRAQ_BOT_TOKEN"), // Required
	})
	if err != nil {
		panic(err)
	}
	h := handler.NewHandler(bot)
	h.BotHandler()

}
