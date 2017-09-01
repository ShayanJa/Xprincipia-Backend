package gorm

import (
	"errors"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//Feedback : Struct containing a feedback
type Feedback struct {
	gorm.Model
	Username    string
	Description string `gorm:"size:100000"`
}

//FeedbackForm : Form to make Feedback Struct
type FeedbackForm struct {
	Username    string
	Description string
}

//FeedbackDeleteForm : ~
type FeedbackDeleteForm struct {
	Username string
	ID       int
}

/*
API
*/

//CreateFeedback : Creates a feedback
func CreateFeedback(form FeedbackForm) {
	f := Feedback{}
	f.Username = form.Username
	f.Description = form.Description
	db.Create(&f)
}

//GetFeedbackByID : Returns a Suggestion based on an int ID
func (f *Feedback) GetFeedbackByID(id uint) {
	err := db.Where("id = ?", id).First(&f)
	if err == nil {
		glog.Info("There was an error")
	}
}

//GetAllFeedback : Return all Feedbacks
func GetAllFeedback() []Feedback {
	f := []Feedback{}
	err := db.Order("created_at desc").Find(&f)
	if err == nil {
		glog.Info("There was an error")
	}
	return f
}

//DeleteFeedbackByID : //DELETE
func DeleteFeedbackByID(form FeedbackDeleteForm) error {
	f := Feedback{}
	f.GetFeedbackByID(uint(form.ID))
	if f.Username == form.Username {
		db.Delete(&f)
		return nil
	}
	return errors.New("UnAuthorized User")
}

// UpdateFeedback : Updates a problem with problemForm as input
func (f *Feedback) UpdateFeedback(form FeedbackForm) {
	err := db.First(&f)
	if err == nil {
		glog.Error("There was an error")
	}

	f.Description = form.Description
	db.Save(&f)
}
