package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"work/xprincipia/backend/util"
)

//Suggestion : Struct containing a question
type Suggestion struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string `gorm:"size:100000"`
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

//SuggestionDeleteForm : ~
type SuggestionDeleteForm struct {
	Username string
	ID       int
}

/*
API
*/

//CreateSuggestion : Creates a suggestion
func CreateSuggestion(form SuggestionForm) error {

	//Handle form Field Errors
	switch {
	case form.Description == "":
		return errors.New("Description is empty: Please fill in field")
	}

	//Create Suggestion
	s := Suggestion{}
	s.Username = form.Username
	intType, _ := strconv.Atoi(form.Type)
	s.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	s.TypeID = intTypeID
	s.Description = form.Description
	s.Rank = 0
	db.Create(&s)

	return nil
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
	err := db.Order("rank desc").Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}
	return s
}

//GetAllSuggestionsByTypeID :
func GetAllSuggestionsByTypeID(dataType int, typeID int) []Suggestion {
	s := []Suggestion{}
	err := db.Order("rank desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&s)
	if err == nil {
		glog.Info("There was an error")
	}

	return s
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

//DeleteSuggestionByID : //DELETE
func DeleteSuggestionByID(form SuggestionDeleteForm) error {
	s := Suggestion{}
	s.GetSuggestionByID(uint(form.ID))
	if s.Username == form.Username {
		db.Delete(&s)
		return nil
	}
	return errors.New("UnAuthorized User")
}

//VoteSuggestion : ~
func (s *Suggestion) VoteSuggestion(id int, vote bool) {
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
	suggestions := GetAllSuggestionsByTypeID(s.Type, s.TypeID)
	for i := 0; i < len(suggestions); i++ {
		totalVotes += suggestions[i].Rank
	}

	for i := 0; i < len(suggestions); i++ {
		var percentRank = float32(0.0)
		if totalVotes > 0 {
			percentRank = float32(suggestions[i].Rank) / float32(totalVotes)
		}
		db.Model(&suggestions[i]).Update("percent_rank", percentRank)
	}

}
