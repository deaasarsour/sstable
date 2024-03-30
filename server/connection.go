package server

import (
	"fmt"
	"net"
	"sstable/dbms/core"
)

type databaseMessage struct {
	key     string
	content []byte
}

type databaseConnection struct {
	dbms          core.DatabaseManagementSystem
	readingBuffer []byte
}

func expandMessageBuffer(conn net.Conn, readingBuffer []byte, messageBuffer *[]byte) {
	n, err := conn.Read(readingBuffer)
	if err != nil {
		panic(err)
	}

	*messageBuffer = append(*messageBuffer, readingBuffer[:n]...)
}

func getMessageIndexFromBufferIfExist(messageBuffer []byte, searchIndex int) int {
	for i := searchIndex; i < len(messageBuffer); i++ {
		if messageBuffer[i] == byte('?') {
			return i
		}
	}

	return len(messageBuffer)
}

func parseMessage(message []byte) *databaseMessage {
	key := string(message[:2])
	content := message[2:]

	return &databaseMessage{
		key:     key,
		content: content,
	}
}

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func readAndExecMessage(message []byte, dbms *core.DatabaseManagementSystem, conn net.Conn) {
	parsedMessage := parseMessage(message)
	fmt.Printf("Key:%v, Content len:%v\n", parsedMessage.key, len(parsedMessage.content))

	executeMessage(parsedMessage, dbms, conn)
}

func handleConnection(conn net.Conn, dbms *core.DatabaseManagementSystem) {
	defer conn.Close()
	defer recoverPanic()

	buffer := make([]byte, 1024)
	currentMessage := make([]byte, 0)
	offset := 0

	for {

		offset = getMessageIndexFromBufferIfExist(currentMessage, offset)
		isMessageAvailable := offset != len(currentMessage)

		if isMessageAvailable {
			message := currentMessage[:offset]
			isEndOfMessage := len(message) > 1
			if isEndOfMessage {
				currentMessage = currentMessage[offset+1:]
				readAndExecMessage(message, dbms, conn)
			} else {
				break
			}

		} else {
			expandMessageBuffer(conn, buffer, &currentMessage)
		}

	}
}
