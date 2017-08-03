package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"work/xprincipia/backend/util"
)

//Resource : Struct containing a pro
type Resource struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string `gorm:"size:10000"`
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
	r.Rank = 0
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

//GetAllResourcesByTypeID : Use typeID because resources are for both problems and solutions
func GetAllResourcesByTypeID(dataType int, typeID int) []Resource {
	r := []Resource{}
	err := db.Order("rank desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&r)
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

// UpdateResource : //UPDATE METHOD
func (r *Resource) UpdateResource(form ResourceForm) {
	err := db.First(&r)
	if err == nil {
		glog.Error("There was an error")
	}

	r.Description = form.Description
	db.Save(&r)
}

//VoteResource : ~
func (r *Resource) VoteResource(id int, vote bool) {
	err := db.Where("id = ?", id).Find(&r)
	if err == nil {
		glog.Info("There was an error")
	}
	//check if upVote or downVote
	if vote == util.VOTEUP {
		r.Rank++
	} else {
		r.Rank--
	}
	db.Model(&r).Update("rank", r.Rank)

	var totalVotes = 0
	resources := GetAllResourcesByTypeID(r.Type, r.TypeID)
	for i := 0; i < len(resources); i++ {
		totalVotes += resources[i].Rank
	}

	for i := 0; i < len(resources); i++ {
		var percentRank = float32(0.0)
		if totalVotes != 0 {
			percentRank = float32(resources[i].Rank) / float32(totalVotes)
		}
		db.Model(&resources[i]).Update("percent_rank", percentRank)
	}

}
