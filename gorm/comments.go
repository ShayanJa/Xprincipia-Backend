package gorm

import (
	"github.com/jinzhu/gorm"
)


type Comment struct {
    gorm.Model
    OriginalPoster: gorm.User
    Text: string
    Replies: []gorm.Comment
}

//API Functions