package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//LearnContent : Struct containing a learnContent
type LearnContent struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string
	Rank        int
	PercentRank float32
}

//LearnContentForm : Form to make LearnContent Struct
type LearnContentForm struct {
	Type        string
	TypeID      string
	Username    string
	Description string
}

//LearnContentDeleteForm : ~
type LearnContentDeleteForm struct {
	Username string
	ID       int
}

/*
API
*/

//CreateLearnContent : Creates a learnContent
func CreateLearnContent(form LearnContentForm) {
	l := LearnContent{}
	intType, _ := strconv.Atoi(form.Type)
	l.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	l.TypeID = intTypeID
	l.Username = form.Username
	l.Description = form.Description
	l.Rank = 1
	db.Create(&l)
}

//GetLearnContentByID : Returns a Suggestion based on an int ID
func (l *LearnContent) GetLearnContentByID(id uint) {
	err := db.Where("id = ?", id).First(&l)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllLearnContents : Return all LearnContents
func GetAllLearnContents() []LearnContent {
	l := []LearnContent{}
	err := db.Order("created_at desc").Find(&l)
	if err == nil {
		glog.Info("There was an error")
	}
	return l
}

//GetAllLearnContentsByTypeID : Use typeID because questions are for both problems and solutions
func GetAllLearnContentsByTypeID(dataType int, typeID int) []LearnContent {
	l := []LearnContent{}
	err := db.Order("created_at desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&l)
	if err == nil {
		glog.Info("There was an error")
	}

	return l
}

//DeleteLearnContentByID : //DELETE
func DeleteLearnContentByID(form LearnContentDeleteForm) error {
	l := LearnContent{}
	l.GetLearnContentByID(uint(form.ID))
	if l.Username == form.Username {
		db.Delete(&l)
		return nil
	}
	return errors.New("UnAuthorized User")
}

//VoteLearnContent : ~
func (l *LearnContent) VoteLearnContent(id int) {
	err := db.Where("id = ?", id).Find(&l)
	if err == nil {
		glog.Info("There was an error")
	}
	l.Rank++
	db.Model(&l).Update("rank", l.Rank)

	var totalVotes = 0
	questions := GetAllLearnContentsByTypeID(l.Type, l.TypeID)
	for i := 0; i < len(questions); i++ {
		totalVotes += questions[i].Rank
	}

	for i := 0; i < len(questions); i++ {
		var percentRank = float32(questions[i].Rank) / float32(totalVotes)
		db.Model(&questions[i]).Update("percent_rank", percentRank)
	}

}
