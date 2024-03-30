package server

func (dbConnection *databaseConnection) expandMessageBuffer() {
	n, err := dbConnection.conn.Read(dbConnection.readingBuffer)
	if err != nil {
		panic(err)
	}

	dbConnection.messageBuffer = append(dbConnection.messageBuffer, dbConnection.readingBuffer[:n]...)
}

func getEndOfMessage(messageBuffer []byte, searchIndex int) int {
	for i := searchIndex; i < len(messageBuffer); i++ {
		if messageBuffer[i] == byte('\n') {
			return i
		}
	}

	return len(messageBuffer)
}

func (dbConnection *databaseConnection) tryReadMessage() []byte {
	dbConnection.readingOffset = getEndOfMessage(dbConnection.messageBuffer, dbConnection.readingOffset)
	isMessageAvailable := dbConnection.readingOffset != len(dbConnection.messageBuffer)
	if isMessageAvailable {
		message := dbConnection.messageBuffer[:dbConnection.readingOffset]
		dbConnection.messageBuffer = dbConnection.messageBuffer[dbConnection.readingOffset+1:]
		dbConnection.readingOffset = 0
		return message
	}
	return nil
}
