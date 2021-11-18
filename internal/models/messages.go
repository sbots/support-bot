package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	MessageTypeFromUser     MessageType = "FROM_USER"
	MessageTypeFromCustomer MessageType = "FROM_CUSTOMER"
	ContentTypeText         ContentType = "Text"
)

type MessageType string
type ContentType string

func (m MessageType) String() string {
	return string(m)
}

func (c ContentType) String() string {
	return string(c)
}

type Message struct {
	ID          string      `json:"id" db:"id"`
	Chat        string      `json:"chat_id" db:"chat_id"`
	Text        string      `json:"text" db:"text"`
	Type        MessageType `json:"type" db:"type"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at" db:"deleted_at"`
	ContentType ContentType `json:"content_type"`
}

type Chat struct {
	ID        string     `json:"id" db:"id"`
	Company   string     `json:"company_id" db:"company_id"`
	User      string     `json:"user_id" db:"user_id"`
	Customer  string     `json:"customer_id" db:"customer_id"`
	IsActive  bool       `json:"content_type" db:"is_active"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

func NewMessage(chat, text string, msgType MessageType) *Message {
	return &Message{
		ID:        uuid.NewV4().String(),
		Chat:      chat,
		Text:      text,
		Type:      msgType,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
