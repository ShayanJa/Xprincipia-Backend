package gorm

import (
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"work/xprincipia/backend/util"
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
		ParentID:       0,
		OriginalPoster: shayan,
		Title:          "What causes the emergent phenomena of consciousness?",
		Field:          "Consciousness",
		Summary:        "The name Amazon is said to arise from a war Francisco de Orellana fought with the Tapuyas and other tribes. The women of the tribe fought alongside the men, as was their custom.[3] Orellana derived the name Amazonas from the Amazons of Greek mythology, described by Herodotus and Diodorus.",
		Description:    "The Amazon rainforest (Portuguese: Floresta Amazônica or Amazônia; Spanish: Selva Amazónica, Amazonía or usually Amazonia; French: Forêt amazonienne; Dutch: Amazoneregenwoud), also known in English as Amazonia or the Amazon Jungle, is a moist broadleaf forest that covers most of the Amazon basin of South America. This basin encompasses 7,000,000 square kilometres (2,700,000 sq mi), of which 5,500,000 square kilometres (2,100,000 sq mi) are covered by the rainforest. This region includes territory belonging to nine nations. The majority of the forest is contained within Brazil, with 60% of the rainforest, followed by Peru with 13%, Colombia with 10%, and with minor amounts in Venezuela, Ecuador, Bolivia, Guyana, Suriname and French Guiana. States or departments in four nations contain 'Amazonas' in their names. The Amazon represents over half of the planet's remaining rainforests,[1] and comprises the largest and most biodiverse tract of tropical rainforest in the world, with an estimated 390 billion individual trees divided into 16,000 species.[2]",
		Rank:           1,
	}
	db.Create(&problem)

	problem2 := Problem{
		ParentID:       0,
		OriginalPoster: shayan,
		Field:          "Computer Science",
		Title:          "How Quantum Cryptology Works",
		Summary:        "Despite all of the security it offers, quantum cryptology also has a few fundamental flaws. Chief among these flaws is the length under which the system will work: It?s too short.",
		Description:    "When we use the Internet, we're not always just clicking around and passively taking in information, such as reading news articles or blog posts -- a great deal of our time online involves sending others our own information. Ordering something over the Internet, whether it's a book, a CD or anything else from an online vendor, or signing up for an online account, requires entering in a good deal of sensitive personal information. A typical transaction might include not only our names, e-mail addresses and physical address and phone number, but also passwords and personal identification numbers (PINs).",
		Rank:           1,
	}
	db.Create(&problem2)

	problem3 := Problem{
		OriginalPoster: shayan,
		ParentID:       0,
		Field:          "Physics",
		Title:          "Naked singularity",
		Summary:        "In general relativity, a naked singularity is a gravitational singularity without an event horizon. In a black hole, the singularity is completely enclosed by a boundary known as the event horizon, inside which the gravitational force of the singularity is so strong that light cannot escape. Hence, objects inside the event horizon?including the singularity itself?cannot be directly observed. A naked singularity, by contrast, is observable from the outside.",
		Description:    "The theoretical existence of naked singularities is important because their existence would mean that it would be possible to observe the collapse of an object to infinite density. It would also cause foundational problems for general relativity, because general relativity cannot make predictions about the future evolution of space-time near a singularity. In generic black holes, this is not a problem, as an outside viewer cannot observe the space-time within the event horizon.ome research has suggested that if loop quantum gravity is correct, then naked singularities could exist in nature,[1][2][3] implying that the cosmic censorship hypothesis does not hold. Numerical calculations[4] and some other arguments[5] have also hinted at this possibility.At LIGO, first observation of gravitational waves were detected after the collision two black holes, known as event GW150914. This event did not produce a naked singularity based on observation.[6]",
		Rank:           1,
	}
	db.Create(&problem3)

	problem4 := Problem{
		OriginalPoster: jackDaniels,
		ParentID:       0,
		Field:          "Investment",
		Title:          "High-yield investment program",
		Summary:        "A high-yield investment program (HYIP) is a type of Ponzi scheme, an investment scam that promises unsustainably high return on investment by paying previous investors with the money invested by new investors.",
		Description:    "Operators generally set up a website offering an 'investment program' which promises very high returns, such as 1% per day (3778% APY when returns are compounded every day), disclosing little or no detail about the underlying management, location, or other aspects of how money is to be invested. The U.S. Securities and Exchange Commission (SEC) has said that 'these fraudulent schemes involve the purported issuance, trading, or use of so-called 'prime' bank, 'prime' European bank or 'prime' world bank financial instruments, or other 'high yield investment programs.' (HYIP's) The fraud artists … seek to mislead investors by suggesting that well regarded and financially sound institutions participate in these bogus programs.'[1] In 2010, the Financial Industry Regulatory Authority (FINRA) warned that '[t]he con artists behind HYIPs are experts at using social media — including YouTube, Twitter and Facebook — to lure investors and create the illusion of social consensus that these investments are legitimate'",
		Rank:           1,
	}
	db.Create(&problem4)

	problem5 := Problem{
		OriginalPoster: jackDaniels,
		Field:          "Consciousness",
		ParentID:       0,
		Title:          "A model of consciousness for the human brain and phenomenal mind",
		Summary:        "A model of consciousness is needed that is a theoretical description relating brain properties of consciousness to phenomenal properties of consciousness. Useful models can be either mathematical/logical or verbal/conceptual.",
		Description:    "Models of consciousness should be distinguished from so-called neural correlates of consciousness (Crick & Koch 1990). While the identification of correlations between aspects of brain activity and aspects of consciousness may constrain the specification of neurobiologically plausible models, such correlations do not by themselves provide explanatory links between neural activity and consciousness. Models should also be distinguished from theories that do not propose any mechanistic implementation (e.g., Rosenthal’s ‘higher-order thought’ theories, Rosenthal 2005). Consciousness models are valuable precisely to the extent that they propose such explanatory links (Seth, 2009). This article summarizes models that include computational, informational, or neurodynamic elements that propose explanatory links between neural properties and phenomenal properties.",
		Requirements:   "Must include computational, informational, or neurodynamic elements that propose explanatory links between neural properties and phenomenal properties",
		References:     "Seth, A. K. & Baars, B. J. 2005 Neural Darwinism and consciousness. Consciousness and Cognition 14, 140-168. Seth, A. K., Baars, B. J. & Edelman, D. B. 2005 Criteria for consciousness in humans and other mammals. Consciousness and Cognition 14, 119-139. Seth, Anil. 'Models of Consciousness.' Scholarpedia. N.p., n.d. Web. 05 Apr. 2017. ",
		Rank:           1,
	}
	db.Create(&problem5)

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

	solution4 := Solution{
		ProblemID:      problem5.ID,
		OriginalPoster: shayan,
		Title:          "Global Workspace",
		Summary:        "Conscious cognitive content is globally available for diverse cognitive processes. Explains the association of consciousness with integrative cognitive processes. May naturally account for the serial nature of conscious experience.",
		Description:    "GW theory was originally described in terms of a ‘blackboard’ architecture in which separate, quasi-independent processing modules interface with a centralized, globally available resource (Baars 1988). This cognitive level of description is preserved in the computational models of Franklin and Graesser (1999), who proposed a model consisting of a population of interacting ‘software agents’, and Shanahan (2005), whose model incorporates aspects of internal simulation supporting executive control and more recently spiking neurons (Shanahan, 2008).Dehaene, Changeux and colleagues have proposed a neuronal implementation of a global workspace architecture, the so-called ‘’neuronal global workspace’’ (see Figure and (Dehaene et al. 2003)). In this model, sensory stimuli mobilize excitatory neurons with long-range cortico-cortical axons, leading to the genesis of a global activity pattern among workspace neurons. Any such global pattern can inhibit alternative activity patterns among workspace neurons, thus preventing the conscious processing of alternative stimuli (for example, during the so-called attentional blink). The global neuronal workspace model predicts that conscious presence is a nonlinear function of stimulus salience; i.e., a gradual increase in stimulus visibility should be accompanied by a sudden transition of the neuronal workspace into a corresponding activity pattern (Dehaene et al. 2003).Wallace has advocated a network-theoretic modelling perspective on global workspace theory (Wallace 2005). In this view, transient links among specialized processing modules comprise dynamically formed networks. The ignition of a global workspace corresponds to the formation of a ‘giant component’ whereby previously disconnected sub-networks coalesce into a single network encompassing the majority of modules. The emergence of giant components in dynamic networks can be considered as a phase transition.",
		Evidence:       "May naturally account for the serial nature of conscious experience. ",
		Experiments:    "No Future Experiments yet",
		References:     "Baars, B. J. 1988 A cognitive theory of consciousness. New York, NY: Cambridge University Press. ",
		Rank:           10,
	}
	db.Create(&solution4)

	//Create Suggestions
	suggestion1 := Suggestion{
		Username:    "Ben Francis",
		Description: "You could add the fact that entropy is actually increased through abiogenesis through some sort of teleology. Also it could be noted the history of complexity increase in the universe is the dependent variable of a parabolic curve where entropy always increases as the independent variable.",
	}
	db.Create(&suggestion1)

	//Create questions
	question1 := Question{
		Type:        util.PROBLEM,
		TypeID:      2,
		Description: "How can we bind ourselves with computers?",
	}
	db.Create(&question1)

	question2 := Question{
		Username:    "Shyshawn",
		Type:        util.PROBLEM,
		TypeID:      2,
		Description: "Why does section 4 mean?",
	}
	db.Create(&question2)

	question3 := Question{
		Username:    "Shyshawn",
		Type:        util.PROBLEM,
		TypeID:      1,
		Description: "Because does section 4 mean?",
	}
	db.Create(&question3)

	//Create freeForms
	freeForm1 := FreeForm{
		Username:    "Shyshawn",
		Type:        util.PROBLEM,
		TypeID:      6,
		Description: "Because does section 4 mean?",
	}
	db.Create(&freeForm1)

	//Create freeForms
	pro1 := Pro{
		Username:    "Shyshawn",
		Type:        util.SOLUTION,
		TypeID:      6,
		Description: "Because does section 4 mean?",
	}
	db.Create(&pro1)

	//Create freeForms
	con1 := Con{
		Username:    "Shyshawn",
		Type:        util.SOLUTION,
		TypeID:      6,
		Description: "Because does section 4 mean?",
	}
	db.Create(&con1)

}
