package gorm

import (
	"github.com/jinzhu/gorm"
)


type Comment struct {
    gorm.Model
    OriginalPoster: User
    Text: string
    Replies: []Comment
}

//API Functions