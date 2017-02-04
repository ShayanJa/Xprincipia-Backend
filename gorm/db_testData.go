package gorm

import "github.com/jinzhu/gorm"
import "github.com/golang/glog"

func populateDBtestData(db *gorm.DB) {

	//User Creation
	shayan := User{
		FirstName:   "Shayan",
		LastName:    "Talebi",
		Email:       "Shay.talebi@gmail.com",
		Address:     "5515 Malibu Dr",
		Username:    "Shyshawn",
		PhoneNumber: "9528075184",
	}
	db.Create(&shayan)

	jackDaniels := User{
		FirstName:   "Jack",
		LastName:    "Daniels",
		Email:       "dev@xPrincipia.com",
		Address:     "5515 Malibu Dr",
		Username:    "jackDaniels",
		PhoneNumber: "9528015184",
	}
	db.Create(&jackDaniels)

	//Problem Creation
	problem := Problem{
		OriginalPoster: shayan,
		Title:          "How to save humanity",
		Description:    "I propose that we first sove humanity's deepest struggle",
	}
	db.Create(&problem)

	//Solution Creation
	solution := Solution{
		ProblemID:      problem.ID,
		OriginalPoster: jackDaniels,
		Text:           "The answer here is something very intense",
		Rating:         10,
	}
	glog.Info(solution.OriginalPoster)
	db.Create(&solution)

	solution2 := Solution{
		ProblemID:      problem.ID,
		OriginalPoster: shayan,
		Text:           "What we can do is enable users to try a solution",
		Rating:         10,
	}
	db.Create(&solution2)
	// solution2 := Solution{}
	// db.Create(&solution2)

	//Comment Creation
	comment := Comment{
		Type: PROBLEM,
		OP:   shayan,
		Text: "I understand what you are saying can you elaborate more",
	}
	comment1 := Comment{
		Type: PROBLEM,
		OP:   shayan,
		Text: "I just don't understand",
	}

	problem.MakeComment(comment)
	problem.MakeComment(comment1)
}
