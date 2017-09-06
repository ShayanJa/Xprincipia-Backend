package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"work/xprincipia/backend/util"
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
	Description            string `gorm:"size:100000"`
	Requirements           string `gorm:"size:1500"`
	References             string `gorm:"size:1500"`
	Rank                   uint
	PercentRank            float32
	Private                uint
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
	Private      string
}

// ProblemDeleteForm :
type ProblemDeleteForm struct {
	Username string
	ID       int
}

// GetProblemByID : returns a solution by its id
func (p *Problem) GetProblemByID(id uint) {
	err := db.Where("id = ?", id).First(&p)
	if err == nil {
		glog.Error("There was an error")
	}
}

//GetProblemByTitle :
func (p *Problem) GetProblemByTitle(title string) {
	err := db.Where("title = ?", title).First(&p)
	if err == nil {
		glog.Error("There was an error")
	}
}

//CreateProblem : Creates a problem from a problemForm
func CreateProblem(form ProblemForm) error {

	//Handle form Field Errors
	switch {
	case form.Title == "":
		return errors.New("Title is empty: Please fill in field")
	case form.Summary == "":
		return errors.New("Additional Information is empty: Please fill in field")
	}

	//Create Problem with Form Items
	p := Problem{}
	p.OriginalPosterUsername = form.Username
	intID, _ := strconv.Atoi(form.ParentID)
	p.ParentID = intID
	p.Title = form.Title
	p.Field = form.Field
	p.Summary = form.Summary
	p.Description = form.Description
	p.References = form.References
	p.Requirements = form.Requirements
	p.Rank = 0
	intPrivate, _ := strconv.Atoi(form.Private)
	p.Private = uint(intPrivate)
	db.Create(&p)
	return nil
}

// UpdateProblem : Updates a problem with problemForm as input
func (p *Problem) UpdateProblem(form ProblemForm) {
	err := db.First(&p)
	if err == nil {
		glog.Error("There was an error")
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
		glog.Error("There was an error")
	}
	return p
}

//GetAllProblemsWithLimit : Returns all problem objects
func GetAllProblemsWithLimit(limit int) []Problem {
	p := []Problem{}
	err := db.Limit(limit).Find(&p)
	if err == nil {
		glog.Error("There was an error")
	}
	return p
}

//GetAllProblemsByPage : Returns all problem objects
func GetAllProblemsByPage(page int) []Problem {
	p := []Problem{}
	err := db.Find(&p)
	if err == nil {
		glog.Error("There was an error")
	}

	//Check to see if there are more problems than util.PAGINGCONSTANT
	//if so set the pagingConstant to the number of problems
	var pageConstant = 0
	pagedProblems := []Problem{}
	if len(p) < util.PAGINGCONSTANT {
		pageConstant = len(p)
	} else {
		pageConstant = util.PAGINGCONSTANT
	}

	for i := (page * pageConstant); i < (pageConstant * (page + 1)); i++ {
		pagedProblems = append(pagedProblems, p[i])
	}

	return pagedProblems
}

//GetSubProblemsByID : Get all subproblems to a parent ID,
//add order them by highest rank
func GetSubProblemsByID(parentID int) []Problem {
	p := []Problem{}
	db.Where("parent_id = ?", parentID).Order("rank desc").Find(&p)
	return p
}

//QueryProblems : Return problems that are related to the query String
func QueryProblems(q string) []Problem {
	p := []Problem{}
	err := db.Where("title LIKE ?", "%"+q+"%").Order("rank desc").Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	return p
}

//QueryLast10Problems : Return problems that are related to the query String
func QueryLast10Problems(q string) []Problem {
	p := []Problem{}
	err := db.Limit(10).Where("title LIKE ?", "%"+q+"%").Order("rank desc").Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	return p
}

// //QueryProblems : Return problems that are related to the query String
// func QueryProblems(q string, pageNumber int) []Problem {
// 	p := []Problem{}
// 	err := db.Where("title LIKE ?", "%"+q+"%").Where("id < ? AND id > ?", util.PAGINGCONSTANT*pageNumber, util.PAGINGCONSTANT*(1+pageNumber)).Find(&p)
// 	if err == nil {
// 		glog.Info("There was an error")
// 	}
// 	return p
// }

//DeleteProblemByID : //DELETE
func DeleteProblemByID(form ProblemDeleteForm) error {
	p := Problem{}
	p.GetProblemByID(uint(form.ID))
	if p.OriginalPosterUsername == form.Username {
		db.Delete(&p)
		return nil
	}
	return errors.New("UnAuthorized User")
}

//VoteProblem : ~
func (p *Problem) VoteProblem(id int, vote bool) {
	err := db.Where("id = ?", id).First(&p)
	if err == nil {
		glog.Error("There was an error")
	}
	if vote == util.VOTEUP {
		p.Rank++
	} else {
		//Check if rank is below zero to prevent negatie votes
		if p.Rank <= 0 {
			return
		}
		p.Rank--
	}

	db.Model(&p).Update("rank", p.Rank)

	totalVotes := uint(0)
	problems := GetSubProblemsByID(int(p.ParentID))
	for i := 0; i < len(problems); i++ {
		totalVotes += problems[i].Rank
	}

	for i := 0; i < len(problems); i++ {
		var percentRank = float32(0.0)
		if totalVotes > 0 {
			//Verify percentRank is never below zero
			percentRank = float32(problems[i].Rank) / float32(totalVotes)
			db.Model(&problems[i]).Update("percent_rank", percentRank)
		}
	}
}
