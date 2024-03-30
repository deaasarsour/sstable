package server

import "fmt"

func (dbConnection *databaseConnection) WriteBack(key string, content any) {

	contentText := fmt.Sprintf("%v", content)

	rawMessage := append([]byte{}, []byte(key)...)
	rawMessage = append(rawMessage, []byte(contentText)...)
	rawMessage = append(rawMessage, byte('\n'))
	dbConnection.writeBackRaw(rawMessage)
}

func (dbConnection *databaseConnection) writeBackRaw(message []byte) {
	for len(message) > 0 {
		n, err := dbConnection.conn.Write(message)
		if err != nil {
			panic(err)
		}
		message = message[n:]
	}
}
