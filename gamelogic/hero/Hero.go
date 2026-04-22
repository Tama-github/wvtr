package hero

import (
	"strings"
	"wvtrserv/data"
	"wvtrserv/gamelogic"
	"wvtrserv/nanapi/client"
)

type Hero data.Hero

func GenerateGrowthRateFromRank(rank string) []float64 {
	res := make([]float64, 5)
	min := 0.0
	max := 0.0
	switch rank {
	case "S":
		min = 0.1
		max = 0.3
	case "A":
		min = 0.05
		max = 0.2
	case "B":
		min = 0.05
		max = 0.15
	case "C":
		min = 0
		max = 0.1
	case "D":
		min = 0
		max = 0.05
	}
	res[data.HPgtID] = gamelogic.NaturalRoll(min, max)
	res[data.SgtID] = gamelogic.NaturalRoll(min, max)
	res[data.IgtID] = gamelogic.NaturalRoll(min, max)
	res[data.DgtID] = gamelogic.NaturalRoll(min, max)
	res[data.LgtID] = gamelogic.NaturalRoll(min, max)
	return res
}

func CreateNewHeroFromDBWaifuInfos(wc *client.JoinWC) *data.Hero {
	// select class
	class := GetRandomHeroClass()
	attributes := data.NewHeroAttribute(class, GenerateGrowthRateFromRank(wc.Rank))
	uniqueSkill := GetRandomUniqueSkill()
	resHero := &data.Hero{
		ImageUrl:       wc.ImageLarge,
		Name:           wc.NameUserPreferred,
		Class:          class,
		Rank:           wc.Rank,
		Attributes:     attributes,
		UniqueSkill:    uniqueSkill,
		WaifuID:        wc.ID,
		AnilistCharaID: uint(wc.IdAl),
	}

	resHero.LevelUp()
	resHero.Attributes.XPToLvlUp = resHero.Attributes.LevelThreshold()
	resHero.Attributes.CurrentHP = resHero.Attributes.MaxHP
	return resHero
}

func Roll(h *data.Hero, min float64, max float64) float64 {
	uskill := UniqueSkill(*h.UniqueSkill)
	resRoll := gamelogic.NaturalRoll(min, max)

	// Lucky
	if h.HasUniqueSkill(data.Lucky) && strings.Compare(uskill.Use(h, nil), "Activate") == 0 {
		skillRoll := gamelogic.NaturalRoll(min, max)
		if resRoll < skillRoll {
			resRoll = skillRoll
		}
	}
	return resRoll
}

func (h Hero) Attack(target Hero) string {
	logResult := ""
	// get attack value

	// check miss

	// check critic

	// damage target

	// check if target dodged

	// check if target blocked

	// add leach if there is

	// check if target is dead

	return logResult
}

func (h *Hero) takeFlatDamage(dmg float64) {
	h.Attributes.CurrentHP -= dmg
}

func (h *Hero) TakeDamage(dmg float64, takeFrom Hero) (float64, float64) {
	// check dodge

	// check blocked

	// check resistances

	// get total tamage taken

	h.takeFlatDamage(dmg)

	// send reflected damage if there are any

	// check if dead

	return 0, dmg
}
