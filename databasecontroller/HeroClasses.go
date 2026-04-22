package databasecontroller

import "wvtrserv/data"

func CreateHeroClasse(hc *data.HeroClass) {
	db.Create(hc)
}

func GetHeroClasses() []*data.HeroClass {
	res := []*data.HeroClass{}
	db.Find(&res)
	return res
}
