package gorm

import (
	"strconv"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

// Problem : User generated problem
type Problem struct {
	gorm.Model
	ParentID               int
	OriginalPoster         User `gorm:"ForeignKey:OriginalPosterUsername;AssociationForeignKey:Username" json:"originalPoster" form:"originalPoster"`
	OriginalPosterUsername string
	Title                  string `gorm:"size:151"`
	Field                  string `gorm:"size:151"`
	Summary                string `gorm:"size:1500"`
	Description            string `gorm:"size:10000"`
	Requirements           string `gorm:"size:1500"`
	References             string `gorm:"size:1500"`
	Rank                   int
	SubProblems            []Problem
	Suggestions            []Suggestion
	Questions              []Question
}

//ProblemForm : form to create problem
type ProblemForm struct {
	Username     string
	ParentID     string
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
	p.OriginalPosterUsername = form.Username
	intID, _ := strconv.Atoi(form.ParentID)
	p.ParentID = intID
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
	err := db.Where("parent_id < ?", 1).Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	return p
}

//GetSubProblemsByID : ~
func GetSubProblemsByID(parentID int) []Problem {
	p := []Problem{}
	db.Where("parent_id = ?", parentID).Find(&p)
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

//VoteProblem : ~
func (p *Problem) VoteProblem(id int) {
	err := db.Where("id = ?", id).Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	p.Rank++
	db.Model(&p).Update("rank", p.Rank)

}
