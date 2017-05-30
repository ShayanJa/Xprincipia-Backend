package gorm

import (
	"strconv"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
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
	c.Rank = 1
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
	err := db.Order("created_at desc").Find(&c)
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

//DeleteCommentByID : //DELETE
func DeleteCommentByID(id int) {
	c := Comment{}
	c.GetCommentByID(uint(id))
	db.Delete(&c)
}

//VoteComment : ~
func (c *Comment) VoteComment(id int) {
	err := db.Where("id = ?", id).Find(&c)
	if err == nil {
		glog.Info("There was an error")
	}
	c.Rank++
	db.Model(&c).Update("rank", c.Rank)

	var totalVotes = 0
	comments := GetAllCommentsBySuggestionID(c.SuggestionID)
	for i := 0; i < len(comments); i++ {
		totalVotes += comments[i].Rank
	}

	for i := 0; i < len(comments); i++ {
		var percentRank = float32(comments[i].Rank) / float32(totalVotes)
		db.Model(&comments[i]).Update("percent_rank", percentRank)
	}

}
