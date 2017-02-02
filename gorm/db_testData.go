package gorm

import "github.com/jinzhu/gorm"

func populateDBtestData(db *gorm.DB) {

	solution := Solution{}
	db.Create(&solution)
	solution1 := Solution{}
	db.Create(&solution1)
	solution2 := Solution{}
	db.Create(&solution2)

	problem1 := Problem{}
	db.Create(&problem1)

}
