package gorm

import "github.com/jinzhu/gorm"

func populateDBtestData(db *gorm.DB) {

	solution := Solution{}
	db.FirstOrCreate(&solution)

}
