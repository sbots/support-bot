package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	MessageTypeText MessageType = "Text"
)

type MessageType string

func (m MessageType) String() string {
	return string(m)
}

type Message struct {
	ID          string      `json:"id" db:"id"`
	Company     string      `json:"company_id" db:"company_id"`
	User        string      `json:"user_id" db:"user_id"`
	Customer    string      `json:"customer_id" db:"customer_id"`
	Text        string      `json:"text" db:"text"`
	ContentType MessageType `json:"content_type"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at" db:"deleted_at"`
}

func NewMessage(company, user, customer, text string, contentType MessageType) *Message {
	return &Message{
		ID:          uuid.NewV4().String(),
		Company:     company,
		User:        user,
		Customer:    customer,
		Text:        text,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		ContentType: contentType,
	}
}
