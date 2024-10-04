package handler

import (
	"log"
	"time"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func NewHandler(bot *traqwsbot.Bot) *Handler {
	return &Handler{bot: bot, messageCount: make(map[string]int), ankas: make(map[int]string)}
}

func (h *Handler) BotHandler() {
	log.Println(time.Now())
	h.bot.OnJoined(func(p *payload.Joined) {
		log.Println("Joined:" + p.Channel.ID)
		h.BotSimplePost(p.Channel.ID, "Joined")
	})
	h.bot.OnLeft(func(p *payload.Left) {
		log.Println("Left:" + p.Channel.ID)
		h.BotSimplePost(p.Channel.ID, "Left")
	})
	h.bot.OnMessageCreated(h.ankaProcessor)

	if err := h.bot.Start(); err != nil {
		panic(err)
	}

}
