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
	Title          string
	Summary        string
	Description    string
	Evidence       string
	Rank           int
	Comments       []Comment
}

//SolutionForm : Solution Form
type SolutionForm struct {
	Title       string `json:"title" form:"title"`
	Summary     string `json:"summary" form:"summary"`
	Description string `json:"description" form:"description"`
	Evidence    string `json:"evidence" form:"evidence"`
}

// GetSolutionByID : returns a solution by its id
func (s *Solution) GetSolutionByID(id uint) {
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

//GetAllSolutions : Get all Solutions in db
func GetAllSolutions() []Solution {
	s := []Solution{}
	err := db.Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}
	return s
}

// CreateSolution : Creates solution from solutionForm
func CreateSolution(form SolutionForm) {
	s := Solution{}

	//Create Solution object based on solutionForm info
	s.Title = form.Title
	s.Summary = form.Summary
	s.Description = form.Description
	s.Evidence = form.Evidence
	s.Rank = 1

	db.Create(&s)
}
