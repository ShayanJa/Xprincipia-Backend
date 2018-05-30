package gorm

import "github.com/jinzhu/gorm"
import "github.com/golang/glog"
import "../util"

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

	//Check type of vote
	switch {
	case v.Type == util.SOLUTION:
		s := Solution{}
		s.VoteSolution(v.TypeID, util.VOTEUP)
	case v.Type == util.PROBLEM:
		p := Problem{}
		p.VoteProblem(v.TypeID, util.VOTEUP)
	case v.Type == util.QUESTION:
		q := Question{}
		q.VoteQuestion(v.TypeID, util.VOTEUP)
	case v.Type == util.SUGGESTION:
		s := Suggestion{}
		s.VoteSuggestion(v.TypeID, util.VOTEUP)
	case v.Type == util.ANSWER:
		a := Answer{}
		a.VoteAnswer(v.TypeID, util.VOTEUP)
	case v.Type == util.COMMENT:
		c := Comment{}
		c.VoteComment(v.TypeID, util.VOTEUP)
	case v.Type == util.FREEFORM:
		f := FreeForm{}
		f.VoteFreeForm(v.TypeID, util.VOTEUP)
	case v.Type == util.LEARNITEM:
		l := LearnItem{}
		l.VoteLearnItem(v.TypeID, util.VOTEUP)
	case v.Type == util.RESOURCE:
		r := Resource{}
		r.VoteResource(v.TypeID, util.VOTEUP)
	case v.Type == util.PRO:
		p := Pro{}
		p.VotePro(v.TypeID, util.VOTEUP)
	case v.Type == util.CON:
		r := Con{}
		r.VoteCon(v.TypeID, util.VOTEUP)
	}

	return true

}

//GetNumberOfVotesByTypeID : !
func GetNumberOfVotesByTypeID(id int) int {
	v := []Vote{}
	db.Where("type_id = ?", id).Find(&v)
	return len(v)
}

// GetVotesByTypeID(typeID, type){
// 	v := []

// }

// IsVotedOn : ~
func IsVotedOn(Type int, typeID int, username string) bool {
	v := Vote{}
	db.Where("type_id = ? AND type = ? AND username = ?", typeID, Type, username).First(&v)

	if v.ID != 0 {
		return true
	}
	return false
}

//DeleteVote : Delete vote
func DeleteVote(Type int, typeID int, username string) {
	v := Vote{}
	db.Where("type_id = ? AND type = ? AND username = ?", typeID, Type, username).First(&v)
	if v.ID != 0 {
		db.Delete(&v)

	}
	//Check type of vote
	switch {
	case v.Type == util.SOLUTION:
		s := Solution{}
		s.VoteSolution(v.TypeID, util.VOTEDOWN)
	case v.Type == util.PROBLEM:
		p := Problem{}
		p.VoteProblem(v.TypeID, util.VOTEDOWN)
	case v.Type == util.QUESTION:
		q := Question{}
		q.VoteQuestion(v.TypeID, util.VOTEDOWN)
	case v.Type == util.SUGGESTION:
		s := Suggestion{}
		s.VoteSuggestion(v.TypeID, util.VOTEDOWN)
	case v.Type == util.ANSWER:
		a := Answer{}
		a.VoteAnswer(v.TypeID, util.VOTEDOWN)
	case v.Type == util.COMMENT:
		c := Comment{}
		c.VoteComment(v.TypeID, util.VOTEDOWN)
	case v.Type == util.FREEFORM:
		f := FreeForm{}
		f.VoteFreeForm(v.TypeID, util.VOTEDOWN)
	case v.Type == util.LEARNITEM:
		l := LearnItem{}
		l.VoteLearnItem(v.TypeID, util.VOTEDOWN)
	case v.Type == util.RESOURCE:
		r := Resource{}
		r.VoteResource(v.TypeID, util.VOTEDOWN)
	case v.Type == util.PRO:
		p := Pro{}
		p.VotePro(v.TypeID, util.VOTEDOWN)
	case v.Type == util.CON:
		r := Con{}
		r.VoteCon(v.TypeID, util.VOTEDOWN)
	}

	return
}
