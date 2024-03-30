package common

const WriteKey string = "wd"
const ReadKey string = "rd"

type DatabaseMessage struct {
	Key     string
	Content []byte
}

func ParseMessage(message []byte) *DatabaseMessage {
	key := string(message[:2])
	content := message[2:]

	return &DatabaseMessage{
		Key:     key,
		Content: content,
	}
}
