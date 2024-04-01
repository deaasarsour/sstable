package server

import (
	"encoding/json"
	"sstable/communication/common"
)

type writeMessageContent struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func (dbConnection *databaseConnection) executeMessage(dbMessage *common.DatabaseMessage) {

	databaseIO := dbConnection.dbms.DatabaseIO

	if dbMessage.Key == common.ReadKey {
		key := string(dbMessage.Content)
		result, err := databaseIO.Read(key)

		if err != nil {
			panic(err)
		}

		dbConnection.WriteBack("rr", result)
	} else if dbMessage.Key == common.WriteKey {
		var writeMessage writeMessageContent
		if err := json.Unmarshal(dbMessage.Content, &writeMessage); err != nil {
			panic(err)
		}
		databaseIO.Write(writeMessage.Key, writeMessage.Value)
		dbConnection.WriteBack("wr", "write successful!")
	} else {
		dbConnection.WriteBack("kn", "key not found!")
	}
}

func (dbConnection *databaseConnection) executeRawMessage(message []byte) {
	parsedMessage := common.ParseMessage(message)
	dbConnection.executeMessage(parsedMessage)
}
