package models

const (
	MessageTypeText MessageType = "Text"
)

type MessageType string

func (m MessageType) String() string {
	return string(m)
}

type Message struct {
	ContentType MessageType `json:"type"`
	Text        string      `json:"text"`
}

func NewMessage(contentType MessageType, text string) *Message {
	return &Message{ContentType: contentType, Text: text}
}
