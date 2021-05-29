package models

type Message struct {
	MessageID int   `json:"message_id"`
	From      *User `json:"from"`
	// Date of the message was sent in Unix time
	Date                 int              `json:"date"`
	Chat                 *Chat            `json:"chat"`
	ForwardFrom          *User            `json:"forward_from"`
	ForwardFromChat      *Chat            `json:"forward_from_chat"`
	ForwardFromMessageID int              `json:"forward_from_message_id"`
	ForwardDate          int              `json:"forward_date"`
	ReplyToMessage       *Message         `json:"reply_to_message"`
	MediaGroupID         string           `json:"media_group_id"`
	AuthorSignature      string           `json:"author_signature"`
	Text                 string           `json:"text"`
	Entities             *[]MessageEntity `json:"entities"`
	Photo                *[]PhotoSize     `json:"photo"`
	Contact              *Contact         `json:"contact"`
	Location             *Location        `json:"location"`
	Invoice              *Invoice         `json:"invoice"`
	// SuccessfulPayment message is a service message about a successful payment,
	// information about the payment;
	//
	// optional
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment"`
}
