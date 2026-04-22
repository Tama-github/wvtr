package data

import "wvtrserv/gamelogic"

func NewHero() *Hero {
	return &Hero{}
}

func (h Hero) HasUniqueSkill(skillId SkillID) bool {
	return h.UniqueSkill.Identifier == skillId
}

func (h *Hero) GoToLevel(level int) {
	if level < h.Attributes.Level {
		return
	}

	for range h.Attributes.Level - level {
		h.GainXP(h.Attributes.XPToLvlUp - h.Attributes.CurrentXP)
	}
}

func (h *Hero) GainXP(amount float64) {
	// if h.HasUniqueSkill(FastLearner) {
	// 	amount = amount * fast_learner.xp_multiplier
	// }

	for amount > 0 {
		thresholdForCurrentLevel := h.Attributes.LevelThreshold()
		xpToGainForLevel := amount

		if amount >= thresholdForCurrentLevel {
			xpToGainForLevel = thresholdForCurrentLevel
			h.LevelUp()
		}

		h.Attributes.CurrentXP += xpToGainForLevel
		amount -= xpToGainForLevel
	}
}

func (h *Hero) IncreaseAttributeWithRate() {
	attrs := h.Attributes.GetAttributesArray()
	grs := h.Attributes.GetGRArray()
	for i := range len(grs) {
		toadd := float64(int(grs[i]))
		proba := attrs[i] - toadd
		if gamelogic.RollCheck(gamelogic.NaturalRoll(0, 1), proba) {
			toadd++
		}
		attrs[i] += toadd
	}
	h.Attributes.SetAttributesWithArray(attrs)
}

func (h *Hero) LevelUp() {
	h.Attributes.Level += 1

	h.IncreaseAttributeWithRate()
}
