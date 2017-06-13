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
	PercentRank float32
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
	return
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

//DeleteSuggestionByID : //DELETE
func DeleteSuggestionByID(id int) {
	s := Suggestion{}
	s.GetSuggestionByID(uint(id))
	db.Delete(&s)
}

//VoteSuggestion : ~
func (s *Suggestion) VoteSuggestion(id int) {
	err := db.Where("id = ?", id).Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}
	s.Rank++
	db.Model(&s).Update("rank", s.Rank)

	var totalVotes = 0
	suggestions := GetAllSuggestionsByTypeID(s.Type, s.TypeID)
	for i := 0; i < len(suggestions); i++ {
		totalVotes += suggestions[i].Rank
	}

	for i := 0; i < len(suggestions); i++ {
		var percentRank = float32(suggestions[i].Rank) / float32(totalVotes)
		db.Model(&suggestions[i]).Update("percent_rank", percentRank)
	}

}

// UpdateSuggestion : Updates a problem with problemForm as input
func (s *Suggestion) UpdateSuggestion(form SuggestionForm) {
	err := db.First(&s)
	if err == nil {
		glog.Error("There was an error")
	}

	s.Description = form.Description
	db.Save(&s)
}
