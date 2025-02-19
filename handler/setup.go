package handler

import (
	db "bot_anka/DB"
	"log"
	"time"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func NewHandler(bot *traqwsbot.Bot) *Handler {
	newDB := &db.DB{}
	newDB.Setup()
	manager := &AnkaManager{ankas: make([]Anka, 0), db: newDB}
	manager.ManagerSetupFromDB()
	return &Handler{bot: bot, messageCount: make(map[string]int), ankas: make(map[string](map[int]string)), ankaManager: manager}
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
