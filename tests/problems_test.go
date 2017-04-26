package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"work/xprincipia/backend/gorm"
)

/*
Test Problems API
*/

func TestGetSolutionByID(t *testing.T) {
	//TODO : FIX expectedResult
	db := SetupTestingDB()
	defer db.Close()

	actualResult := gorm.Problem{}
	actualResult.GetProblemByID(1)
	expectedResult := gorm.Solution{
		ProblemID:              1,
		OriginalPosterUsername: "jackDaniels",
		Title: "The answer here is something very intense",
		Rank:  10,
	}

	assert.Equal(t, actualResult.OriginalPosterUsername, expectedResult.OriginalPosterUsername, "the OriginalPosterUsername's should be the same")
	assert.Equal(t, actualResult.Title, expectedResult.Title, "the Title's should be the same")
}
