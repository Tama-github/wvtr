package databasecontroller

import "wvtrserv/data"

func GetArmorByID(id uint) *data.Armor {
	var armor *data.Armor = &data.Armor{}
	db.Preload("BlockScore").
		Preload("EvadeScore").
		Preload("BaseResistancesRange").
		Preload("Affixes").
		Find(&armor, id)

	for i := range armor.Affixes {
		armor.Affixes[i] = GetAffixByID(armor.Affixes[i].ID)
	}
	return armor
}

func SaveArmor(o *data.Armor) {
	SaveStatsRange(o.BlockScore)
	SaveStatsRange(o.EvadeScore)
	SaveDamage(o.BaseResistancesRange)
	for _, a := range o.Affixes {
		SaveAffixe(a)
	}
	o.BlockScoreID = o.BlockScore.ID
	o.EvadeScoreID = o.EvadeScore.ID
	db.Save(o)
}

func CreateArmor(o *data.Armor) {
	CreateStatsRange(o.BlockScore)
	CreateStatsRange(o.EvadeScore)
	CreateDamage(o.BaseResistancesRange)
	for _, a := range o.Affixes {
		CreateAffixe(a)
	}
	db.Create(o)
}
