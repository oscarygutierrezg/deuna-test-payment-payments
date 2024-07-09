package models

import (
	"github.com/google/uuid"
	"payment-payments-api/pkg/auth"
	"time"
)

func NewUser(firstName, lastName, email string) (*User, error) {
	u := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Enabled:   true,
		CreatedAt: time.Now(),
	}
	u.DefaultPassword()
	return u, nil
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Password     string    `json:"password,omitempty"`
	PasswordSalt string    `json:"passwordSalt,omitempty"`
	Enabled      bool      `json:"enabled"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (u *User) DefaultPassword() {
	defaultPassword := "admin123"
	u.SetPassword(defaultPassword)
}

func (u *User) SetPassword(password string) {
	u.PasswordSalt, u.Password = auth.EncryptPassword(password)
}

func (u *User) CleanSensitiveInfo() {
	u.Password = ""
	u.PasswordSalt = ""
}
