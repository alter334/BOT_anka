package handler

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
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

	posttext, isAnkaInvoke := h.ankaManager.ankaChecker(p.Message.ChannelID, p.Message.ID)
	if isAnkaInvoke {
		h.BotSimplePost(p.Message.ChannelID, posttext)
	}
	sep := strings.Fields(p.Message.Text)

	if len(sep) == 2 {
		if sep[0] != "@BOT_anka" {
			return
		}
		if sep[1] == "join" {
			log.Println("Received join command")
			h.BotJoiner(p.Message.ChannelID)
		}
		if sep[1] == "leave" {
			log.Println("Received leave command")
			h.BotLeaver(p.Message.ChannelID)
		}
	}

	ankames := []rune(sep[len(sep)-1])

	if ankames[0] != '↓' {
		log.Println(ankames[0])
		return
	}
	amount := string([]rune(ankames)[1:])
	num, err := strconv.Atoi(amount)
	if err != nil {
		log.Println("Failed to parse")
		return
	}

	if num < 1 {
		log.Println("Invalid number")
		return
	}

	ankaID := uuid.New().String()
	newanka := &Anka{
		id:           ankaID,
		messageID:    p.Message.ID,
		channelID:    p.Message.ChannelID,
		messageCount: num,
	}
	h.ankaManager.AddAnka(*newanka)

	log.Println("Add Ancor:" + "after" + strconv.Itoa(num) + ",in:" + channel.Name)

}

// 発火すべき安価があるか確認する
func (am *AnkaManager) ankaChecker(channelid string, messageId string) (string, bool) {
	ankas := am.DecrementAnkaMessageCount(channelid)
	var ankaids []string
	if !(len(ankas) > 0) {
		return "", false
	}
	posttext := ""
	for _, anka := range ankas {
		originUrl := "https://q.trap.jp/messages/" + anka.messageID
		posttext += originUrl + "\n"
		log.Println("Anka/in:", channelid)
		ankaids = append(ankaids, anka.id)
	}
	ancorUrl := "https://q.trap.jp/messages/" + messageId
	// h.BotSimplePost("baaf247d-125a-47e4-82a8-ffcccab5f0b8", originUrl+"\n"+ancorUrl)
	for _, ankaid := range ankaids {
		am.RemoveAnkaByID(ankaid)
		log.Println("Remove Anka:" + ankaid + ",in:" + channelid)

	}

	return posttext + ancorUrl, true
}
