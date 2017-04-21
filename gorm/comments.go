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

//GetAnswerByID : Returns a Suggestion based on an int ID
func (a *Comment) GetAnswerByID(id uint) {
	err := db.Where("id = ?", id).First(&a)
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
	err := db.Order("created_at desc").Where("question_id = ?", suggestionID).Find(&c)
	if err == nil {
		glog.Info("There was an error")
	}

	return c
}
