package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Resource : Struct containing a pro
type Resource struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string
	Rank        int
	PercentRank float32
}

//ResourceForm : Form to make Resource Struct
type ResourceForm struct {
	Type        string
	TypeID      string
	Username    string
	Description string
}

//ResourceDeleteForm : ~
type ResourceDeleteForm struct {
	Username string
	ID       int
}

/*
API
*/

//CreateResource : Creates a pro
func CreateResource(form ResourceForm) {
	r := Resource{}
	intType, _ := strconv.Atoi(form.Type)
	r.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	r.TypeID = intTypeID
	r.Username = form.Username
	r.Description = form.Description
	r.Rank = 1
	db.Create(&r)
}

//GetResourceByID : Returns a Suggestion based on an int ID
func (r *Resource) GetResourceByID(id uint) {
	err := db.Where("id = ?", id).First(&r)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllResources : Return all Resources
func GetAllResources() []Resource {
	r := []Resource{}
	err := db.Order("created_at desc").Find(&r)
	if err == nil {
		glog.Info("There was an error")
	}
	return r
}

//GetAllResourcesByTypeID : Use typeID because questions are for both problems and solutions
func GetAllResourcesByTypeID(dataType int, typeID int) []Resource {
	r := []Resource{}
	err := db.Order("created_at desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&r)
	if err == nil {
		glog.Info("There was an error")
	}

	return r
}

//DeleteResourceByID : //DELETE
func DeleteResourceByID(form ResourceDeleteForm) error {
	r := Resource{}
	r.GetResourceByID(uint(form.ID))
	if r.Username == form.Username {
		db.Delete(&r)
		return nil
	}
	return errors.New("UnAuthorized User")
}

//VoteResource : ~
func (r *Resource) VoteResource(id int) {
	err := db.Where("id = ?", id).Find(&r)
	if err == nil {
		glog.Info("There was an error")
	}
	r.Rank++
	db.Model(&r).Update("rank", r.Rank)

	var totalVotes = 0
	questions := GetAllResourcesByTypeID(r.Type, r.TypeID)
	for i := 0; i < len(questions); i++ {
		totalVotes += questions[i].Rank
	}

	for i := 0; i < len(questions); i++ {
		var percentRank = float32(questions[i].Rank) / float32(totalVotes)
		db.Model(&questions[i]).Update("percent_rank", percentRank)
	}

}
