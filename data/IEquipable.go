package data

// type IWeapon interface {
// 	IEquipable
// 	ToWeapon() *Weapon
// }

// func (e *Weapon) ToWeapon() *Weapon {
// 	return e
// }

// func (e *Weapon) Equip(h *Hero) IWeapon {
// 	var res IWeapon
// 	res = h.Equipment.Weapon
// 	h.Equipment.Weapon = e
// 	return res
// }

type IEquipable interface {
	Equip(*Hero) IEquipable
	RollBaseStats()
}

func (e *Weapon) Equip(h *Hero) IEquipable {
	var res IEquipable
	res = h.Equipment.Weapon
	h.Equipment.Weapon = e
	h.Equipment.WeaponID = e.ID
	e.InventoryID = 0
	return res
}

func (e *Armor) Equip(h *Hero) IEquipable {
	var res IEquipable
	res = h.Equipment.Armor
	h.Equipment.Armor = e
	h.Equipment.ArmorID = e.ID
	e.InventoryID = 0
	return res
}

func (e *Omamori) Equip(h *Hero) IEquipable {
	var res IEquipable
	res = h.Equipment.Omamori
	h.Equipment.Omamori = e
	h.Equipment.OmamoriID = e.ID
	e.InventoryID = 0
	return res
}

func (e *Weapon) RollBaseStats() {
	e.BaseAttackSpeed.RollValue()
	e.BaseCritRate.RollValue()
	e.BaseAttackSpeed.RollValue()
}

func (e *Armor) RollBaseStats() {
	e.BlockScore.RollValue()
}

func (e *Omamori) RollBaseStats() {
}
