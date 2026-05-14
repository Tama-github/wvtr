package databasecontroller

import "wvtrserv/data"

func GetInventoryByID(id uint) *data.Inventory {
	var res *data.Inventory = &data.Inventory{}
	db.Preload("Weapons").
		Preload("Armors").
		Preload("Omamoris").
		Preload("Currencies").
		Find(&res, id)

	for i := range res.Weapons {
		res.Weapons[i] = GetWeaponByID(res.Weapons[i].ID)
	}

	for i := range res.Armors {
		res.Armors[i] = GetArmorByID(res.Armors[i].ID)
	}

	for i := range res.Omamoris {
		res.Omamoris[i] = GetOmamoriByID(res.Omamoris[i].ID)
	}

	for i := range res.Currencies {
		res.Currencies[i] = GetCurrencyOwnedByID(res.Currencies[i].ID)
	}

	return res
}

func SaveInventory(inv *data.Inventory) {
	for _, w := range inv.Weapons {
		w.InventoryID = inv.ID
		SaveWeapon(w)
	}
	for _, a := range inv.Armors {
		a.InventoryID = inv.ID
		SaveArmor(a)
	}
	for _, o := range inv.Omamoris {
		o.InventoryID = inv.ID
		SaveOmamori(o)
	}
	for _, c := range inv.Currencies {
		SaveCurrencyOwned(c)
	}
}

func CreateInventory(inv *data.Inventory) {
	for _, w := range inv.Weapons {
		w.InventoryID = inv.ID
		CreateWeapon(w)
	}
	for _, a := range inv.Armors {
		a.InventoryID = inv.ID
		CreateArmor(a)
	}
	for _, o := range inv.Omamoris {
		o.InventoryID = inv.ID
		CreateOmamori(o)
	}
	db.Create(inv)
}
