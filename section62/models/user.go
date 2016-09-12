package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `db:"Username"`
	Secret []byte `db:"Secret"`
}

func NewUser(un, pw string) *User {
	secret, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return &User{
		Username: un, 
		Secret: secret,
	}
}

func (u *User) Authenticate(pw string) bool {
	return bcrypt.CompareHashAndPassword(u.Secret, []byte(pw)) == nil
}