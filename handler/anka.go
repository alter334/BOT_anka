package handler

import (
	"log"
	"strconv"
	"strings"

	"github.com/traPtitech/traq-ws-bot/payload"
)

// 安価登録
func (h *Handler) ankaProcessor(p *payload.MessageCreated) {
	log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
	h.messageCount++
	h.ankaChecker(p.Message.ChannelID, h.messageCount, p.Message.ID)
	sep := strings.Fields(p.Message.Text)
	anka := sep[len(sep)-1]

	if rune(anka[0]) != '↓' {
		return
	}
	amount := string([]rune(anka)[1:])
	num, err := strconv.Atoi(amount)
	if err != nil {
		log.Println("Failed to parse")
		return
	}
	h.ankas[h.messageCount+num] = p.Message.ID
	log.Println("Add Ancor:" + strconv.Itoa(h.messageCount+num))

}

func (h *Handler) ankaChecker(channelid string, messageNum int, messageId string) {
	originID, exist := h.ankas[messageNum]
	if !exist {
		return
	}
	originUrl := "https://q.trap.jp/messages/" + originID
	ancorUrl := "https://q.trap.jp/messages/" + messageId
	h.BotSimplePost(channelid, originUrl+"\n"+ancorUrl)
	delete(h.ankas, messageNum)
}
