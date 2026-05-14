package gamedata

import (
	"time"
	"wvtrserv/data"
)

// should be inserted in db once
var skillCollection []*data.Skill = []*data.Skill{
	{
		Identifier:  data.Lucky,
		Name:        "Lucky",
		Type:        data.Unique,
		Targeting:   data.Self,
		ImageURL:    "/imgs/skills/Skill_lucky.png",
		Description: "Can (lck) reroll any luck based action and pick the highest score.",
		Weight:      10,
	},
	{
		Identifier:  data.GoodRest,
		Name:        "Good Rest",
		Type:        data.Unique,
		Targeting:   data.Self,
		ImageURL:    "/imgs/skills/Skill_good_rest.png",
		Description: "This Hero rest faster and better.",
		Weight:      20,
	},
	{
		Identifier:  data.SecondWind,
		Name:        "Second Wind",
		Type:        data.Unique,
		Targeting:   data.Self,
		ImageURL:    "/imgs/skills/Skill_second_wind.png",
		Description: "Once per expedition this hero survive a fatal blow and gain back all their hp.",
		Weight:      10,
	},
	{
		Identifier:  data.Prodigy,
		Name:        "Prodigy",
		Type:        data.Unique,
		Targeting:   data.Self,
		ImageURL:    "/imgs/skills/Skill_prodigy.png",
		Description: "Better chance (lck/int) of gaining attributes on level up",
		Weight:      20,
	},
	{
		Identifier:  data.Berserk,
		Name:        "Berserk",
		Type:        data.Unique,
		Targeting:   data.Self,
		ImageURL:    "/imgs/skills/Skill_berserk.png",
		Description: "Apply a damage multiplicator to pure physical damage. Applied last",
		Weight:      50,
	},
	{
		Identifier:  data.Trickster,
		Name:        "Trickster",
		Type:        data.Unique,
		Targeting:   data.Self,
		ImageURL:    "/imgs/skills/Skill_trickster.png",
		Description: "Can (lck/dex) reduce the necessary time to execute any actions.",
		Weight:      20,
	},
	{
		Identifier:  data.FastLearner,
		Name:        "Fast Learner",
		Type:        data.Unique,
		Targeting:   data.Self,
		ImageURL:    "/imgs/skills/Skill_fast_learner.png",
		Description: "Gain 20% more xp",
		Weight:      30,
	},
	// Active
	{
		Identifier:           data.Spit,
		Name:                 "Charge",
		Type:                 data.Active,
		Targeting:            data.Enemy,
		ImageURL:             "",
		Description:          "",
		RecuperationDuration: 2 * time.Second,
	},
	{
		Identifier:           data.Charge,
		Name:                 "Charge",
		Type:                 data.Active,
		Targeting:            data.Enemy,
		ImageURL:             "",
		Description:          "",
		RecuperationDuration: 5 * time.Second,
	},
	{
		Identifier:           data.Tackle,
		Name:                 "Tackle",
		Type:                 data.Active,
		Targeting:            data.Enemy,
		ImageURL:             "",
		Description:          "",
		RecuperationDuration: 3 * time.Second,
	},
	{
		Identifier:           data.Sweep,
		Name:                 "Sweep",
		Type:                 data.Active,
		Targeting:            data.Enemy,
		ImageURL:             "",
		Description:          "",
		RecuperationDuration: 3 * time.Second,
	},
	{
		Identifier:           data.FireBolt,
		Name:                 "Fire bolt",
		Type:                 data.Active,
		Targeting:            data.Enemy,
		ImageURL:             "",
		Description:          "",
		RecuperationDuration: 5 * time.Second,
	},
	{
		Identifier:           data.IceBolt,
		Name:                 "Ice bolt",
		Type:                 data.Active,
		Targeting:            data.Enemy,
		ImageURL:             "",
		Description:          "",
		RecuperationDuration: 3 * time.Second,
	},
	{
		Identifier:           data.LightningBolt,
		Name:                 "Lightning bolt",
		Type:                 data.Active,
		Targeting:            data.Enemy,
		ImageURL:             "",
		Description:          "",
		RecuperationDuration: 3 * time.Second,
	},
}

var WeaponAttackSkill *data.Skill = &data.Skill{
	Identifier:           data.WeaponAttack,
	Name:                 "Attack",
	Type:                 data.Active,
	Targeting:            data.Enemy,
	RecuperationDuration: 1 * time.Second,
	ImageURL:             "",
	Description:          "Attack with whatever they have in hand.",
	Weight:               1,
}

func GetSkills() []*data.Skill {
	return skillCollection
}

func GetUniqueSkills() []*data.Skill {
	res := make([]*data.Skill, 0)
	for _, s := range skillCollection {
		if s.Type == data.Unique {
			res = append(res, s)
		}
	}
	return res
}

func GetActiveSkills() []*data.Skill {
	res := make([]*data.Skill, 0)
	for _, s := range skillCollection {
		if s.Type == data.Active {
			res = append(res, s)
		}
	}
	return res
}

func GetAttackSkill() *data.Skill {
	return WeaponAttackSkill
}

func GetRandomUniqueSkill(skills []*data.Skill) *data.Skill {
	weights := make([]float64, 0)
	for _, c := range skills {
		weights = append(weights, c.Weight)
	}
	weights = data.NormalizeArray(weights)
	idx := data.RollInArrayWithRate(data.NaturalRoll(0, 1), weights)
	return skills[idx]
}
