package models

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Surname   string     `json:"surname" db:"surname"`
	Password  string     `json:"password" db:"password"`
	Company   string     `json:"company" db:"company_id"`
	Email     string     `json:"email" db:"email"`
	Phone     string     `json:"phone" db:"phone"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

func NewUser(name, surname, password, companyID, email, phone string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if surname == "" {
		return nil, fmt.Errorf("surname is required")
	}
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	if phone == "" {
		return nil, fmt.Errorf("phone is required")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, fmt.Errorf("hasging the password")
	}

	return &User{
		ID:        uuid.NewV4().String(),
		Name:      name,
		Surname:   surname,
		Password:  string(hashedPassword),
		Company:   companyID,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}

func (u *User) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type Tenant struct {
	ID        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

func NewTenant(name string) *Tenant {
	return &Tenant{
		ID:        uuid.NewV4().String(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
