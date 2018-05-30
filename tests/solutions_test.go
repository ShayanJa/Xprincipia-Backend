package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"../gorm"
)

/*
Test Solutions API
*/

//Test GetSolutionByID function in Gorm
func TestGetSolutionByID(t *testing.T) {

	db := SetupTestingDB()
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
	assert.Equal(t, actualResult.Title, expectedResult.Title, "the Title's should be the same")
}

func TestGetSolutionsByProblemID(t *testing.T) {

	db := SetupTestingDB()
	defer db.Close()

	actualResult := gorm.GetSolutionsByProblemID(1)
	expectedResult := []gorm.Solution{}
	s := gorm.Solution{}
	s.GetSolutionByID(1)
	expectedResult = append(expectedResult, s)
	s = gorm.Solution{}
	s.GetSolutionByID(2)
	expectedResult = append(expectedResult, s)
	s = gorm.Solution{}
	s.GetSolutionByID(3)
	expectedResult = append(expectedResult, s)

	assert.Equal(t, actualResult, expectedResult, "The should be 3 solutions with the same values")

}

func TestGetAllSolutions(t *testing.T) {

	db := SetupTestingDB()
	defer db.Close()

	actualResult := gorm.GetAllSolutions()
	expectedResult := []gorm.Solution{}
	s := gorm.Solution{}
	s.GetSolutionByID(1)
	expectedResult = append(expectedResult, s)
	s = gorm.Solution{}
	s.GetSolutionByID(2)
	expectedResult = append(expectedResult, s)
	s = gorm.Solution{}
	s.GetSolutionByID(3)
	expectedResult = append(expectedResult, s)
	s = gorm.Solution{}
	s.GetSolutionByID(4)
	expectedResult = append(expectedResult, s)

	assert.Equal(t, actualResult, expectedResult, "The should be 4 solutions with the same values as expected")

}

//TODO
func TestVoteSolution(t *testing.T) {
	//TODO
	db := SetupTestingDB()
	defer db.Close()
}

func DeleteSolutionByID(t *testing.T) {
	//TODO
	db := SetupTestingDB()
	defer db.Close()
}
