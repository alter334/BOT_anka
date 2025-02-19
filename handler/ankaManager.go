package handler

import (
	db "bot_anka/DB"

	"github.com/labstack/gommon/log"
)

// Ankaを追加するメソッド
func (am *AnkaManager) AddAnka(a Anka) {
	am.ankas = append(am.ankas, a)
	am.addAnkatoDB(&a)
}

// AnkaをIDで検索するメソッド
func (am *AnkaManager) FindAnkaByID(id string) (*Anka, bool) {
	for _, a := range am.ankas {
		if a.id == id {
			return &a, true
		}
	}
	return nil, false
}

// AnkaをIDで削除するメソッド
func (am *AnkaManager) RemoveAnkaByID(id string) (*Anka, bool) {
	for i, a := range am.ankas {
		if a.id == id {
			am.ankas = append(am.ankas[:i], am.ankas[i+1:]...)
			return &a, true
		}
	}
	return nil, false
}

// AnkaをChannelIDで検索するメソッド
func (am *AnkaManager) FindAnkaByChannelID(channelID string) ([]*Anka, bool) {
	var result []*Anka
	for _, a := range am.ankas {
		if a.channelID == channelID {
			result = append(result, &a)
		}
	}
	return result, len(result) > 0
}

// 引数のChannelIDのAnkaのmessageCountを全て1減らす
// messageCountが0になったAnkaを返す
func (am *AnkaManager) DecrementAnkaMessageCount(channelID string) []*Anka {
	var result []*Anka
	log.Printf("AnkaManager")
	err := am.db.AnkaNumDecrementByChannel(channelID)
	if err != nil {
		log.Printf("AnkaDeleteError", err.Error())
		return result
	}
	err = am.db.AnkaDeleteFromDB()
	if err != nil {
		log.Printf("AnkaDeleteError", err.Error())
		return result
	}
	for i := range am.ankas {
		if am.ankas[i].channelID == channelID {
			am.ankas[i].messageCount--
			log.Printf("AnkaNumDecrare:id=%v,num=%v", am.ankas[i].id, am.ankas[i].messageCount)
			if am.ankas[i].messageCount == 0 {
				result = append(result, &am.ankas[i])
			}
		}
	}
	return result
}

// ankaをDBに追加するメソッド

func (am *AnkaManager) addAnkatoDB(anka *Anka) {
	data := am.ankaDataConverttoDBdata(anka)
	am.db.AnkaInserttoDB(data)
}

func (am *AnkaManager) ankaDataConverttoDBdata(anka *Anka) *db.AnkaDBData {
	return &db.AnkaDBData{Id: anka.id, OriginMessageId: anka.messageID, ChannelId: anka.channelID, AnkaInvokeMessageCount: anka.messageCount}
}
