package gorm

import (
	"strconv"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Question : Struct containing a question
type Question struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string
	// Answers     []string
	Rank int
}

//QuestionForm : Form to make Question Struct
type QuestionForm struct {
	Type        string
	TypeID      string
	Description string
}

/*
API
*/

//CreateQuestion : Creates a question
func CreateQuestion(form QuestionForm) {
	q := Question{}
	intType, _ := strconv.Atoi(form.Type)
	q.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	q.TypeID = intTypeID
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

//GetAllQuestions : Return all Questions
func GetAllQuestions() []Question {
	q := []Question{}
	err := db.Order("created_at desc").Find(&q)
	if err == nil {
		glog.Info("There was an error")
	}
	return q
}

//GetAllQuestionsByTypeID :
func GetAllQuestionsByTypeID(dataType int, typeID int) []Question {
	q := []Question{}
	err := db.Order("created_at desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&q)
	if err == nil {
		glog.Info("There was an error")
	}

	return q
}
