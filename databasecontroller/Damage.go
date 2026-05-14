package databasecontroller

import "wvtrserv/data"

func GetDamageByID(id uint) *data.Damage {
	var dmg *data.Damage = &data.Damage{}
	db.Find(&dmg, id)
	return dmg
}

func SaveDamage(d *data.Damage) {
	db.Save(d)
}

func CreateDamage(d *data.Damage) {
	db.Create(d)
}
