package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Customer struct {
	ID             string     `json:"id" db:"id"`
	Name           string     `json:"name" db:"name"`
	Surname        string     `json:"surname" db:"surname"`
	Platform       string     `json:"platform" db:"platform"`
	PlatformUserID string     `json:"platform_user_id" db:"platform_user_id"`
	Company        string     `json:"company_id" db:"company_id"`
	Phone          string     `json:"phone" db:"phone"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
}

func NewCustomer(name, surname, platform, platformUserID, company, phone string) *Customer {
	return &Customer{
		ID:             uuid.NewV4().String(),
		Name:           name,
		Surname:        surname,
		Platform:       platform,
		PlatformUserID: platformUserID,
		Company:        company,
		Phone:          phone,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}
}
