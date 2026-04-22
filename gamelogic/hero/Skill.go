package hero

import (
	"wvtrserv/data"
	"wvtrserv/gamelogic"
)

func GetRandomUniqueSkill() *data.Skill {
	skills := data.GetUniqueSkills()
	weights := make([]float64, 0)
	for _, c := range skills {
		weights = append(weights, c.Weight)
	}
	weights = gamelogic.NormalizeArray(weights)
	idx := gamelogic.RollInArrayWithRate(gamelogic.NaturalRoll(0, 1), weights)
	return skills[idx]
}

type SkillInterface interface {
	Use(from *Hero, to *Hero) *string
}

type UniqueSkill data.Skill

func (s UniqueSkill) Use(from *data.Hero, to *Hero) string {
	return "Used Unique"
}

type Lucky struct {
	UniqueSkill
}

func (s Lucky) CanBeUse(x float64) bool {
	roll := gamelogic.NaturalRoll(0, 1)
	target := (x * x) / (200 * (x * x))
	return gamelogic.RollCheck(roll, target)
}

func (s Lucky) Use(from *Hero, to *Hero) string {
	if s.CanBeUse(float64(from.Attributes.GetLuck())) {
		return "Activate"
	}
	return ""
}
