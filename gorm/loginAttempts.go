package gorm

import "github.com/jinzhu/gorm"

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
	if CheckLoginAttempt(username, token) {
		return
	}
	db.Create(&l)
	return
}

// CheckLoginAttempt : Check for table entry
func CheckLoginAttempt(username string, token string) bool {
	l := LoginAttempt{}
	db.Where(" username = ? AND token = ? ", username, token).Find(&l)
	if l.ID == 0 {
		return false
	}
	return true
}
