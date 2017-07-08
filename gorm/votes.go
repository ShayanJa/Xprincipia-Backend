package gorm

import "github.com/jinzhu/gorm"
import "github.com/golang/glog"
import "work/xprincipia/backend/util"

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
		s.VoteSolution(v.TypeID)
	case v.Type == util.PROBLEM:
		p := Problem{}
		p.VoteProblem(v.TypeID)
	case v.Type == util.QUESTION:
		q := Question{}
		q.VoteQuestion(v.TypeID)
	case v.Type == util.SUGGESTION:
		s := Suggestion{}
		s.VoteSuggestion(v.TypeID)
	case v.Type == util.ANSWER:
		a := Answer{}
		a.VoteAnswer(v.TypeID)
	case v.Type == util.COMMENT:
		c := Comment{}
		c.VoteComment(v.TypeID)
	case v.Type == util.FREEFORM:
		f := FreeForm{}
		f.VoteFreeForm(v.TypeID)
	case v.Type == util.LEARNITEM:
		l := LearnItem{}
		l.VoteLearnItem(v.TypeID)
	case v.Type == util.RESOURCE:
		r := Resource{}
		r.VoteResource(v.TypeID)
	case v.Type == util.PRO:
		p := Pro{}
		p.VotePro(v.TypeID)
	case v.Type == util.CON:
		r := Con{}
		r.VoteCon(v.TypeID)
	}

	return true

}

//GetNumberOfVotesByTypeID : !
func GetNumberOfVotesByTypeID(id int) int {
	v := []Vote{}
	db.Where("type_id = ?", id).Find(&v)
	return len(v)
}
