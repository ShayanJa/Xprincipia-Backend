package gorm

import "github.com/jinzhu/gorm"

//Question : Struct containing a question
type Question struct {
	gorm.Model
	Username    string
	Description string
	// Answers     []string
	Rank int
}

//QuestionForm : Form to make Question Struct
type QuestionForm struct {
	Description string
}

/*
API
*/

//CreateQuestion : Creates a question
func CreateQuestion(form QuestionForm) {
	q := Question{}
	q.Description = form.Description
	q.Rank = 1
	db.Create(&q)
}
