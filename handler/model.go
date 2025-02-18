package handler

import traqwsbot "github.com/traPtitech/traq-ws-bot"

type Handler struct {
	bot          *traqwsbot.Bot
	messageCount map[string]int // 安価管理用の累積メッセージ数
	ankas        map[string](map[int]string) // 安価管理
	ankaManager AnkaManager // 安価管理
}

type AnkaManager struct {
	ankas []Anka // 安価配列
}

type Anka struct {
	id string
	messageID string
	channelID string
	messageCount int
	delete func(am *AnkaManager) // 削除コールバック
}

