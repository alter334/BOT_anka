package db

import "log"

type AnkaDBData struct {
	Id                     string `db:"Id"`
	OriginMessageId        string `db:"originMessageId"`
	ChannelId              string `db:"channelId"`
	AnkaInvokeMessageCount int    `db:"ankaInvokeMessageCount"`
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

func (d *DB) AnkaNumDecrementByChannel(channelId string) (err error) {
	_, err = d.dataBase.Exec("UPDATE `anka` SET `ankaInvokeMessageCount` = `ankaInvokeMessageCount` - 1 WHERE `channelId` = ?", channelId)
	if err != nil {
		log.Println("ankaDecrementError:", err.Error())
		return err
	}
	return nil
}

func (d *DB) AnkaDeleteFromDB() (err error) {
	_, err = d.dataBase.Exec("DELETE FROM `anka` WHERE `ankaInvokeMessageCount` <= 0")
	if err != nil {
		log.Println("ankaDeleteError:", err.Error())
		return err
	}
	return nil
}

func (d *DB) AnkaDataSetupFromDB() (ankasDBData []AnkaDBData, err error) {
	result := []AnkaDBData{}
	err = d.dataBase.Select(&result, "SELECT * FROM `anka` ORDER BY `ankaInvokeMessageCount` DESC")
	if err != nil {
		log.Println("ankaSetupFromDBError:", err.Error())
	}
	return result, err
}
