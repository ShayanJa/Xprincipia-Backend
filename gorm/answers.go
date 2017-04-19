package gorm

import (
	"strconv"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Answer : Struct containing a question
type Answer struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string
	Rank        int
}

//AnswerForm : Form to make Question Struct
type AnswerForm struct {
	Username    string
	Type        string
	TypeID      string
	Description string
}

/*
API
*/

//CreateAnswer : Creates a question
func CreateAnswer(form AnswerForm) {
	a := Answer{}
	a.Username = form.Username
	intType, _ := strconv.Atoi(form.Type)
	a.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	a.TypeID = intTypeID
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

//GetAllAnswers : get all suggestions
func GetAllAnswers() []Answer {
	a := []Answer{}
	err := db.Order("created_at desc").Find(&a)
	if err == nil {
		glog.Info("There was an error")
	}
	return a
}

//GetAllAnswersByTypeID :
func GetAllAnswersByTypeID(dataType int, typeID int) []Answer {
	a := []Answer{}
	err := db.Order("created_at desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&a)
	if err == nil {
		glog.Info("There was an error")
	}

	return a
}
