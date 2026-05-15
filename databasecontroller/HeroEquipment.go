package databasecontroller

import "wvtrserv/data"

func GetHeroEquipmentByID(id uint) *data.HeroEquipment {
	var equi *data.HeroEquipment = &data.HeroEquipment{}

	db.Preload("Weapon").
		Preload("Armor").
		Preload("Omamori").
		Find(&equi, id)

	if equi.Weapon != nil {
		equi.Weapon = GetWeaponByID(equi.Weapon.ID)
	}

	if equi.Armor != nil {
		equi.Armor = GetArmorByID(equi.Armor.ID)
	}

	if equi.Omamori != nil {
		equi.Omamori = GetOmamoriByID(equi.Omamori.ID)
	}

	return equi
}

func SaveHeroEquipment(he *data.HeroEquipment) {
	SaveWeapon(he.Weapon)
	SaveArmor(he.Armor)
	SaveOmamori(he.Omamori)
	db.Save(he)
}
