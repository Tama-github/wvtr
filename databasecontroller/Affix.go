package databasecontroller

import "wvtrserv/data"

func GetAffixByID(id uint) *data.Affix {
	var aff *data.Affix = &data.Affix{}
	db.Preload("Ranges").Find(&aff, id)
	return aff
}

func SaveAffixe(a *data.Affix) {
	for _, r := range a.Ranges {
		db.Save(r)
	}
	db.Save(a)
}

func CreateAffixe(a *data.Affix) {
	db.Create(a)
}
