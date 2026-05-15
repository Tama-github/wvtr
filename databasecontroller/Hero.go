package databasecontroller

import (
	"wvtrserv/data"
	"wvtrserv/logger"
)

func CreateHero(hero *data.Hero) error {
	logger.DumpLog.Print("CreateNewHero")
	tx := db.Create(hero)
	return tx.Error
}

func SaveHero(hero *data.Hero) {
	SaveHeroAttributes(hero.Attributes)
	SaveHeroEquipment(hero.Equipment)
	db.Save(hero)
}

func GetHeroes() []*data.Hero {
	res := []*data.Hero{}
	db.Find(&res)
	return res
}

func GetHeroByID(id uint) *data.Hero {
	var res *data.Hero
	db.Preload("Class").
		Preload("Attributes").
		Preload("Class").
		Preload("WeaponAttack").
		Preload("UniqueSkill").
		Preload("ActiveSkill").
		Preload("Equipment").
		Find(&res, id)

	if res.Class != nil {
		res.Class = GetHeroClassByID(res.Class.ID)
	}
	if res.Attributes != nil {
		res.Attributes = GetHeroAttributesByID(res.Attributes.ID)
	}
	if res.WeaponAttack != nil {
		res.WeaponAttack = GetSkillByID(res.WeaponAttack.ID)
	}
	if res.UniqueSkill != nil {
		res.UniqueSkill = GetSkillByID(res.UniqueSkill.ID)
	}
	if res.ActiveSkill != nil {
		res.ActiveSkill = GetSkillByID(res.ActiveSkill.ID)
	}
	if res.Equipment != nil {
		res.Equipment = GetHeroEquipmentByID(res.Equipment.ID)
	}
	return res
}
