package gorm

import (
	"github.com/jinzhu/gorm"
)

// Solution : Generic Problem Solution
type Solution struct {
	gorm.Model
	ProblemID      int
	OriginalPoster User
	Text           string
	Rating         int
	Comments       []Comment
}

func getSolutionByID(id int) {
	db.Where("ID = ?", id).Find(&User{})
}
