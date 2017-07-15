package gorm

import (
	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"strconv"
	"work/xprincipia/backend/util"
)

//Comment : Struct containing a question
type Comment struct {
	gorm.Model
	SuggestionID int
	Username     string
	Description  string
	Rank         int
	PercentRank  float32
}

//CommentForm : Form to make Question Struct
type CommentForm struct {
	Username     string
	SuggestionID string
	Description  string
}

//CommentDeleteForm : ~
type CommentDeleteForm struct {
	Username string
	ID       int
}

/*
API
*/

//CreateComment : Creates a question
func CreateComment(form CommentForm) {
	c := Comment{}
	c.Username = form.Username
	intCommentID, _ := strconv.Atoi(form.SuggestionID)
	c.SuggestionID = intCommentID
	c.Description = form.Description
	c.Rank = 0
	db.Create(&c)
}

//GetCommentByID : Returns a Suggestion based on an int ID
func (c *Comment) GetCommentByID(id uint) {
	err := db.Where("id = ?", id).First(&c)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllComments : get all suggestions
func GetAllComments() []Comment {
	c := []Comment{}
	err := db.Order("rank desc").Find(&c)
	if err == nil {
		glog.Info("There was an error")
	}
	return c
}

//GetAllCommentsBySuggestionID :
func GetAllCommentsBySuggestionID(suggestionID int) []Comment {
	c := []Comment{}
	err := db.Order("created_at desc").Where("suggestion_id = ?", suggestionID).Find(&c)
	if err == nil {
		glog.Info("There was an error")
	}

	return c
}

// UpdateComment : Updates a problem with problemForm as input
func (c *Comment) UpdateComment(form CommentForm) {
	err := db.First(&c)
	if err == nil {
		glog.Error("There was an error")
	}

	c.Description = form.Description
	db.Save(&c)
}

//DeleteCommentByID : //DELETE
func DeleteCommentByID(form CommentDeleteForm) error {
	c := Comment{}
	c.GetCommentByID(uint(form.ID))
	if c.Username == form.Username {
		db.Delete(&c)
		return nil
	}
	return errors.New("UnAuthorized User")
}

//VoteComment : ~
func (c *Comment) VoteComment(id int, vote bool) {
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
	comments := GetAllCommentsBySuggestionID(c.SuggestionID)
	for i := 0; i < len(comments); i++ {
		totalVotes += comments[i].Rank
	}

	for i := 0; i < len(comments); i++ {
		var percentRank = float32(0.0)
		if totalVotes != 0 {
			percentRank = float32(comments[i].Rank) / float32(totalVotes)
		}
		db.Model(&comments[i]).Update("percent_rank", percentRank)
	}

}
