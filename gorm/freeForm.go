package gorm

import (
	"strconv"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"work/xprincipia/backend/util"
)

//FreeForm : Struct containing a question
type FreeForm struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string `gorm:"size:100000"`
	Rank        int
	PercentRank float32
}

//FreeFormForm : Form to make Question Struct
type FreeFormForm struct {
	Username    string
	Type        string
	TypeID      string
	Description string
}

/*
API
*/

//CreateFreeForm : Creates a question
func CreateFreeForm(form FreeFormForm) {
	f := FreeForm{}
	f.Username = form.Username
	intType, _ := strconv.Atoi(form.Type)
	f.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	f.TypeID = intTypeID
	f.Description = form.Description
	f.Rank = 0
	db.Create(&f)
	return
}

//GetFreeFormByID : Returns a FreeForm based on an int ID
func (f *FreeForm) GetFreeFormByID(id uint) {
	err := db.Where("id = ?", id).First(&f)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllFreeForms : get all suggestions
func GetAllFreeForms() []FreeForm {
	f := []FreeForm{}
	err := db.Order("rank desc").Find(&f)
	if err == nil {
		glog.Info("There was an error")
	}
	return f
}

//GetAllFreeFormsByTypeID :
func GetAllFreeFormsByTypeID(dataType int, typeID int) []FreeForm {
	f := []FreeForm{}
	err := db.Order("rank desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&f)
	if err == nil {
		glog.Info("There was an error")
	}

	return f
}

//DeleteFreeFormByID : //DELETE
func DeleteFreeFormByID(id int) {
	f := FreeForm{}
	f.GetFreeFormByID(uint(id))
	db.Delete(&f)
}

// UpdateFreeForm : Updates a problem with problemForm as input
func (f *FreeForm) UpdateFreeForm(form FreeFormForm) {
	err := db.First(&f)
	if err == nil {
		glog.Error("There was an error")
	}

	f.Description = form.Description
	db.Save(&f)
}

//VoteFreeForm : ~
func (f *FreeForm) VoteFreeForm(id int, vote bool) {
	err := db.Where("id = ?", id).Find(&f)
	if err == nil {
		glog.Info("There was an error")
	}

	//check if upVote or downVote
	if vote == util.VOTEUP {
		f.Rank++
	} else {
		f.Rank--
	}

	db.Model(&f).Update("rank", f.Rank)

	var totalVotes = 0
	freeForms := GetAllFreeFormsByTypeID(f.Type, f.TypeID)
	for i := 0; i < len(freeForms); i++ {
		totalVotes += freeForms[i].Rank
	}

	for i := 0; i < len(freeForms); i++ {
		var percentRank = float32(0.0)
		if totalVotes > 0 {
			percentRank = float32(freeForms[i].Rank) / float32(totalVotes)
		}
		db.Model(&freeForms[i]).Update("percent_rank", percentRank)
	}

}
