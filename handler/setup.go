package handler

import (
	"log"
	"time"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

func NewHandler(bot *traqwsbot.Bot) *Handler {
	return &Handler{bot: bot}
}

func (h *Handler) BotHandler() {
	log.Println(time.Now())
	h.bot.OnMessageCreated(h.ankaProcessor)

	if err := h.bot.Start(); err != nil {
		panic(err)
	}

}
