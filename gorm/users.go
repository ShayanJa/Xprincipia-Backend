package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"work/xprincipia/backend/util"
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
	IsDisabled          bool
}

//LoginForm : ~
type LoginForm struct {
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
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

//API Functions

// CreateUser : check if user is already created,
// if not use RegistrationForm to populate a new one
func CreateUser(form RegistrationForm) {

	//check DB if Username is already taken
	u := User{}
	err := db.Where("username = ?", form.Username).First(&u).Value
	if err == nil {
		glog.Info("error has occured")
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

		db.Create(&u)
	} else {
		glog.Error("Username is already taken")
	}

}

//GetUserByUsername : get user by name
func (u *User) GetUserByUsername(name string) bool {
	glog.Info("Getting Username : " + name + " ...")
	err := db.Where("username = ?", name).First(&u)
	if err == nil {
		glog.Info("There was an error...")
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
		glog.Info("There was an error")
	}
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
	u.ProblemsPostedIDs = append(u.ProblemsPostedIDs, p)
}

// getFollowedProblems : returns problemIDs of all problems followed by the user
//TODO:
//THis doesn't work right
func (u User) getFollowedProblems() []int {
	var followedProblems []int
	err := db.Where("followed_problems = ?").Find(&followedProblems)
	if err == nil {
		glog.Error("Unable to retrieve users followed problems")
	}
	return followedProblems
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
	//this needs to be fixed. to be dynamic
	for index, vote := range votes {
		s := Solution{}
		s.GetSolutionByID(uint(vote.TypeID))
		glog.Info(s)
		solutions[index] = s

	}
	return solutions[:len(votes)]
}

// VoteOnSolution : User votes on a solution to increase it's rank
func (u *User) VoteOnSolution(solutionID uint) {
	solution := Solution{}
	solution.GetSolutionByID(solutionID)

	//Check if user has already voted on this problem.
	//if found look for what problem this is and lower rank on a different solution
	for _, votedProblem := range u.VotedProblemIDs {
		for _, votedSolution := range u.VotedSolutionsIDs {
			if votedProblem.ID == votedSolution.ProblemID {
				s := Solution{}
				s.GetSolutionByID(votedSolution.ID)
				s.Rank--
			}
		}
	}
	solution.Rank++
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
