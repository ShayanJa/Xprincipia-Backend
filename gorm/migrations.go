package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

//migrations
func runMigrations(db *gorm.DB) {

	glog.Info("CREATING USER TABLE")
	db.AutoMigrate(&User{})

	glog.Info("CREATING PROBLEM TABLE")
	db.AutoMigrate(&Problem{})

	glog.Info("CREATING SOLUTION TABLE")
	db.AutoMigrate(&Solution{})

	glog.Info("CREATING STORY TABLE")
	db.AutoMigrate(&Story{})

	glog.Info("CREATING COMMENTS TABLE")
	db.AutoMigrate(&Comment{})

	glog.Info("CREATING ADDRESS TABLE")
	db.AutoMigrate(&Address{})

}
