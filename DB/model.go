package db

import "log"

type AnkaDBData struct {
	Id                     string `db:"Id"`
	OriginMessageId        string `db:"originMessageId"`
	ChannelId              string `db:"channelId"`
	AnkaInvokeMessageCount int    `db:"ankaInvokemessageCount"`
}

func (d *DB) AnkaInserttoDB(ankaData *AnkaDBData) (err error) {

	_, err = d.dataBase.Exec("INSERT INTO `anka` (`Id`, `originMessageId`, `channelId`, `ankaInvokeMessageCount`) VALUES (?,?,?,?)",
		ankaData.Id, ankaData.OriginMessageId, ankaData.ChannelId, ankaData.AnkaInvokeMessageCount)
	if err != nil {
		log.Println("ankaInsertError:", err.Error())
		return err
	}

	return nil

}
