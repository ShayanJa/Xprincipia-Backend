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

	problem2 := Problem{
		OriginalPoster: shayan,
		Title:          "How Quantum Cryptology Works",
		Summary:        "Despite all of the security it offers, quantum cryptology also has a few fundamental flaws. Chief among these flaws is the length under which the system will work: It?s too short.",
		Description:    "When we use the Internet, we're not always just clicking around and passively taking in information, such as reading news articles or blog posts -- a great deal of our time online involves sending others our own information. Ordering something over the Internet, whether it's a book, a CD or anything else from an online vendor, or signing up for an online account, requires entering in a good deal of sensitive personal information. A typical transaction might include not only our names, e-mail addresses and physical address and phone number, but also passwords and personal identification numbers (PINs).",
	}
	db.Create(&problem2)

	problem3 := Problem{
		OriginalPoster: shayan,
		Title:          "Naked singularity",
		Summary:        "In general relativity, a naked singularity is a gravitational singularity without an event horizon. In a black hole, the singularity is completely enclosed by a boundary known as the event horizon, inside which the gravitational force of the singularity is so strong that light cannot escape. Hence, objects inside the event horizon?including the singularity itself?cannot be directly observed. A naked singularity, by contrast, is observable from the outside.",
		Description:    "The theoretical existence of naked singularities is important because their existence would mean that it would be possible to observe the collapse of an object to infinite density. It would also cause foundational problems for general relativity, because general relativity cannot make predictions about the future evolution of space-time near a singularity. In generic black holes, this is not a problem, as an outside viewer cannot observe the space-time within the event horizon.ome research has suggested that if loop quantum gravity is correct, then naked singularities could exist in nature,[1][2][3] implying that the cosmic censorship hypothesis does not hold. Numerical calculations[4] and some other arguments[5] have also hinted at this possibility.At LIGO, first observation of gravitational waves were detected after the collision two black holes, known as event GW150914. This event did not produce a naked singularity based on observation.[6]",
	}
	db.Create(&problem3)

	problem4 := Problem{
		OriginalPoster: jackDaniels,
		Title:          "High-yield investment program",
		Summary:        "A high-yield investment program (HYIP) is a type of Ponzi scheme, an investment scam that promises unsustainably high return on investment by paying previous investors with the money invested by new investors.",
		Description:    "Operators generally set up a website offering an 'investment program' which promises very high returns, such as 1% per day (3778% APY when returns are compounded every day), disclosing little or no detail about the underlying management, location, or other aspects of how money is to be invested. The U.S. Securities and Exchange Commission (SEC) has said that 'these fraudulent schemes involve the purported issuance, trading, or use of so-called 'prime' bank, 'prime' European bank or 'prime' world bank financial instruments, or other 'high yield investment programs.' (HYIP's) The fraud artists … seek to mislead investors by suggesting that well regarded and financially sound institutions participate in these bogus programs.'[1] In 2010, the Financial Industry Regulatory Authority (FINRA) warned that '[t]he con artists behind HYIPs are experts at using social media — including YouTube, Twitter and Facebook — to lure investors and create the illusion of social consensus that these investments are legitimate'",
	}
	db.Create(&problem4)

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

	solution3 := Solution{
		ProblemID:      problem.ID,
		OriginalPoster: shayan,
		Title:          "Quantum Computation in the Microtubule Tryptophans",
		Summary:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras ut dolor ante. Duis id pretium metus. Nunc suscipit in ipsum eu condimentum. Fusce nec egestas sapien, id mattis nisl. Curabitur venenatis dui lorem, a rutrum nibh sollicitudin id.",
		Rank:           10,
	}
	db.Create(&solution3)

	//Create Suggestions

	suggestion1 := Suggestion{
		Username:    "Ben Francis",
		Description: "You could add the fact that entropy is actually increased through abiogenesis through some sort of teleology. Also it could be noted the history of complexity increase in the universe is the dependent variable of a parabolic curve where entropy always increases as the independent variable.",
	}
	db.Create(&suggestion1)

	//Create questions
	question1 := Question{
		Type:        1,
		TypeID:      1,
		Description: "What does section 4 mean?",
	}
	db.Create(&question1)

	question2 := Question{
		Type:        1,
		TypeID:      2,
		Description: "Why does section 4 mean?",
	}
	db.Create(&question2)

	question3 := Question{
		Type:        1,
		TypeID:      1,
		Description: "Because does section 4 mean?",
	}
	db.Create(&question3)

}
