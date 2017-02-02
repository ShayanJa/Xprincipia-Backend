package gorm

import (
	"github.com/golang/glog"
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

// GetSolutionByID : returns a solution by its id
func (s *Solution) GetSolutionByID(id int) {
	err := db.Where("id = ?", id).First(&s)
	if err == nil {
		glog.Info("There was an error")
	}
}
