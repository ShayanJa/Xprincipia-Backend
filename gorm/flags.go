package gorm

import "github.com/jinzhu/gorm"

//Flag : ~
type Flag struct {
	gorm.Model
	Type        int
	TypeID      int
	Username    string
	Description string `gorm:"size:10000"`
}

//FlagForm : ~
type FlagForm struct {
	Type        int
	TypeID      int
	Username    string
	Description string
}

//CreateFlag : ~
func CreateFlag(form FlagForm) {
	//Create Vote
	f := Flag{}
	f.Type = form.Type
	f.TypeID = form.TypeID
	f.Username = form.Username
	f.Description = form.Description

	db.Create(&f)
}
