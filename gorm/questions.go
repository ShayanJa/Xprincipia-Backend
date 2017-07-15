package gorm

import (
	"strconv"

	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"work/xprincipia/backend/util"
)

//Question : Struct containing a question
type Question struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string
	Rank        int
	PercentRank float32
}

//QuestionForm : Form to make Question Struct
type QuestionForm struct {
	Type        string
	TypeID      string
	Username    string
	Description string
}

//QuestionDeleteForm : ~
type QuestionDeleteForm struct {
	Username string
	ID       int
}

/*
API
*/

//CreateQuestion : Creates a question
func CreateQuestion(form QuestionForm) {
	q := Question{}
	intType, _ := strconv.Atoi(form.Type)
	q.Type = intType
	intTypeID, _ := strconv.Atoi(form.TypeID)
	q.TypeID = intTypeID
	q.Username = form.Username
	q.Description = form.Description
	q.Rank = 1
	db.Create(&q)
}

//GetQuestionByID : Returns a Suggestion based on an int ID
func (q *Question) GetQuestionByID(id uint) {
	err := db.Where("id = ?", id).First(&q)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllQuestions : Return all Questions
func GetAllQuestions() []Question {
	q := []Question{}
	err := db.Order("rank desc").Find(&q)
	if err == nil {
		glog.Info("There was an error")
	}
	return q
}

//GetAllQuestionsByTypeID : Use typeID because questions are for both problems and solutions
func GetAllQuestionsByTypeID(dataType int, typeID int) []Question {
	q := []Question{}
	err := db.Order("rank desc").Where("type_id = ? AND type = ?", typeID, dataType).Find(&q)
	if err == nil {
		glog.Info("There was an error")
	}

	return q
}

//DeleteQuestionByID : //DELETE
func DeleteQuestionByID(form QuestionDeleteForm) error {
	q := Question{}
	q.GetQuestionByID(uint(form.ID))
	if q.Username == form.Username {
		db.Delete(&q)
		return nil
	}
	return errors.New("UnAuthorized User")
}

// UpdateQuestion : Updates a problem with problemForm as input
func (q *Question) UpdateQuestion(form QuestionForm) {
	err := db.First(&q)
	if err == nil {
		glog.Error("There was an error")
	}

	q.Description = form.Description
	db.Save(&q)
}

//VoteQuestion : ~
func (q *Question) VoteQuestion(id int, vote bool) {
	err := db.Where("id = ?", id).Find(&q)
	if err == nil {
		glog.Info("There was an error")
	}
	//check if upVote or downVote
	if vote == util.VOTEUP {
		q.Rank++
	} else {
		q.Rank--
	}
	db.Model(&q).Update("rank", q.Rank)

	var totalVotes = 0
	questions := GetAllQuestionsByTypeID(q.Type, q.TypeID)
	for i := 0; i < len(questions); i++ {
		totalVotes += questions[i].Rank
	}

	for i := 0; i < len(questions); i++ {
		var percentRank = float32(questions[i].Rank) / float32(totalVotes)
		db.Model(&questions[i]).Update("percent_rank", percentRank)
	}

}
