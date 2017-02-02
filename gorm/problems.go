package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// Problem : User generated problem
type Problem struct {
	gorm.Model
	OriginalPoster User
	Title          string
	Description    string
	SubProblems    []Problem
	Comments       []Comment
}

// GetProblemByID : returns a solution by its id
func (p *Problem) GetProblemByID(id int) {
	err := db.Where("id = ?", id).First(&p)
	if err == nil {
		glog.Info("There was an error")
	}
}
