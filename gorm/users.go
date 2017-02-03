package gorm

import "github.com/jinzhu/gorm"

//User : ~
type User struct {
	gorm.Model
	FirstName   string `json:"firstName" form:"firstName"`
	LastName    string `json:"LastName" form:"lastName"`
	Email       string `json:"email"`
	Address     string
	Username    string
	PhoneNumber string
	// Friends         []User
	// ProblemsPosted  []Problem
	// SolutionsPosted []Solution
	// Comments        []Comment
}

//LoginForm : ~
type LoginForm struct {
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
}

//PasswordResetForm : ~
type PasswordResetForm struct {
	Email string `json:"email" form:"email"`
}

//API Functions

//LoginAttempt : Logs everytime someone logs on
func (l LoginForm) LoginAttempt() {
	db.Create(l)
}

// PostProblem : ~
func (u *User) PostProblem() {

}
