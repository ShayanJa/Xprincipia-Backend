package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Pro : Struct containing a pro
type Pro struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string
	Rank        int
	PercentRank float32
}

//ProForm : Form to make Pro Struct
type ProForm struct {
	Type        string
	TypeID      string
	Username    string
	Description string
}

//ProDeleteForm : ~
type ProDeleteForm struct {
	Username string
	ID       int
}

/*
API
*/

//CreatePro : Creates a pro
func CreatePro(form ProForm) {
	p := Pro{}
	intType, _ := strconv.Atoi(form.Type)
	p.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	p.TypeID = intTypeID
	p.Username = form.Username
	p.Description = form.Description
	p.Rank = 1
	db.Create(&p)
}

//GetProByID : Returns a Pro based on an int ID
func (p *Pro) GetProByID(id uint) {
	err := db.Where("id = ?", id).First(&p)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllPros : Return all Pros
func GetAllPros() []Pro {
	p := []Pro{}
	err := db.Order("created_at desc").Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	return p
}

//GetAllProsByTypeID : Use typeID because questions are for both problems and solutions
func GetAllProsByTypeID(dataType int, typeID int) []Pro {
	p := []Pro{}
	err := db.Order("created_at desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}

	return p
}

//DeleteProByID : //DELETE
func DeleteProByID(form ProDeleteForm) error {
	p := Pro{}
	p.GetProByID(uint(form.ID))
	if p.Username == form.Username {
		db.Delete(&p)
		return nil
	}
	return errors.New("UnAuthorized User")
}

// UpdatePro : Updates a problem with problemForm as input
func (p *Pro) UpdatePro(form ProForm) {
	err := db.First(&p)
	if err == nil {
		glog.Error("There was an error")
	}

	p.Description = form.Description
	db.Save(&p)
}

//VotePro : ~
func (p *Pro) VotePro(id int) {
	err := db.Where("id = ?", id).Find(&p)
	if err == nil {
		glog.Info("There was an error")
	}
	p.Rank++
	db.Model(&p).Update("rank", p.Rank)

	var totalVotes = 0
	questions := GetAllProsByTypeID(p.Type, p.TypeID)
	for i := 0; i < len(questions); i++ {
		totalVotes += questions[i].Rank
	}

	for i := 0; i < len(questions); i++ {
		var percentRank = float32(questions[i].Rank) / float32(totalVotes)
		db.Model(&questions[i]).Update("percent_rank", percentRank)
	}

}
