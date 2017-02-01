package gorm

import (
	"github.com/jinzhu/gorm"
)

// Comment : Generic Comment that can be used anywhere
type Comment struct {
	gorm.Model
	OriginalPoster User
	Text           string
	Replies        []Comment
}

//API Functions
