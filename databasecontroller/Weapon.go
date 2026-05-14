package databasecontroller

import "wvtrserv/data"

func GetWeaponByID(id uint) *data.Weapon {
	var weap *data.Weapon = &data.Weapon{}
	db.Preload("BaseDamage").
		Preload("BaseCritRate").
		Preload("BaseAttackSpeed").
		Preload("Affixes").
		Find(&weap, id)
	for i := range weap.Affixes {
		weap.Affixes[i] = GetAffixByID(weap.Affixes[i].ID)
	}
	return weap
}

func SaveWeapon(o *data.Weapon) {
	SaveDamage(o.BaseDamage)
	SaveStatsRange(o.BaseAttackSpeed)
	SaveStatsRange(o.BaseCritRate)
	for _, a := range o.Affixes {
		SaveAffixe(a)
	}
	db.Save(o)
}

func CreateWeapon(o *data.Weapon) {
	CreateDamage(o.BaseDamage)
	CreateStatsRange(o.BaseAttackSpeed)
	CreateStatsRange(o.BaseCritRate)
	for _, a := range o.Affixes {
		CreateAffixe(a)
	}
	o.BaseAttackSpeedID = o.BaseAttackSpeed.ID
	o.BaseCritRateID = o.BaseCritRate.ID
	db.Save(o)
}
