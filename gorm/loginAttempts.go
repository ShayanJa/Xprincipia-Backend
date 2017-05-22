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
func CreateLoginAttempt(form LoginAttemptForm) {
	l := LoginAttempt{}
	l.Username = form.Username
	l.Token = form.Token
}

// CheckLoginAttempt : Check for table entry
func CheckLoginAttempt(form LoginAttemptForm) bool {
	l := LoginAttempt{}
	db.Where(" username = ? AND token = ? ", form.Username, form.Token).Find(&l)
	if l.ID == 0 {
		return false
	}
	return true
}
