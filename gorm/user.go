package gorm

import "github.com/jinzhu/gorm"

//User : a normal user on XPrincipia
type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	//Friends   []gorm.Friends
	//ProblemsPosted []gorm.problem
	//SolutionsPosted []gorm.solution
	//Comments []gorm.comments
}
