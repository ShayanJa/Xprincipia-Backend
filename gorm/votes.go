package gorm

import "github.com/jinzhu/gorm"
import "github.com/golang/glog"

//Vote : ~
type Vote struct {
	gorm.Model
	Type     int
	TypeID   int
	Username string
}

//VoteForm : ~
type VoteForm struct {
	Type     int
	TypeID   int
	Username string
}

//CreateVote : ~
func CreateVote(form VoteForm) bool {
	//Create Vote
	v := Vote{}
	v.Type = form.Type
	v.TypeID = form.TypeID
	v.Username = form.Username

	foundVotes := []Vote{}
	db.Where("username = ? AND type_id = ? AND type = ?", v.Username, v.TypeID, v.Type).Find(&foundVotes)
	glog.Info(foundVotes)
	if len(foundVotes) > 0 {
		return false
	}
	db.Create(&v)

	//Change solution rank
	if v.Type == 1 {
		s := Solution{}
		s.VoteSolution(v.TypeID)
	} else {
		if v.Type == 0 {
			p := Problem{}
			p.VoteProblem(v.TypeID)
		}
	}
	return true

}

//GetNumberOfVotesByTypeID : !
func GetNumberOfVotesByTypeID(id int) int {
	v := []Vote{}
	db.Where("type_id = ?", id).Find(&v)
	return len(v)
}
