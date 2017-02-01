package gorm

import (
	"github.com/jinzhu/gorm"
)

// Story : Generic newfeed story
type Story struct {
	gorm.Model
	OriginalPoster User
	Title          string
	Description    string
}
