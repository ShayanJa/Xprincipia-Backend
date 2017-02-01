package gorm

import "github.com/jinzhu/gorm"

//User : a normal user on XPrincipia
type User struct {
	gorm.Model
	FirstName       string
	LastName        string
	Email           string
	Address         string
	Username        string
	PhoneNumber     int
	Friends         []User
	ProblemsPosted  []Problem
	SolutionsPosted []Solution
	Comments        []Comment
}

//API Functions
