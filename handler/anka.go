package handler

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/traPtitech/traq-ws-bot/payload"
)

// 安価登録
func (h *Handler) ankaProcessor(p *payload.MessageCreated) {
	log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
	channel, _, _ := h.bot.API().ChannelApi.GetChannel(context.Background(), p.Message.ChannelID).Execute()

	if _, exist := h.messageCount[p.Message.ChannelID]; !exist {
		h.messageCount[p.Message.ChannelID] = 0
		log.Println(h.messageCount[p.Message.ChannelID], ":"+channel.Name)
	} else {
		h.messageCount[p.Message.ChannelID]++
		log.Println(h.messageCount[p.Message.ChannelID], ":"+channel.Name)
	}

	h.ankaChecker(p.Message.ChannelID, h.messageCount[p.Message.ChannelID], p.Message.ID)
	sep := strings.Fields(p.Message.Text)

	if len(sep) == 2 {
		if sep[1] == "join" {
			log.Println("Received join command")
			h.BotJoiner(p.Message.ChannelID)
		}
		if sep[1] == "leave" {
			log.Println("Received leave command")
			h.BotLeaver(p.Message.ChannelID)
		}
	}

	anka := []rune(sep[len(sep)-1])

	if anka[0] != '↓' {
		log.Println(anka[0])
		return
	}
	amount := string([]rune(anka)[1:])
	num, err := strconv.Atoi(amount)
	if err != nil {
		log.Println("Failed to parse")
		return
	}
	if _, exist := h.ankas[p.Message.ChannelID]; !exist {
		h.ankas[p.Message.ChannelID] = make(map[int]string)
	}
	h.ankas[p.Message.ChannelID][h.messageCount[p.Message.ChannelID]+num] = p.Message.ID

	log.Println("Add Ancor:" + strconv.Itoa(h.messageCount[p.Message.ChannelID]+num) + ",in:" + channel.Name)

}

func (h *Handler) ankaChecker(channelid string, messageNum int, messageId string) {
	originID, exist := h.ankas[channelid][messageNum]
	if !exist {
		return
	}
	originUrl := "https://q.trap.jp/messages/" + originID
	ancorUrl := "https://q.trap.jp/messages/" + messageId
	h.BotSimplePost(channelid, originUrl+"\n"+ancorUrl)
	h.BotSimplePost("baaf247d-125a-47e4-82a8-ffcccab5f0b8", originUrl+"\n"+ancorUrl)
	log.Println("Anka/in:", channelid, " at:", messageNum)
	delete(h.ankas[channelid], messageNum)
}
