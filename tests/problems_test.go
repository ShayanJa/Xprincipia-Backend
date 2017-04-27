package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"work/xprincipia/backend/gorm"
)

/*
Test Problems API
*/

func TestGetProblemByID(t *testing.T) {

	db := SetupTestingDB()
	defer db.Close()

	actualResult := gorm.Problem{}
	actualResult.GetProblemByID(1)

	//test with first problem from db_testdata
	expectedResult := gorm.Problem{
		ParentID:               0,
		OriginalPosterUsername: "Shyshawn",
		Title:       "What causes the emergent phenomena of consciousness?",
		Field:       "Consciousness",
		Summary:     "The name Amazon is said to arise from a war Francisco de Orellana fought with the Tapuyas and other tribes. The women of the tribe fought alongside the men, as was their custom.[3] Orellana derived the name Amazonas from the Amazons of Greek mythology, described by Herodotus and Diodorus.",
		Description: "The Amazon rainforest (Portuguese: Floresta Amazônica or Amazônia; Spanish: Selva Amazónica, Amazonía or usually Amazonia; French: Forêt amazonienne; Dutch: Amazoneregenwoud), also known in English as Amazonia or the Amazon Jungle, is a moist broadleaf forest that covers most of the Amazon basin of South America. This basin encompasses 7,000,000 square kilometres (2,700,000 sq mi), of which 5,500,000 square kilometres (2,100,000 sq mi) are covered by the rainforest. This region includes territory belonging to nine nations. The majority of the forest is contained within Brazil, with 60% of the rainforest, followed by Peru with 13%, Colombia with 10%, and with minor amounts in Venezuela, Ecuador, Bolivia, Guyana, Suriname and French Guiana. States or departments in four nations contain 'Amazonas' in their names. The Amazon represents over half of the planet's remaining rainforests,[1] and comprises the largest and most biodiverse tract of tropical rainforest in the world, with an estimated 390 billion individual trees divided into 16,000 species.[2]",
	}

	assert.Equal(t, expectedResult.OriginalPosterUsername, actualResult.OriginalPosterUsername, "the OriginalPosterUsername's should be the same")
	assert.Equal(t, expectedResult.Title, actualResult.Title, "the Title's should be the same")
	assert.Equal(t, expectedResult.Summary, actualResult.Summary, "the Summary's should be the same")
	assert.Equal(t, expectedResult.Field, actualResult.Field, "the Fields should be the same")
}

func TestCreateProblem_GetProblemByTitle_DeleteProblem(t *testing.T) {
	/*
	* Tests multiple API functions because it shortens code
	* CreateProblem
	* GetProblemByTitle
	* DeleteProblem
	*
	 */
	db := SetupTestingDB()
	defer db.Close()

	form := gorm.ProblemForm{
		Username: "Shyshawn",
		ParentID: "1",
		Title:    "Projectile vomiting in cats",
		Field:    "Aerospace",
		Summary:  "cats can projectile vomit pretty far",
	}

	//CreateProblem
	gorm.CreateProblem(form)

	//GetProblemByTitle
	actualResult := gorm.Problem{}
	actualResult.GetProblemByTitle(form.Title)

	assert.Equal(t, form.Title, actualResult.Title, "Titles's should be the same")
	assert.Equal(t, form.Field, actualResult.Field, "Field's should be the same")
	assert.Equal(t, form.Summary, actualResult.Summary, "Summary's should be the same")

	//DeleteProblem
	gorm.DeleteProblemByID(int(actualResult.ID))

	actualResult = gorm.Problem{}
	actualResult.GetProblemByTitle(form.Title)
	assert.NotEqual(t, form.Title, actualResult.Title, "Titles's should not be the same")
	assert.NotEqual(t, form.Field, actualResult.Field, "Field's should not be the same")
	assert.NotEqual(t, form.Summary, actualResult.Summary, "Summary's should not be the same")
}
