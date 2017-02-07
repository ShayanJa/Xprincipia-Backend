package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//User : ~
type User struct {
	gorm.Model
	FirstName      string `json:"firstName" form:"firstName"`
	LastName       string `json:"LastName" form:"lastName"`
	Email          string `json:"email" form:"email"`
	Address        string `json:"address" form:"address"`
	Username       string `json:"username" form:"username"`
	PhoneNumber    string `json:"phoneNumber" form:"phoneNumber"`
	HashedPassword []byte `json:"hashedPassword" form:"hashedPassword"`
	// Friends         []User
	ProblemsPosted []Problem
	// SolutionsPosted []Solution
	// Comments        []Comment
	IsDisabled bool
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

//GetUserByUsername : get user by name
func (u *User) GetUserByUsername(name string) bool {
	err := db.Where("username = ?", name).First(&u)
	if err == nil {
		glog.Info("There was an error")
	}
	if u.ID == 0 {
		return false
	}
	return true
}

//VerifyUser : Checks db credentials
func (u *User) VerifyUser(username string, password string) bool {
	err := db.Where("username = ? AND hashed_password", username, password).First(&u)
	if err == nil {
		glog.Info("There was an error")
	}
	if u.ID == 0 {
		return false
	}
	return true
}

//LoginAttempt : Logs everytime someone logs on
func (l LoginForm) LoginAttempt() {
	db.Create(l)
}

// PostProblem : User Auth Required> Post Problem
func (u *User) PostProblem(text string, description string) {
	p := Problem{
		OriginalPoster: *u,
		Title:          text,
		Description:    description,
	}
	db.Create(&p)
	u.ProblemsPosted = append(u.ProblemsPosted, p)
}

//PostSolution : User Auth Required> Post Solution
func (u *User) PostSolution(p Problem, text string, description string) {
	s := Solution{
		ProblemID:      p.ID,
		OriginalPoster: *u,
		Text:           text,
		Rating:         0,
	}
	db.Create(s)
}
