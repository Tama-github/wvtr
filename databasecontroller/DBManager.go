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
	db, err = gorm.Open(sqlite.Open("./madb.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database")
	}

	// Skills must be in db, if there are not this means that the db is empty
	// So we create the DB and populate it with everything that is needed (classes, skills)
	s := GetSkills()
	if len(s) == 0 {
		// Lets do it once
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
}

func LaunchExpedition(user *data.User, expedition *data.ExpeditionDB) {
	// TODO: Check expedition integrity
	if user.UserIsHome() && !user.UserHasAProblem() {
		user.State.CurrentExpedition = expedition
		user.State.State = expedition.WhatHappened[0].StepState
		if user.State.State == data.Fight {
			glTeam := gamedata.GetEnemyTeamForEvent(expedition.Identifier, 0)
			user.State.ETeam = (*data.Team)(glTeam)
		}
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
	for _, s := range data.GetSkills() {
		CreateSkill(s)
	}
}

func InsertHeroClassesInDB() {
	// check if there if it is necessary
	skills := GetHeroClasses()
	if len(skills) > 0 {
		logger.DumpLog.Println("Hero classes already inside db")
		return
	}
	logger.DumpLog.Println("Insert hero classes in db")
	for _, s := range data.GetHeroClasses() {
		CreateHeroClasse(s)
	}
}
