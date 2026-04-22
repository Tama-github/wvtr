package hero

import (
	"wvtrserv/data"
	"wvtrserv/gamelogic"
)

func GetRandomHeroClass() *data.HeroClass {
	classes := data.GetHeroClasses()
	weights := make([]float64, 0)
	for _, c := range classes {
		weights = append(weights, c.Weight)
	}
	weights = gamelogic.NormalizeArray(weights)
	idx := gamelogic.RollInArrayWithRate(gamelogic.NaturalRoll(0, 1), weights)
	return classes[idx]
}
