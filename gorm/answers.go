package gorm

import (
	"strconv"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Answer : Struct containing a question
type Answer struct {
	gorm.Model
	QuestionID  int
	Username    string
	Description string
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

//CreateAnswer : Creates a question
func CreateAnswer(form AnswerForm) {
	a := Answer{}
	a.Username = form.Username
	intQuestionID, _ := strconv.Atoi(form.QuestionID)
	a.QuestionID = intQuestionID
	a.Description = form.Description
	a.Rank = 1
	db.Create(&a)
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
	err := db.Order("created_at desc").Find(&a)
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

//DeleteAnswerByAnswerID : //DELETE
func DeleteAnswerByAnswerID(id int) {
	a := Answer{}
	a.GetAnswerByID(uint(id))
	db.Delete(&a)
}

//VoteAnswer : ~
func (a *Answer) VoteAnswer(id int) {
	err := db.Where("id = ?", id).Find(&a)
	if err == nil {
		glog.Info("There was an error")
	}
	a.Rank++
	db.Model(&a).Update("rank", a.Rank)

	var totalVotes = 0
	answers := GetAllAnswersByQuestionID(a.QuestionID)
	for i := 0; i < len(answers); i++ {
		totalVotes += answers[i].Rank
	}

	for i := 0; i < len(answers); i++ {
		var percentRank = float32(answers[i].Rank) / float32(totalVotes)
		db.Model(&answers[i]).Update("percent_rank", percentRank)
	}

}
