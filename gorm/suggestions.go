package gorm

import (
	"strconv"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Suggestion : Struct containing a question
type Suggestion struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string
	Rank        int
}

//SuggestionForm : Form to make Question Struct
type SuggestionForm struct {
	Username    string
	Type        string
	TypeID      string
	Description string
}

/*
API
*/

//CreateSuggestion : Creates a question
func CreateSuggestion(form SuggestionForm) {
	s := Suggestion{}
	s.Username = form.Username
	intType, _ := strconv.Atoi(form.Type)
	s.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	s.TypeID = intTypeID
	s.Description = form.Description
	s.Rank = 1
	db.Create(&s)
}

//GetSuggestionByID : Returns a Suggestion based on an int ID
func (s *Suggestion) GetSuggestionByID(id uint) {
	err := db.Where("id = ?", id).First(&s)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllSuggestions : get all suggestions
func GetAllSuggestions() []Suggestion {
	s := []Suggestion{}
	err := db.Order("created_at desc").Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}
	return s
}

//GetAllSuggestionsByTypeID :
func GetAllSuggestionsByTypeID(dataType int, typeID int) []Suggestion {
	s := []Suggestion{}
	err := db.Order("created_at desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}

	return s
}
