package gorm

import "github.com/jinzhu/gorm"

//User : a normal user on XPrincipia
type User struct {
	gorm.Model
	FirstName       string `json:"firstName" form:"firstName"`
	LastName        string `json:"LastName" form:"lastName"`
	Email           string `json:"email"`
	Address         string
	Username        string
	PhoneNumber     int
	Friends         []User
	ProblemsPosted  []Problem
	SolutionsPosted []Solution
	Comments        []Comment
}

//LoginForm :LoginForm Struct
type LoginForm struct {
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
}

//API Functions

//LoginAttempt : Logs everytime someone logs on
func (l *LoginForm) LoginAttempt() {
	db.Create(l)
}
