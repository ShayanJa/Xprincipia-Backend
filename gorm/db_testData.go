package gorm

import (
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
		Username:       "XprincipiaTeam",
		PhoneNumber:    "9528075184",
		HashedPassword: hashedShyshawnPassword,
	}
	db.Create(&shayan)

	// jackDaniels := User{
	// 	FirstName:   "Jack",
	// 	LastName:    "Daniels",
	// 	Email:       "dev@xPrincipia.com",
	// 	Address:     "5515 Malibu Dr",
	// 	Username:    "jackDaniels",
	// 	PhoneNumber: "9528015184",
	// }
	// db.Create(&jackDaniels)

	//Problem Creation
	// 	problem := Problem{
	// 		ParentID:       0,
	// 		OriginalPoster: shayan,
	// 		Title:          `The Applications of Time Crystals `,
	// 		Field:          "Physics",
	// 		Summary:        "What are the  applications of time crystals, ranked by value to humanity?  Detail each as precisely as possible.",
	// 		Description: `Concept

	// In March 2017, two research teams reported the creation of time crystals. This is a type of matter that exhibits a fundamental oscillation over time. [1] These systems are in non-equilibrium with their environment and exhibit discrete time-translation symmetry breaking. This is similar to how atomic crystals break continuous space-translation symmetry by only allowing atoms a discrete choice of where to exist, as opposed to continuous. In time crystals, the  system allows only a discrete choice of when to exist. In time, this means the system repeats some internal state with constant separations of time, as a regular crystal exhibits a periodic pattern through space. The crystal also exhibits a resiliency to change from these periodic states, much as regular crystals do in their atomic lattices. [4]

	// Procedure

	// The concept was initially introduced in 2012. By 2016, a procedure was outlined for creating these states. Using an open system, as opposed to a closed system where no energy is allowed in or out, a laser is used as an external input of energy to force oscillating states. This creates the condition of non-equilibrium. The states used are the alignment of electron spin in a chain of ions. The spins line up with each other due to interacting magnetic fields, either in direct alignment or opposite alignment. These are both lower energy states than random alignments. The well-ordered and periodic electromagnetic waves from the laser cause the spins of the ions to oscillate. The period of the spin oscillations is an integer multiple of the driving period from the laser. [1]

	// However, when the laser is stopped or changed, the spin oscillations continue. The system continues to be in a state of non-equilibrium. This is due to the fact the oscillations themselves are the ground states of the system. The oscillations are sustained internally and will resist change due to an outside perturbation.

	// The time crystals exhibit topological order, an emergent phenomenon where correlations in the particles are encoded in the entire wave-function of the system, which allows the quantum states to stabilize against decoherence effects. This can greatly enhance the efficiency of some information theory and quantum thermodynamic tasks. [3] They are characterized by this many-body localization , as well as by a subharmonic temporal response. [5]

	// After their creation, the time crystals are stable in their ground state without consuming or producing energy. Importantly, the ground states would be time-dependent and non-random. This is a distinction between normal matter in thermal equilibrium, where there is only random internal motion.

	// Creation

	// Two teams used alternative methods for the 2017 creation of the time crystals. Each showed that it is impossible for the crystals to be in equilibrium with their environment over time. This ensures they are in the state of non-equilibrium, although still in their ground states. One team used a chain of 10 ytterbium ions driven by a pair of lasers in a similar method as proposed in 2016. The chain was trapped in a quadrupole ion trap, a type of ion trap that is also used in trapped ion quantum computers. [7]

	// Subharmonic oscillation of the drive was observed. This experiment showed the rigidity of the time crystal, where the oscillation frequency remained unchanged even upon purposeful perturbation of the state, up to a certain degree. The relationship between interaction strength between the ions versus the imperfections in the spin-flip driving signal can be seen in the phase diagram shown in Reference 2.

	// When the variations in the driving signal become too strong and the interaction strength too weak, the time crystal “melts” into regular time-symmetric matter. Here the symmetry remains unbroken and the ion chain follows the driving signal without independent rhythm of its own. If the connections between the ions are too strong, random thermal effects take over and the symmetry again remains unbroken. [1]

	// The other team used an alternative method, using microwaves to generate oscillations in the spins of nitrogen impurities inside a diamond. A similar phase diagram was also observed. This system developed a spin period three times that of the microwave period, while the ytterbium ion oscillations were twice the period of the laser.  [1]

	// Applications

	// Time crystals are a new discovery, and thus their applications are not fully known. Initial ideas include use in quantum computation and furthering the theoretical framework of time.

	// In quantum computing, the most popular approach to building a quantum computing memory element is to use electron spins, where the up-down direction of the spin represents the classical 1 and 0. Maintaining these states is difficult, as even small quantities of random motion from heat ruin the prepared array of entangled spin alignments used in the calculation. The resiliency of the spin-flip cycle could contribute to building stable quantum memory. [1]

	// Until this discovery, time was an exception as a major symmetry that hadn’t yet been broken. As the symmetry of time is directly related to the conservation of energy, fields or particles in the presence of a time crystal background will appear to violate the conservation of energy. However, this is in reality a transfer to the vacuum field. The predicted properties of the time crystals introduce the concept of quasienergy. [3]
	// `,
	// 		Requirements: `1. Describe each possible application of time crystals.
	// 2.Order the applications by their importance to the future of humanity. Explanations can be given for the choice of ranking.
	// 3.Assign an estimation of date, with error bars, for when each application could become physically realized.`,
	// 		References: `PBS SPACETIME TIME CRYSTALS: https://www.youtube.com/watch?v=5l1KxgHH2Ek
	// PHASE DIAGRAM: https://www.eurekalert.org/multimedia/pub/131712.php?from=348795
	// WIKIPEDIA TIME CRYSTAL: https://en.wikipedia.org/wiki/Time_crystal
	// 2013 WIRED ARTICLE: https://www.wired.com/2013/04/time-crystals
	// 2016 PROCEDURE PAPER: https://physics.aps.org/featured-article-pdf/10.1103/PhysRevLett.118.030401
	// 2017 NITROGEN TIME CRYSTAL PAPER: https://arxiv.org/pdf/1610.08057v1.pdf
	// 2017 YTTERBIUM TIME CRYSTAL PAPER: https://arxiv.org/pdf/1609.08684.pdf
	// `,
	// 	}
	// 	db.Create(&problem)

	// 	problem2 := Problem{
	// 		ParentID:       0,
	// 		OriginalPoster: shayan,
	// 		Field:          "Computer Science",
	// 		Title:          "What is the probability of life arising independently on each of the seven planets in the TRAPPIST-1 solar system?",
	// 		Summary:        "Despite all of the security it offers, quantum cryptology also has a few fundamental flaws. Chief among these flaws is the length under which the system will work: It?s too short.",
	// 		Description: `TRAPPIST-1 Solar System

	// In 2017, seven temperate terrestrial planets were discovered orbiting the ultracool red dwarf star TRAPPIST-1. These planets are labeled TRAPPIST-1b, c, d, e, f, g, and h. The star is located 39.5 light years from the Sun in the Aquarius Constellation. Planets b, c, e, f and g are similar in size to Earth while planets d and h are intermediate in size between Earth and Mars. Planets e, f, and g orbit in the circumstellar habitable zone. Given the proper atmospheric conditions, these three are more likely to support liquid water. Despite this, it is considered possible for each of the seven planets to possess liquid water somewhere on their surface. [6]

	// Each planet in the system orbits much closer than Mercury orbits the sun. Each planet except b orbits farther than the Galilean satellites around Jupiter but closer than most of Jupiter’s other moons.  All seven planets are likely to be tidally locked, where one side of each planet faces the star at all times. If so, the most likely places for life may be near the terminator lines, the mid-twilight regions between the two sides. [6] The atmospheres of the planets are unknown, but information can be inferred.

	// The combined transmission spectrum of planets b and c from the Hubble Space Telescope rule out cloud-free hydrogen-dominated atmospheres for both planets. They are then unlikely to harbor an extended gas envelope unless it is cloudy out to high altitudes. The featureless spectrum still allows for a range of other atmospheric structures, from a cloud-free water vapor to a possibly Venus-like atmosphere.[6] TRAPPIST-1h is considered to have an equilibrium temperature of 169 K, placing it at the snow line, where volatile compounds such as water, ammonia, methane, carbon dioxide, and carbon monoxide can condense into solid ice grains. [4] Even here, liquid water could still exist either under an H2-rich atmosphere or a 2.7 km layer of ice. [2]

	// Recent photometric data collected by the Kepler-2 mission suggests that TRAPPIST-1 is prone to magnetic instability and large solar flares.  Over an 80-day observation period, 42 strong flaring events were identified with five multiple peaks and an average frequency of every 28 hours. These high energy flares are possibly as strong as  the Carryington Event in 1859. The strongest of the events could possibly damage the atmospheres to the degree requiring approximately 30,000 years to return to normal. This likely alters the atmospheres of the planets, lowering the likelihood of abiogenesis in the TRAPPIST-1  planetary system. [1]

	// The exact values of the masses and radii of the solar system members are thought to be: [6]

	// TRAPPIST-1 Star: 0.08±0.009 Solar Masses, 0.114±0.006 Solar Radii
	// b: 0.85±0.72 Earth Masses, 1.086±0.035 Earth Radii
	// c: 1.38±0.61 Earth Masses, 1.056±0.035 Earth Radii
	// d: 0.41±0.27 Earth Masses, 0.772±0.030 Earth Radii
	// e: 0.62±0.58 Earth Masses, 0.918±0.039 Earth Radii
	// f: 0.68±0.18 Earth Masses, 1.045±0.038 Earth Radii
	// g: 1.34±0.88 Earth Masses, 1.127±0.041 Earth Radii
	// h: Unknown, likely < 1 Earth Masses, 0.715+0.047,-0.043 Earth Radii

	// Abiogenesis

	// Abiogenesis is the natural process by which life arises from non-living matter, such as organic compounds. On Earth, the earliest undisputed evidence of life dates back to at least 3.5 billion years ago. However, recent discoveries of microfossils in Quebec have been dated with rocks between 3.77 and 4.28 billion years ago. With the formation of the ocean 4.4 billion years ago, this indicates abiogenesis could be possible relatively quickly after oceans or liquid water arise on a planet. [3] TRAPPIST-1 is estimated to be between 3-8 billion years old. The likelihood that life has formed increases with the age of the planets and their possible oceans or liquid water. [2]

	// Life on Earth is dependent on a specialized chemistry of carbon and water. Due to the low luminosity of TRAPPIST-1, much of the electromagnetic spectrum falls in the infrared part of the electromagnetic spectrum. Because water strongly absorbs red and infrared light, less energy would be available for aquatic life. [5]

	// As far as life found on Earth can indicate, liquid water holds the key towards the creation of life. The first step towards determining the chance of life around the nearby TRAPPIST-1 star is likely directly linked with the odds of liquid water on each newly discovered planet.
	// `,
	// 		Requirements: `1.Express with percentages the likelihood of abiogenesis on each of the seven planets.
	// 2.Associate error bars with each estimation, due to the unknown nature of the planets.
	// `,
	// 		References: `1. 2017 FREQUENT FLARING PAPER: https://arxiv.org/pdf/1703.10130.pdf
	// 2. 2017 SNOW LINE PAPER: https://arxiv.org/pdf/1703.04166.pdf
	// 3. WIKIPEDIA ABIOGENESIS: https://en.wikipedia.org/wiki/Abiogenesis
	// 4. WIKIPEDIA FROST LINE: https://en.wikipedia.org/wiki/Frost_line_(astrophysics)
	// 5. WIKIPEDIA HABITABLIITY OF RED DWARF SYSTEMS: https://en.wikipedia.org/wiki/Habitability_of_red_dwarf_systems
	// 6. WIKIPEDIA TRAPPIST-1: https://en.wikipedia.org/wiki/TRAPPIST-1
	// `,
	// 	}
	// 	db.Create(&problem2)

	// problem3 := Problem{
	// 	OriginalPoster: shayan,
	// 	ParentID:       0,
	// 	Field:          "Physics",
	// 	Title:          "Naked singularity",
	// 	Summary:        "In general relativity, a naked singularity is a gravitational singularity without an event horizon. In a black hole, the singularity is completely enclosed by a boundary known as the event horizon, inside which the gravitational force of the singularity is so strong that light cannot escape. Hence, objects inside the event horizon?including the singularity itself?cannot be directly observed. A naked singularity, by contrast, is observable from the outside.",
	// 	Description:    "The theoretical existence of naked singularities is important because their existence would mean that it would be possible to observe the collapse of an object to infinite density. It would also cause foundational problems for general relativity, because general relativity cannot make predictions about the future evolution of space-time near a singularity. In generic black holes, this is not a problem, as an outside viewer cannot observe the space-time within the event horizon.ome research has suggested that if loop quantum gravity is correct, then naked singularities could exist in nature,[1][2][3] implying that the cosmic censorship hypothesis does not hold. Numerical calculations[4] and some other arguments[5] have also hinted at this possibility.At LIGO, first observation of gravitational waves were detected after the collision two black holes, known as event GW150914. This event did not produce a naked singularity based on observation.[6]",
	// }
	// db.Create(&problem3)

	// problem4 := Problem{
	// 	OriginalPoster: shayan,
	// 	ParentID:       0,
	// 	Field:          "Investment",
	// 	Title:          "High-yield investment program",
	// 	Summary:        "A high-yield investment program (HYIP) is a type of Ponzi scheme, an investment scam that promises unsustainably high return on investment by paying previous investors with the money invested by new investors.",
	// 	Description:    "Operators generally set up a website offering an 'investment program' which promises very high returns, such as 1% per day (3778% APY when returns are compounded every day), disclosing little or no detail about the underlying management, location, or other aspects of how money is to be invested. The U.S. Securities and Exchange Commission (SEC) has said that 'these fraudulent schemes involve the purported issuance, trading, or use of so-called 'prime' bank, 'prime' European bank or 'prime' world bank financial instruments, or other 'high yield investment programs.' (HYIP's) The fraud artists … seek to mislead investors by suggesting that well regarded and financially sound institutions participate in these bogus programs.'[1] In 2010, the Financial Industry Regulatory Authority (FINRA) warned that '[t]he con artists behind HYIPs are experts at using social media — including YouTube, Twitter and Facebook — to lure investors and create the illusion of social consensus that these investments are legitimate'",
	// }
	// db.Create(&problem4)

	// problem5 := Problem{
	// 	OriginalPoster: jackDaniels,
	// 	Field:          "Consciousness",
	// 	ParentID:       0,
	// 	Title:          "A model of consciousness for the human brain and phenomenal mind",
	// 	Summary:        "A model of consciousness is needed that is a theoretical description relating brain properties of consciousness to phenomenal properties of consciousness. Useful models can be either mathematical/logical or verbal/conceptual.",
	// 	Description:    "Models of consciousness should be distinguished from so-called neural correlates of consciousness (Crick & Koch 1990). While the identification of correlations between aspects of brain activity and aspects of consciousness may constrain the specification of neurobiologically plausible models, such correlations do not by themselves provide explanatory links between neural activity and consciousness. Models should also be distinguished from theories that do not propose any mechanistic implementation (e.g., Rosenthal’s ‘higher-order thought’ theories, Rosenthal 2005). Consciousness models are valuable precisely to the extent that they propose such explanatory links (Seth, 2009). This article summarizes models that include computational, informational, or neurodynamic elements that propose explanatory links between neural properties and phenomenal properties.",
	// 	Requirements:   "Must include computational, informational, or neurodynamic elements that propose explanatory links between neural properties and phenomenal properties",
	// 	References:     "Seth, A. K. & Baars, B. J. 2005 Neural Darwinism and consciousness. Consciousness and Cognition 14, 140-168. Seth, A. K., Baars, B. J. & Edelman, D. B. 2005 Criteria for consciousness in humans and other mammals. Consciousness and Cognition 14, 119-139. Seth, Anil. 'Models of Consciousness.' Scholarpedia. N.p., n.d. Web. 05 Apr. 2017. ",
	// }
	// db.Create(&problem5)

	// //Solution Creation
	// solution := Solution{
	// 	ProblemID:      problem.ID,
	// 	OriginalPoster: jackDaniels,
	// 	Title:          "The answer here is something very intense",
	// 	Rank:           10,
	// }
	// glog.Info(solution.OriginalPoster)
	// db.Create(&solution)

	// solution2 := Solution{
	// 	ProblemID:      problem.ID,
	// 	OriginalPoster: shayan,
	// 	Title:          "What we can do is enable users to try a solution",
	// 	Rank:           10,
	// }
	// db.Create(&solution2)

	// solution3 := Solution{
	// 	ProblemID:      problem.ID,
	// 	OriginalPoster: shayan,
	// 	Title:          "Quantum Computation in the Microtubule Tryptophans",
	// 	Summary:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras ut dolor ante. Duis id pretium metus. Nunc suscipit in ipsum eu condimentum. Fusce nec egestas sapien, id mattis nisl. Curabitur venenatis dui lorem, a rutrum nibh sollicitudin id.",
	// 	Rank:           10,
	// }
	// db.Create(&solution3)

	// solution4 := Solution{
	// 	ProblemID:      problem5.ID,
	// 	OriginalPoster: shayan,
	// 	Title:          "Global Workspace",
	// 	Summary:        "Conscious cognitive content is globally available for diverse cognitive processes. Explains the association of consciousness with integrative cognitive processes. May naturally account for the serial nature of conscious experience.",
	// 	Description:    "GW theory was originally described in terms of a ‘blackboard’ architecture in which separate, quasi-independent processing modules interface with a centralized, globally available resource (Baars 1988). This cognitive level of description is preserved in the computational models of Franklin and Graesser (1999), who proposed a model consisting of a population of interacting ‘software agents’, and Shanahan (2005), whose model incorporates aspects of internal simulation supporting executive control and more recently spiking neurons (Shanahan, 2008).Dehaene, Changeux and colleagues have proposed a neuronal implementation of a global workspace architecture, the so-called ‘’neuronal global workspace’’ (see Figure and (Dehaene et al. 2003)). In this model, sensory stimuli mobilize excitatory neurons with long-range cortico-cortical axons, leading to the genesis of a global activity pattern among workspace neurons. Any such global pattern can inhibit alternative activity patterns among workspace neurons, thus preventing the conscious processing of alternative stimuli (for example, during the so-called attentional blink). The global neuronal workspace model predicts that conscious presence is a nonlinear function of stimulus salience; i.e., a gradual increase in stimulus visibility should be accompanied by a sudden transition of the neuronal workspace into a corresponding activity pattern (Dehaene et al. 2003).Wallace has advocated a network-theoretic modelling perspective on global workspace theory (Wallace 2005). In this view, transient links among specialized processing modules comprise dynamically formed networks. The ignition of a global workspace corresponds to the formation of a ‘giant component’ whereby previously disconnected sub-networks coalesce into a single network encompassing the majority of modules. The emergence of giant components in dynamic networks can be considered as a phase transition.",
	// 	Evidence:       "May naturally account for the serial nature of conscious experience. ",
	// 	Experiments:    "No Future Experiments yet",
	// 	References:     "Baars, B. J. 1988 A cognitive theory of consciousness. New York, NY: Cambridge University Press. ",
	// 	Rank:           10,
	// }
	// db.Create(&solution4)

	// //Create Suggestions

	// suggestion1 := Suggestion{
	// 	Username:    "Ben Francis",
	// 	Description: "You could add the fact that entropy is actually increased through abiogenesis through some sort of teleology. Also it could be noted the history of complexity increase in the universe is the dependent variable of a parabolic curve where entropy always increases as the independent variable.",
	// }
	// db.Create(&suggestion1)

	// //Create questions
	// question1 := Question{
	// 	Type:        1,
	// 	TypeID:      1,
	// 	Description: "What does section 4 mean?",
	// }
	// db.Create(&question1)

	// question2 := Question{
	// 	Type:        1,
	// 	TypeID:      2,
	// 	Description: "Why does section 4 mean?",
	// }
	// db.Create(&question2)

	// question3 := Question{
	// 	Type:        1,
	// 	TypeID:      1,
	// 	Description: "Because does section 4 mean?",
	// }
	// db.Create(&question3)

}
