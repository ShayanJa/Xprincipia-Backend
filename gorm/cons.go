package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"../util"
)

//Con : Struct containing a con
type Con struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string `gorm:"size:100000"`
	Rank        int
	PercentRank float32
}

//ConForm : Form to make Con Struct
type ConForm struct {
	Type        string
	TypeID      string
	Username    string
	Description string
}

//ConDeleteForm : ~
type ConDeleteForm struct {
	Username string
	ID       int
}

/*
API
*/

//CreateCon : Creates a con
func CreateCon(form ConForm) {
	c := Con{}
	intType, _ := strconv.Atoi(form.Type)
	c.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	c.TypeID = intTypeID
	c.Username = form.Username
	c.Description = form.Description
	c.Rank = 0
	db.Create(&c)
}

//GetConByID : Returns a Suggestion based on an int ID
func (c *Con) GetConByID(id uint) {
	err := db.Where("id = ?", id).First(&c)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllCons : Return all Cons
func GetAllCons() []Con {
	c := []Con{}
	err := db.Order("rank desc").Find(&c)
	if err == nil {
		glog.Info("There was an error")
	}
	return c
}

//GetAllConsByTypeID : Use typeID because questions are for both problems and solutions
func GetAllConsByTypeID(dataType int, typeID int) []Con {
	c := []Con{}
	err := db.Order("created_at desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&c)
	if err == nil {
		glog.Info("There was an error")
	}

	return c
}

//DeleteConByID : //DELETE
func DeleteConByID(form ConDeleteForm) error {
	c := Con{}
	c.GetConByID(uint(form.ID))
	if c.Username == form.Username {
		db.Delete(&c)
		return nil
	}
	return errors.New("UnAuthorized User")
}

// UpdateCon : Updates a problem with problemForm as input
func (c *Con) UpdateCon(form ConForm) {
	err := db.First(&c)
	if err == nil {
		glog.Error("There was an error")
	}

	c.Description = form.Description
	db.Save(&c)
}

//VoteCon : ~
func (c *Con) VoteCon(id int, vote bool) {
	err := db.Where("id = ?", id).Find(&c)
	if err == nil {
		glog.Info("There was an error")
	}
	//check if upVote or downVote
	if vote == util.VOTEUP {
		c.Rank++
	} else {
		c.Rank--
	}
	db.Model(&c).Update("rank", c.Rank)

	var totalVotes = 0
	questions := GetAllConsByTypeID(c.Type, c.TypeID)
	for i := 0; i < len(questions); i++ {
		totalVotes += questions[i].Rank
	}

	for i := 0; i < len(questions); i++ {
		var percentRank = float32(0.0)
		if totalVotes != 0 {
			percentRank = float32(questions[i].Rank) / float32(totalVotes)
		}
		db.Model(&questions[i]).Update("percent_rank", percentRank)
	}

}
