package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func populateDBtestData(db *gorm.DB) {
	//HashedPasswords
	passwordBytes := []byte("Popcan123")
	hashedShyshawnPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	//User Creation
	shayan := User{
		FirstName:      "Shayan",
		LastName:       "Talebi",
		Email:          "Shay.talebi@gmail.com",
		Address:        "5515 Malibu Dr",
		Username:       "Shyshawn",
		PhoneNumber:    "9528075184",
		HashedPassword: hashedShyshawnPassword,
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
		Title:          "What causes the emergent phenomena of consciousness?",
		Summary:        "The name Amazon is said to arise from a war Francisco de Orellana fought with the Tapuyas and other tribes. The women of the tribe fought alongside the men, as was their custom.[3] Orellana derived the name Amazonas from the Amazons of Greek mythology, described by Herodotus and Diodorus.",
		Description:    "The Amazon rainforest (Portuguese: Floresta Amazônica or Amazônia; Spanish: Selva Amazónica, Amazonía or usually Amazonia; French: Forêt amazonienne; Dutch: Amazoneregenwoud), also known in English as Amazonia or the Amazon Jungle, is a moist broadleaf forest that covers most of the Amazon basin of South America. This basin encompasses 7,000,000 square kilometres (2,700,000 sq mi), of which 5,500,000 square kilometres (2,100,000 sq mi) are covered by the rainforest. This region includes territory belonging to nine nations. The majority of the forest is contained within Brazil, with 60% of the rainforest, followed by Peru with 13%, Colombia with 10%, and with minor amounts in Venezuela, Ecuador, Bolivia, Guyana, Suriname and French Guiana. States or departments in four nations contain 'Amazonas' in their names. The Amazon represents over half of the planet's remaining rainforests,[1] and comprises the largest and most biodiverse tract of tropical rainforest in the world, with an estimated 390 billion individual trees divided into 16,000 species.[2]",
	}
	db.Create(&problem)

	//Solution Creation
	solution := Solution{
		ProblemID:      problem.ID,
		OriginalPoster: jackDaniels,
		Title:          "The answer here is something very intense",
		Rank:           10,
	}
	glog.Info(solution.OriginalPoster)
	db.Create(&solution)

	solution2 := Solution{
		ProblemID:      problem.ID,
		OriginalPoster: shayan,
		Title:          "What we can do is enable users to try a solution",
		Rank:           10,
	}
	db.Create(&solution2)
	// solution2 := Solution{}
	// db.Create(&solution2)

}
