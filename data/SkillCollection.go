package data

type SkillType int

const (
	Unique SkillType = iota
	Active
)

type SkillID int

const (
	Lucky SkillID = iota
	GoodRest
	SecondWind
	Prodigy
	Berserk
	Trickster
	FastLearner
	ElementalCursed
	PhysicalCursed
)

// should be inserted in db once
var skillCollection []*Skill = []*Skill{
	{
		Identifier:  Lucky,
		Name:        "Lucky",
		Type:        Unique,
		ImageURL:    "/imgs/skills/Skill_lucky.png",
		Description: "Can (lck) reroll any luck based action and pick the highest score.",
		Weight:      10,
	},
	{
		Identifier:  GoodRest,
		Name:        "Good Rest",
		Type:        Unique,
		ImageURL:    "/imgs/skills/Skill_good_rest.png",
		Description: "This Hero rest faster and better.",
		Weight:      20,
	},
	{
		Identifier:  SecondWind,
		Name:        "Second Wind",
		Type:        Unique,
		ImageURL:    "",
		Description: "Once per expedition this hero survive a fatal blow and gain back all their hp.",
		Weight:      10,
	},
	{
		Identifier:  Prodigy,
		Name:        "Prodigy",
		Type:        Unique,
		ImageURL:    "",
		Description: "Better chance (lck/int) of gaining attributes on level up",
		Weight:      20,
	},
	{
		Identifier:  Berserk,
		Name:        "Berserk",
		Type:        Unique,
		ImageURL:    "",
		Description: "Apply a damage multiplicator to pure physical damage. Applied last",
		Weight:      50,
	},
	{
		Identifier:  Trickster,
		Name:        "Trickster",
		Type:        Unique,
		ImageURL:    "/imgs/skills/Skill_trickster.png",
		Description: "Can (lck/dex) reduce the necessary time to execute any actions.",
		Weight:      20,
	},
	{
		Identifier:  FastLearner,
		Name:        "Fast Learner",
		Type:        Unique,
		ImageURL:    "",
		Description: "Reduction of a random (lck/dex) amount of XP to level up. Determined each time the Hero level up",
		Weight:      30,
	},
}

func GetSkills() []*Skill {
	return skillCollection
}

func GetUniqueSkills() []*Skill {
	res := make([]*Skill, 0)
	for _, s := range skillCollection {
		if s.Type == Unique {
			res = append(res, s)
		}
	}
	return res
}

func GetActiveSkills() []*Skill {
	res := make([]*Skill, 0)
	for _, s := range skillCollection {
		if s.Type == Active {
			res = append(res, s)
		}
	}
	return res
}
