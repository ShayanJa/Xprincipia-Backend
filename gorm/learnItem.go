package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"../util"
)

//LearnItem : Struct containing a learnItem
type LearnItem struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string `gorm:"size:100000"`
	Rank        int
	PercentRank float32
}

// LearnItemForm : Form to make LearnItem Struct
type LearnItemForm struct {
	Type        string
	TypeID      string
	Username    string
	Description string
}

//LearnItemDeleteForm : ~
type LearnItemDeleteForm struct {
	Username string
	ID       int
}

/*
API
*/

//CreateLearnItem : Creates a learnItem
func CreateLearnItem(form LearnItemForm) {
	l := LearnItem{}
	intType, _ := strconv.Atoi(form.Type)
	l.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	l.TypeID = intTypeID
	l.Username = form.Username
	l.Description = form.Description
	l.Rank = 0
	db.Create(&l)
}

//GetLearnItemByID : Returns a Suggestion based on an int ID
func (l *LearnItem) GetLearnItemByID(id uint) {
	err := db.Where("id = ?", id).First(&l)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllLearnItems : Return all LearnItems
func GetAllLearnItems() []LearnItem {
	l := []LearnItem{}
	err := db.Order("created_at desc").Find(&l)
	if err == nil {
		glog.Info("There was an error")
	}
	return l
}

//GetAllLearnItemsByTypeID : Use typeID because questions are for both problems and solutions
func GetAllLearnItemsByTypeID(dataType int, typeID int) []LearnItem {
	l := []LearnItem{}
	err := db.Order("rank desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&l)
	if err == nil {
		glog.Info("There was an error")
	}

	return l
}

//DeleteLearnItemByID : //DELETE
func DeleteLearnItemByID(form LearnItemDeleteForm) error {
	l := LearnItem{}
	l.GetLearnItemByID(uint(form.ID))
	if l.Username == form.Username {
		db.Delete(&l)
		return nil
	}
	return errors.New("UnAuthorized User")
}

// UpdateLearnItem : Updates a problem with problemForm as input
func (l *LearnItem) UpdateLearnItem(form LearnItemForm) {
	err := db.First(&l)
	if err == nil {
		glog.Error("There was an error")
	}

	l.Description = form.Description
	db.Save(&l)
}

//VoteLearnItem : ~
func (l *LearnItem) VoteLearnItem(id int, vote bool) {
	err := db.Where("id = ?", id).Find(&l)
	if err == nil {
		glog.Info("There was an error")
	}

	//check if upVote or downVote
	if vote == util.VOTEUP {
		l.Rank++
	} else {
		l.Rank--
	}

	db.Model(&l).Update("rank", l.Rank)

	var totalVotes = 0
	learnItems := GetAllLearnItemsByTypeID(l.Type, l.TypeID)
	for i := 0; i < len(learnItems); i++ {
		totalVotes += learnItems[i].Rank
	}

	for i := 0; i < len(learnItems); i++ {
		var percentRank = float32(0.0)
		if totalVotes != 0 {
			percentRank = float32(learnItems[i].Rank) / float32(totalVotes)
		}
		db.Model(&learnItems[i]).Update("percent_rank", percentRank)
	}

}
