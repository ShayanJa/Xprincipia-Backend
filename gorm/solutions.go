package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// Solution : Generic Problem Solution
type Solution struct {
	gorm.Model
	ProblemID      uint
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

// GetSolutionByProblemID : returns a solution by its id
func (s *Solution) GetSolutionByProblemID(id int) {
	err := db.Where("problem_id = ?", id).First(&s)
	if err == nil {
		glog.Info("There was an error")
	}
}

// Upvote : user can upvote on a solution and change it's ranking
func (s *Solution) Upvote(u User) {
	
}
