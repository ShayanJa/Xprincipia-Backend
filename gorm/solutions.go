package gorm

import (
	"strconv"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// Solution : Generic Problem Solution
type Solution struct {
	gorm.Model
	ProblemID              uint
	OriginalPoster         User `gorm:"ForeignKey:OriginalPosterUsername;AssociationForeignKey:Username" json:"originalPoster" form:"originalPoster"`
	OriginalPosterUsername string
	Title                  string
	Summary                string
	Description            string
	Evidence               string
	Experiments            string
	References             string
	Rank                   int
	Suggestions            []Suggestion
	Questions              []Question
}

//SolutionForm : Solution Form
type SolutionForm struct {
	ProblemID   string `json:"problemID" form:"problemID"`
	Title       string `json:"title" form:"title"`
	Summary     string `json:"summary" form:"summary"`
	Description string `json:"description" form:"description"`
	Evidence    string `json:"evidence" form:"evidence"`
	Experiments string `json:"experiments" form:"experiments"`
	References  string `json:"references" form:"references"`
}

// GetSolutionByID : returns a solution by its id
func (s *Solution) GetSolutionByID(id uint) {
	err := db.Where("id = ?", id).First(&s)
	if err == nil {
		glog.Info("There was an error")
	}
}

// GetSolutionsByProblemID : returns a solution by its id
func GetSolutionsByProblemID(id int) []Solution {
	s := []Solution{}
	err := db.Where("problem_id = ?", id).Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}
	glog.Info(s)
	return s
}

//GetAllSolutions : Get all Solutions in db
func GetAllSolutions() []Solution {
	s := []Solution{}
	err := db.Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}
	glog.Info(s)
	return s
}

// CreateSolution : Creates solution from solutionForm
func CreateSolution(form SolutionForm) {
	s := Solution{}

	//Create Solution object based on solutionForm info
	intID, _ := strconv.Atoi(form.ProblemID)
	s.ProblemID = uint(intID)
	s.Title = form.Title
	s.Summary = form.Summary
	s.Description = form.Description
	s.Evidence = form.Evidence
	s.Experiments = form.Experiments
	s.References = form.References
	s.Rank = 1

	db.Create(&s)
}
