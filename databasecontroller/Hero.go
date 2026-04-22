package databasecontroller

import (
	"fmt"
	"wvtrserv/data"
	"wvtrserv/logger"
)

func CreateHero(hero *data.Hero) error {
	logger.DumpLog.Print("CreateNewHero")
	tx := db.Create(hero)
	return tx.Error
}

func GetHeroByID(id uint) data.Hero {
	var res data.Hero
	db.Preload("HeroClass").
		Preload("Attributes").
		Preload("Class").
		Preload("UniqueSkill").
		Preload("ActiveSkill").
		Preload("Team").
		Preload("Hero").
		Find(&res, id)
	fmt.Printf("Got hero by ID: %d\n", id)
	return res
}
