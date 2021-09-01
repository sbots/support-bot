package viber

const (
	messageTypeText = "text"
)

type Message struct {
	Receiver      string  `json:"receiver"`
	MinApiVersion int     `json:"min_api_version"`
	Sender        *Sender `json:"sender"`
	TrackingData  string  `json:"tracking_data"`
	ContentType   string  `json:"type"`
	Text          string  `json:"text"`
}

func newMessage(minApiVersion int, sender *Sender, receiver, trackingData, contentType, text string) *Message {
	return &Message{Receiver: receiver, MinApiVersion: minApiVersion, Sender: sender, TrackingData: trackingData, ContentType: contentType, Text: text}
}

type Sender struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func newSender(name string, avatar string) *Sender {
	return &Sender{Name: name, Avatar: avatar}
}
