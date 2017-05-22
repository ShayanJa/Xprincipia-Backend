package gorm

import (
	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

type LoginAttempt struct {
	gorm.Model
	Username string
	Token    string
}

type LoginAttemptForm struct {
	Username string
	Token    string
}

// CreateLoginAttempt : Create a login attempt, table of usernames and tokens
func CreateLoginAttempt(username string, token string) {
	l := LoginAttempt{}
	l.Username = username
	l.Token = token
	err := CheckToken(username, token)
	if err != nil {
		glog.Error(err)
		return
	}

	//else create db Entry
	db.Create(&l)

}

// CheckLoginAttempt : Check for table entry
func CheckToken(username string, token string) error {
	l := LoginAttempt{}
	db.Where(" username = ? AND token = ? ", username, token).Find(&l)
	if l.ID == 0 {
		return errors.New("Invalid Token")
	}
	return nil
}
