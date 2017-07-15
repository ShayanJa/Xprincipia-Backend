package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"work/xprincipia/backend/util"
)

// Solution : Generic Problem Solution
type Solution struct {
	gorm.Model
	ProblemID              uint
	OriginalPoster         User `gorm:"ForeignKey:OriginalPosterUsername;AssociationForeignKey:Username" json:"originalPoster" form:"originalPoster"`
	OriginalPosterUsername string
	Title                  string `gorm:"size:151"`
	Summary                string `gorm:"size:1500"`
	Description            string `gorm:"size:10000"`
	Evidence               string `gorm:"size:1500"`
	Experiments            string `gorm:"size:1500"`
	References             string `gorm:"size:1500"`
	Rank                   int
	PercentRank            float32
	Suggestions            []Suggestion
	Questions              []Question
}

//SolutionForm : Solution Form
type SolutionForm struct {
	Username    string `json:"username" form:"username"`
	ProblemID   string `json:"problemID" form:"problemID"`
	Title       string `json:"title" form:"title"`
	Summary     string `json:"summary" form:"summary"`
	Description string `json:"description" form:"description"`
	Evidence    string `json:"evidence" form:"evidence"`
	Experiments string `json:"experiments" form:"experiments"`
	References  string `json:"references" form:"references"`
}

//SolutionDeleteForm : ~
type SolutionDeleteForm struct {
	Username string
	ID       int
}

//SolutionVoteForm : Vote form
type SolutionVoteForm struct {
	Username   string `json:"username" form:"username"`
	SolutionID string `json:"solutionID" form:"soutionID"`
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
	err := db.Where("problem_id = ?", id).Order("rank desc").Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}
	return s
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
	intID, _ := strconv.Atoi(form.ProblemID)
	s.ProblemID = uint(intID)
	s.OriginalPosterUsername = form.Username
	s.Title = form.Title
	s.Summary = form.Summary
	s.Description = form.Description
	s.Evidence = form.Evidence
	s.Experiments = form.Experiments
	s.References = form.References
	s.Rank = 1

	db.Create(&s)
}

// UpdateSolution : Updates a problem with problemForm as input
func (s *Solution) UpdateSolution(form SolutionForm) {
	err := db.First(&s)
	if err == nil {
		glog.Error("There was an error")
	}

	s.Description = form.Description
	s.Summary = form.Summary
	s.Title = form.Title
	db.Save(&s)
}

//DeleteSolutionByID : //DELETE
func DeleteSolutionByID(form SolutionDeleteForm) error {
	p := Solution{}
	p.GetSolutionByID(uint(form.ID))
	if p.OriginalPosterUsername == form.Username {
		db.Delete(&p)
		return nil
	}
	return errors.New("UnAuthorized User")
}

//VoteSolution : ~
func (s *Solution) VoteSolution(id int, vote bool) {
	err := db.Where("id = ?", id).Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}
	//check if upVote or downVote
	if vote == util.VOTEUP {
		s.Rank++
	} else {
		s.Rank--
	}
	db.Model(&s).Update("rank", s.Rank)

	var totalVotes = 0
	solutions := GetSolutionsByProblemID(int(s.ProblemID))
	for i := 0; i < len(solutions); i++ {
		totalVotes += solutions[i].Rank
	}

	for i := 0; i < len(solutions); i++ {
		var percentRank = float32(solutions[i].Rank) / float32(totalVotes)
		db.Model(&solutions[i]).Update("percent_rank", percentRank)
	}

}
