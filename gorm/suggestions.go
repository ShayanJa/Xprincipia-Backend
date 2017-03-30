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
func CreateSuggestion(form SuggestionForm) {
	s := Suggestion{}
	s.Description = form.Description
	s.Rank = 1
	db.Create(&s)
}
