package gorm

import "github.com/jinzhu/gorm"

//Suggestion : Struct containing a question
type Suggestion struct {
	gorm.Model
	Username    string
	Description string
	Rank        int
}

//SuggestionForm : Form to make Question Struct
type SuggestionForm struct {
	Description string
}

/*
API
*/

//CreateSuggestion : Creates a question
func CreateSuggestion(form QuestionForm) {
	q := Suggestion{}
	q.Description = form.Description
	q.Rank = 1
	db.Create(&q)
}
