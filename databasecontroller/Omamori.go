package databasecontroller

import "wvtrserv/data"

func GetOmamoriByID(id uint) *data.Omamori {
	var oma *data.Omamori = &data.Omamori{}
	db.Preload("Affixes").
		Find(&oma, id)
	for i := range oma.Affixes {
		oma.Affixes[i] = GetAffixByID(oma.Affixes[i].ID)
	}
	return oma
}

func SaveOmamori(o *data.Omamori) {
	if o == nil {
		return
	}
	for _, a := range o.Affixes {
		SaveAffixe(a)
	}
	db.Save(o)
}

func CreateOmamori(o *data.Omamori) {
	for _, a := range o.Affixes {
		CreateAffixe(a)
	}
	db.Create(o)
}
