package gorm

import (
	"github.com/jinzhu/gorm"
)

// COMMENT TYPES

// PROBLEM : ~
const PROBLEM = 0

// SOLUTION : ~
const SOLUTION = 1

// Comment : Generic Comment that can be used anywhere
type Comment struct {
	gorm.Model
	Type    int
	TypeID  uint
	OP      User
	Text    string
	Replies []Comment
}

//API Functions
