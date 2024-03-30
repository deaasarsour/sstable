package server

import (
	"fmt"
	"net"
	"sstable/dbms/core"
)

type databaseConnection struct {
	conn          net.Conn
	dbms          *core.DatabaseManagementSystem
	readingBuffer []byte
	messageBuffer []byte
	readingOffset int
}

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func (dbConnection *databaseConnection) startCommunication() {
	defer dbConnection.conn.Close()
	defer recoverPanic()

	for {
		if message := dbConnection.tryReadMessage(); message != nil {
			isCloseMessage := len(message) < 2
			if isCloseMessage {
				break
			} else {
				dbConnection.executeRawMessage(message)
			}

		} else {
			dbConnection.expandMessageBuffer()
		}
	}
}

func NewDatabaseConnection(conn net.Conn, dbms *core.DatabaseManagementSystem) *databaseConnection {
	return &databaseConnection{
		conn:          conn,
		dbms:          dbms,
		readingBuffer: make([]byte, 1024),
		messageBuffer: make([]byte, 0),
		readingOffset: 0,
	}
}
