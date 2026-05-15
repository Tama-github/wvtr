package data

func (he *HeroEquipment) GetTotalValueOfAffixInEquipment(af AffixType) *Affix {
	res := &Affix{Type: af}
	if he.Weapon != nil {
		res.Add(he.Weapon.GetTotalValueOfAffixInEquipment(af))
	}
	if he.Armor != nil {
		res.Add(he.Armor.GetTotalValueOfAffixInEquipment(af))
	}
	if he.Omamori != nil {
		res.Add(he.Omamori.GetTotalValueOfAffixInEquipment(af))
	}

	// No affixes have been found
	if len(res.Ranges) == 0 {
		res.Ranges = []*StatsRange{
			{Value: 0},
		}
	}

	return res
}

func (he *HeroEquipment) GetEquipmentTotalDodgeChance() float64 {
	res := 0.0

	// TODO: Add affix
	if he.Armor != nil {
		res += he.Armor.EvadeScore.Value
	}

	return res
}

func (he *HeroEquipment) GetEquipmentTotalBlockChance() float64 {
	res := 0.0

	// TODO: Add affix
	if he.Armor != nil {
		res += he.Armor.BlockScore.Value
	}

	return res
}

func (he *HeroEquipment) GetEquipmenetTotalCritChance() float64 {
	res := 0.0

	// TODO: Add affix
	if he.Weapon != nil && he.Weapon.BaseCritRate != nil {
		res += he.Weapon.BaseCritRate.Value
	}

	return res
}
