package gorm

import (
	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"strconv"
	"work/xprincipia/backend/util"
)

//Answer : Struct containing a question
type Answer struct {
	gorm.Model
	QuestionID  int
	Username    string
	Description string `gorm:"size:100000"`
	Rank        int
	PercentRank float32
}

//AnswerForm : Form to make Question Struct
type AnswerForm struct {
	Username    string
	QuestionID  string
	Description string
}

/*
API
*/

//CreateAnswer : Creates a answer
func CreateAnswer(form AnswerForm) error {

	//Handle form Field Errors
	switch {
	case form.Description == "":
		return errors.New("Description is empty: Please fill in field")
	}

	//Create Answer
	a := Answer{}
	a.Username = form.Username
	intQuestionID, _ := strconv.Atoi(form.QuestionID)
	a.QuestionID = intQuestionID
	a.Description = form.Description
	a.Rank = 0
	db.Create(&a)

	return nil
}

//GetAnswerByID : Returns a Suggestion based on an int ID
func (a *Answer) GetAnswerByID(id uint) {
	err := db.Where("id = ?", id).First(&a)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllAnswers : get all suggestions //READ
func GetAllAnswers() []Answer {
	a := []Answer{}
	err := db.Order("rank desc").Find(&a)
	if err == nil {
		glog.Info("There was an error")
	}
	return a
}

//GetAllAnswersByQuestionID : //READ
func GetAllAnswersByQuestionID(questionID int) []Answer {
	a := []Answer{}
	err := db.Order("created_at desc").Where("question_id = ?", questionID).Find(&a)
	if err == nil {
		glog.Info("There was an error")
	}

	return a
}

// UpdateAnswer : //UPDATE METHOD
func (a *Answer) UpdateAnswer(form AnswerForm) {
	err := db.First(&a)
	if err == nil {
		glog.Error("There was an error")
	}

	a.Description = form.Description
	db.Save(&a)
}

//DeleteAnswerByID : //DELETE
func DeleteAnswerByID(id int) {
	a := Answer{}
	a.GetAnswerByID(uint(id))
	db.Delete(&a)
}

//VoteAnswer : vote paramater takes in true or false to
//denote and upvote or a downvote
func (a *Answer) VoteAnswer(id int, vote bool) {
	err := db.Where("id = ?", id).Find(&a)
	if err == nil {
		glog.Info("There was an error")
	}
	//check if upVote or downVote
	if vote == util.VOTEUP {
		a.Rank++
	} else {
		a.Rank--
	}

	db.Model(&a).Update("rank", a.Rank)

	var totalVotes = 0
	answers := GetAllAnswersByQuestionID(a.QuestionID)
	for i := 0; i < len(answers); i++ {
		totalVotes += answers[i].Rank
	}

	for i := 0; i < len(answers); i++ {
		var percentRank = float32(0.0)
		if totalVotes > 0 {
			percentRank = float32(answers[i].Rank) / float32(totalVotes)
		}
		db.Model(&answers[i]).Update("percent_rank", percentRank)
	}

}
