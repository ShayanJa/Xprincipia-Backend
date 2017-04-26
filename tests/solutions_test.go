package tests

import (
	database "github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"work/xprincipia/backend/gorm"
)

func setupTestingDB() *database.DB {
	err := godotenv.Load("../config/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := gorm.InitializeDB()
	return db
}

//test Solutions api
func TestGetSolutionByID(t *testing.T) {

	db := setupTestingDB()
	defer db.Close()

	actualResult := gorm.Solution{}
	actualResult.GetSolutionByID(1)
	expectedResult := gorm.Solution{
		ProblemID:              1,
		OriginalPosterUsername: "jackDaniels",
		Title: "The answer here is something very intense",
		Rank:  10,
	}

	assert.Equal(t, actualResult.ProblemID, expectedResult.ProblemID, "the problem ID's should be the same")
	assert.Equal(t, actualResult.OriginalPosterUsername, expectedResult.OriginalPosterUsername, "the OriginalPosterUsername's should be the same")
}
