package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Seller struct {
	Metadata
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    []byte `json:"-"`
	ServicesIDs uuid.UUID
}

type UserAutenticate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *Seller) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *Seller) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
