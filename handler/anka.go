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

	posttext, isAnkaInvoke := h.ankaManager.ankaChecker(p.Message.ChannelID, p.Message.ID, p.Message.Text, h)
	if isAnkaInvoke {
		h.BotSimplePost(p.Message.ChannelID, posttext)
		h.BotSimplePost("baaf247d-125a-47e4-82a8-ffcccab5f0b8", posttext)
	}
	sep := strings.Fields(p.Message.Text)

	if len(sep) == 2 {
		if sep[0] == "@BOT_anka" {

			if sep[1] == "join" {
				log.Println("Received join command")
				h.BotJoiner(p.Message.ChannelID)
				return
			}
			if sep[1] == "leave" {
				log.Println("Received leave command")
				h.BotLeaver(p.Message.ChannelID)
				return
			}
		}
	}

	viewMessage := h.ankaManager.AnkaReader(p)

	posttext, isMessageCreate := viewMessage.ViewMessageMaker()
	if isMessageCreate {
		viewMessage.id = h.BotSimplePost(p.Message.ChannelID, posttext)
	}

}

// 発火すべき安価があるか確認する
func (am *AnkaManager) ankaChecker(channelid string, messageId string, messageText string, h *Handler) (string, bool) {
	ankas := am.DecrementAnkaMessageCount(channelid)
	var ankaids []string
	if !(len(ankas) > 0) {
		return "", false
	}
	posttext := ""
	ancorUrl := "https://q.trap.jp/messages/" + messageId
	for _, anka := range ankas {
		originUrl := "https://q.trap.jp/messages/" + anka.messageID
		posttext += originUrl + "\n"
		log.Println("Anka/in:", channelid)
		ankaids = append(ankaids, anka.id)
		replaceText := "[(↓" + strconv.Itoa(anka.originmessageCount) + ")" + messageText + "](" + ancorUrl + ")"
		if anka.viewMessage == nil {
			log.Println("viewMessageisNil (Maybe because of reset of Application)")
			continue
		}
		anka.viewMessage.openedText[anka.inMessageAnkaOrder] = replaceText
		updateAnkaViewText, isUpdate := anka.viewMessage.ViewMessageMaker()
		if !isUpdate {
			continue
		}
		h.BotSimpleUpdate(anka.viewMessage.id, updateAnkaViewText)

	}

	for _, ankaid := range ankaids {
		am.RemoveAnkaByID(ankaid)
		log.Println("Remove Anka:" + ankaid + ",in:" + channelid)

	}

	return posttext + ancorUrl, true
}
