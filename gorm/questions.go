package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Question : Struct containing a question
type Question struct {
	gorm.Model
	Username    string
	Description string
	// Answers     []string
	Rank int
}

//QuestionForm : Form to make Question Struct
type QuestionForm struct {
	Description string
}

/*
API
*/

//CreateQuestion : Creates a question
func CreateQuestion(form QuestionForm) {
	q := Question{}
	q.Description = form.Description
	q.Rank = 1
	db.Create(&q)
}

//GetQuestionByID : Returns a Suggestion based on an int ID
func (q *Question) GetQuestionByID(id uint) {
	err := db.Where("id = ?", id).First(&q)
	if err == nil {
		glog.Info("There was an error")
	}
}
