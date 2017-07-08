package gorm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// LoginAttempt :
type LoginAttempt struct {
	gorm.Model
	Username string
	Token    string
}

// LoginAttemptForm :
type LoginAttemptForm struct {
	Username string
	Token    string
}

// CreateLoginAttempt : Create a login attempt, table of usernames and tokens
func CreateLoginAttempt(username string, token string) {
	l := LoginAttempt{}
	l.Username = username
	l.Token = token

	//else create db Entry
	db.Create(&l)

}

// CheckToken : Check for table entry
func CheckToken(username string, token string) error {
	l := LoginAttempt{}
	db.Where(" username = ? AND token = ? ", username, token).Find(&l)
	if l.ID == 0 {
		return errors.New("Invalid Token")
	}
	return nil
}
