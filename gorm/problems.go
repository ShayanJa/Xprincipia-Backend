package gorm

import (
	"github.com/jinzhu/gorm"
)

// Problem : User generated problem
type Problem struct {
	gorm.Model
	OriginalPoster User
	Title          string
	Description    string
	SubProblems    []Problem
	Comments       []Comment
}
