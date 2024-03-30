package server

import (
	"encoding/json"
	"fmt"
	"net"
	"sstable/dbms/core"
)

type writeMessageContent struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

func executeMessage(dbMessage *databaseMessage, dbms *core.DatabaseManagementSystem, conn net.Conn) {
	if dbMessage.key == "rd" {
		key := string(dbMessage.content)
		result, err := dbms.DatabaseManagement.Read(key)

		if err != nil {
			panic(err)
		}

		fmt.Printf("read result: %v\n", result)
	} else if dbMessage.key == "wd" {
		var writeMessage writeMessageContent
		fmt.Println(string(dbMessage.content))
		if err := json.Unmarshal(dbMessage.content, &writeMessage); err != nil {
			panic(err)
		}
		fmt.Printf("start writing\n")

		dbms.DatabaseManagement.Write(writeMessage.Key, writeMessage.Value)

		fmt.Printf("write done\n")
	}
}
