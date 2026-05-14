package databasecontroller

import (
	"fmt"
	"log"
	"wvtrserv/data"
	"wvtrserv/gamedata"
	"wvtrserv/logger"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

var db *gorm.DB

func DBLogIn() *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open("./db/madb.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database")
	}

	// Skills must be in db, if there are not, this means that the db is empty
	// So we create the DB and populate it with everything that is needed (classes, skills, etc...)
	s := GetSkills()
	if len(s) == 0 {
		createDB(db)
	}

	return db
}

func InsertInDB(db *gorm.DB, toAdd interface{}) {
	result := db.Create(&toAdd)
	if result.Error != nil {
		fmt.Printf("Could not insert: ")
		fmt.Printf("%s", result.Error.Error())
	}
}

func CreateTable(db *gorm.DB, toAdd interface{}) {
	result := db.Create(&toAdd)
	if result.Error != nil {
		fmt.Printf("Could not insert: ")
		fmt.Printf("%s", result.Error.Error())
	}
}

func createDB(db *gorm.DB) {
	for _, a := range data.DBSchema {
		db.AutoMigrate(a)
	}

	InsertSkillsInDB()
	InsertHeroClassesInDB()
	InsertAllCurrenciesInDB()
	InsertEveryEnemiesInDB()
}

func Delete[T any](row T) {
	db.Delete(row)
}

func LaunchExpedition(user *data.User, expedition *data.ExpeditionDB) {
	// TODO: Check expedition integrity
	if user.UserIsHome() && !user.UserHasAProblem() {
		CreateExpeditionDB(expedition)
		user.State.CurrentExpedition = expedition
		user.State.State = expedition.WhatHappened[0].StepState

		UpdateGameState(user.State)
	} else {
		logger.ErrLog.Printf("requested to launch an expedition on user %s but user is [state: %s].", user.Name, user.State.State)
	}
}

func InsertSkillsInDB() {
	// check if there if it is necessary
	skills := GetSkills()
	if len(skills) > 0 {
		logger.DumpLog.Println("Skills already inside db")
		return
	}
	logger.DumpLog.Println("Insert skills in db")
	for _, s := range gamedata.GetSkills() {
		CreateSkill(s)
	}
}

func InsertHeroClassesInDB() {
	// check if there if it is necessary
	heroclasses := GetHeroClasses()
	if len(heroclasses) > 0 {
		logger.DumpLog.Println("Hero classes already inside db")
		return
	}
	logger.DumpLog.Println("Insert hero classes in db")
	for _, s := range gamedata.GetHeroClasses() {
		CreateHeroClass(s)
	}
}

func InsertEveryEnemiesInDB() {
	// check if there if it is necessary
	heroes := GetHeroes()
	if len(heroes) > 0 {
		logger.DumpLog.Println("Hero classes already inside db")
		return
	}
	logger.DumpLog.Println("Insert enemies in db")
	for _, s := range gamedata.GetEveryEnemies() {
		CreateHero(s)
	}
}

func InsertAllCurrenciesInDB() {
	// check if there if it is necessary
	heroes := GetAllCurrencies()
	if len(heroes) > 0 {
		logger.DumpLog.Println("Currencies already inside db")
		return
	}
	logger.DumpLog.Println("Insert currencies in db")
	for _, s := range gamedata.GetAllCurrencies() {
		CreateCurrency(s)
	}
}
