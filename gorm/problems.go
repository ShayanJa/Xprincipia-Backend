package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// Problem : User generated problem
type Problem struct {
	gorm.Model
	ParentID               int
	OriginalPoster         User `gorm:"ForeignKey:OriginalPosterUsername;AssociationForeignKey:Username" json:"originalPoster" form:"originalPoster"`
	OriginalPosterUsername string
	Title                  string
	Field                  string
	Summary                string `gorm:"size:1000"`
	Description            string `gorm:"size:10000"`
	Requirements           string
	References             string
	SubProblems            []Problem
	Suggestions            []Suggestion
	Questions              []Question
}

//ProblemForm : form to create problem
type ProblemForm struct {
	Title        string
	Field        string
	Summary      string
	Description  string
	Requirements string
	References   string
}

// GetProblemByID : returns a solution by its id
func (p *Problem) GetProblemByID(id uint) {
	err := db.Where("id = ?", id).First(&p)
	if err == nil {
		glog.Info("There was an error")
	}
}

// GetProblemBySolutionID : returns a solution by its id
func (p *Problem) GetProblemBySolutionID(id uint) {
	s := Solution{}
	err := db.Where("id = ?", id).First(&s)
	if err == nil {
		glog.Info("There was an error")
	}
	p.GetProblemByID(s.ID)
}

//CreateProblem : Creates a problem from a problemForm
func CreateProblem(form ProblemForm) {
	p := Problem{}
	p.Title = form.Title
	p.Summary = form.Summary
	p.Description = form.Description
	db.Create(&p)
}

// UpdateProblem : Updates a problem with problemForm as input
func (p *Problem) UpdateProblem(form ProblemForm) {
	err := db.First(&p)
	if err == nil {
		glog.Info("There was an error")
	}

	p.Description = form.Description
	p.Summary = form.Summary
	p.Title = form.Title
	db.Save(&p)
}

//GetAllProblems : Returns all problem objects
func GetAllProblems() []Problem {
	p := []Problem{}
	err := db.Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	return p
}

//QueryProblems : Return problems that are related to the query String
func QueryProblems(q string) []Problem {
	p := []Problem{}
	err := db.Where("title LIKE ?", "%"+q+"%").Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	return p
}
