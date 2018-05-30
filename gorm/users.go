package gorm

import (
	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"../util"
)

//User : ~
type User struct {
	gorm.Model
	FirstName           string `json:"firstName" form:"firstName"`
	LastName            string `json:"LastName" form:"lastName"`
	FullName            string
	Email               string `json:"email" form:"email"`
	Address             string `json:"address" form:"address"`
	Username            string `json:"username" form:"username"`
	PhoneNumber         string `json:"phoneNumber" form:"phoneNumber"`
	HashedPassword      []byte `json:"hashedPassword" form:"hashedPassword"`
	FriendsIDs          []User
	ProblemsPostedIDs   []Problem
	SolutionsIDs        []Solution
	FollowedProblemsIDs []Problem
	VotedSolutionsIDs   []Solution
	VotedProblemIDs     []Problem
	Points              int
	IsDisabled          bool
}

//LoginForm : ~
type LoginForm struct {
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
}

//LoginAttempt : Logs everytime someone logs on
func (l LoginForm) LoginAttempt() {
	db.Create(l)
}

// RegistrationForm : A registration struct
type RegistrationForm struct {
	FullName string `json:"fullName" form:"fullName"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

//PasswordResetForm : ~
type PasswordResetForm struct {
	Email string `json:"email" form:"email"`
}

// //ErrorTypes
// var ErrFormSubmission = errors.New("Form was incomplete")

//API Functions

// CreateUser : Validate form fields and check if user is already created,
// if not use RegistrationForm to populate a new one
func CreateUser(form RegistrationForm) error {

	//Validate register form
	switch {
	case form.Username == "":
		return errors.New("Please enter a username. ")
	case form.Email == "":
		return errors.New("Please enter an email address. ")
	case form.FullName == "":
		return errors.New("Please enter your full name. ")
	case len(form.Password) < 8:
		return errors.New("Please enter a password with at least 8 characters. ")
	}

	//check DB if Username is already taken
	u := User{}
	err := db.Where("username = ?", form.Username).First(&u).Value
	if err == nil {
		glog.Error("error has occured")
	}
	//If username does not exist
	if u.Username == "" {
		glog.Info("Username not taken...")
		//generate hashpassword
		passwordBytes := []byte(form.Password)
		hashedPassword, hashError := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
		if hashError != nil {
			glog.Error("Bcrypt failed to hash password")
		}

		//Reinitalize user as new user and set it with form details
		u = User{}
		u.HashedPassword = hashedPassword
		u.Email = form.Email
		u.Username = form.Username
		u.FullName = form.FullName

		db.Create(&u)
	} else {
		return errors.New("Username is already taken")
	}
	return nil
}

//GetUserByUsername : get user by name
func (u *User) GetUserByUsername(name string) bool {
	glog.Info("Getting Username : " + name + " ...")
	err := db.Where("username = ?", name).First(&u)
	if err == nil {
		glog.Error("There was an error...")
	}
	if u.ID == 0 {
		glog.Info("NO USER BY THAT NAME...")
		return false
	}
	glog.Info("USERNAME FOUND RETURNING " + u.Username)
	return true
}

//GetUserByID : get user by ID
func (u *User) GetUserByID(id int) {
	err := db.Where("ID = ?", id).First(&u)
	if err == nil {
		glog.Error("There was an error")
	}
}

//VerifyUser : Checks db credentials
func (u *User) VerifyUser(username string, password string) bool {
	err := db.Where("username = ? AND hashed_password", username, password).First(&u)
	if err == nil {
		glog.Error("There was an error")
	}
	if u.ID == 0 {
		return false
	}
	return true
}

//GetAllCreatedSolutions :
func (u *User) GetAllCreatedSolutions() []Solution {
	solutions := []Solution{}
	db.Where("original_poster_username = ?", u.Username).Find(&solutions)
	return solutions
}

//GetAllCreatedProblems :
func (u *User) GetAllCreatedProblems() []Problem {
	problems := []Problem{}
	db.Where("original_poster_username = ?", u.Username).Find(&problems)
	return problems
}

//GetAllFollowedSolutions :
func (u *User) GetAllFollowedSolutions() []Solution {
	votes := []Vote{}
	db.Where("type = ? AND username = ?", util.SOLUTION, u.Username).Find(&votes)

	var solutions [100]Solution
	for index, vote := range votes {
		s := Solution{}
		s.GetSolutionByID(uint(vote.TypeID))
		solutions[index] = s

	}
	return solutions[:len(votes)]
}

//GetAllFollowedProblems :
func (u *User) GetAllFollowedProblems() []Problem {
	votes := []Vote{}
	db.Where("type = ? AND username = ?", util.PROBLEM, u.Username).Find(&votes)

	problems := []Problem{}
	for _, vote := range votes {
		p := Problem{}
		p.GetProblemByID(uint(vote.TypeID))
		problems = append(problems, p)

	}
	return problems
}

//DeleteUserByID : //DELETE
func DeleteUserByID(id int) {
	u := User{}
	u.GetUserByID(id)
	db.Delete(&u)
}

//DisableUser : disables user
func DisableUser(id int) {
	u := User{}
	u.GetUserByID(id)
	u.IsDisabled = true
	db.Model(&u).Update("is_disabled", u.IsDisabled)
}

/*

DB bool functions


*/

// IsUserinDBbyEmail : checks if a user is in the db
func IsUserinDBbyEmail(email string) bool {
	u := User{}
	db.Where("email = ?", email).First(&u)
	if u.ID == 0 {
		return false
	}
	return true
}

// IsUserinDBbyUsername : checks if a user is in the db
func IsUserinDBbyUsername(username string) bool {
	u := User{}
	db.Where("username = ?", username).First(&u)
	if u.ID == 0 {
		return false
	}
	return true
}

/* Points API */

//AddPoints : add points to user
func (u *User) AddPoints(points int) {
	err := db.First(&u)
	if err == nil {
		glog.Error("There was an error")
	}

	u.Points += points
	db.Save(&u)

}

//GetPoints : get points of the user
func (u *User) GetPoints() int {
	err := db.First(&u)
	if err == nil {
		glog.Error("There was an error")
	}
	return u.Points

}
