package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//migrations
func runMigrations(db *gorm.DB) bool {

	if db.HasTable(&User{}) {
		glog.Info("TABLES HAVE ALREADY BEEN CREATED...")
		return false
	}

	glog.Info("CREATING USER TABLE...")
	db.AutoMigrate(&User{})

	glog.Info("CREATING PROBLEM TABLE...")
	db.AutoMigrate(&Problem{})

	glog.Info("CREATING SOLUTION TABLE...")
	db.AutoMigrate(&Solution{})

	glog.Info("CREATING QUESTION TABLE...")
	db.AutoMigrate(&Question{})

	glog.Info("CREATING ANSWER TABLE...")
	db.AutoMigrate(&Answer{})

	glog.Info("CREATING SUGGESTION TABLE...")
	db.AutoMigrate(&Suggestion{})

	glog.Info("CREATING COMMENT TABLE...")
	db.AutoMigrate(&Comment{})

	glog.Info("CREATING FREEFORM TABLE...")
	db.AutoMigrate(&FreeForm{})

	glog.Info("CREATING PRO TABLE...")
	db.AutoMigrate(&Pro{})

	glog.Info("CREATING LEARN ITEMS TABLE...")
	db.AutoMigrate(&LearnContent{})

	glog.Info("CREATING RESOURCE TABLE...")
	db.AutoMigrate(&Resource{})

	glog.Info("CREATING VOTE TABLE...")
	db.AutoMigrate(&Vote{})

	glog.Info("CREATING LOGINATTEMPTS TABLE...")
	db.AutoMigrate(&LoginAttempt{})

	return true
}
