package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Con : Struct containing a con
type Con struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string
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
	p := Con{}
	intType, _ := strconv.Atoi(form.Type)
	p.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	p.TypeID = intTypeID
	p.Username = form.Username
	p.Description = form.Description
	p.Rank = 1
	db.Create(&p)
}

//GetConByID : Returns a Suggestion based on an int ID
func (p *Con) GetConByID(id uint) {
	err := db.Where("id = ?", id).First(&p)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllCons : Return all Cons
func GetAllCons() []Con {
	p := []Con{}
	err := db.Order("created_at desc").Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	return p
}

//GetAllConsByTypeID : Use typeID because questions are for both problems and solutions
func GetAllConsByTypeID(dataType int, typeID int) []Con {
	p := []Con{}
	err := db.Order("created_at desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}

	return p
}

//DeleteConByID : //DELETE
func DeleteConByID(form ConDeleteForm) error {
	p := Con{}
	p.GetConByID(uint(form.ID))
	if p.Username == form.Username {
		db.Delete(&p)
		return nil
	}
	return errors.New("UnAuthorized User")
}

//VoteCon : ~
func (p *Con) VoteCon(id int) {
	err := db.Where("id = ?", id).Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	p.Rank++
	db.Model(&p).Update("rank", p.Rank)

	var totalVotes = 0
	questions := GetAllConsByTypeID(p.Type, p.TypeID)
	for i := 0; i < len(questions); i++ {
		totalVotes += questions[i].Rank
	}

	for i := 0; i < len(questions); i++ {
		var percentRank = float32(questions[i].Rank) / float32(totalVotes)
		db.Model(&questions[i]).Update("percent_rank", percentRank)
	}

}
