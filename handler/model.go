package handler

import (
	db "bot_anka/DB"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

type Handler struct {
	bot         *traqwsbot.Bot
	ankaManager *AnkaManager // 安価管理
}

type AnkaManager struct {
	ankas []Anka // 安価配列
	db    *db.DB // DB
}

type Anka struct {
	id                 string
	messageID          string
	channelID          string
	messageCount       int
	originmessageCount int
	inMessageAnkaOrder int
	viewMessage        *AnkaViewMessage
}

type AnkaViewMessage struct {
	id         string
	originText []string
	openedText []string
}
