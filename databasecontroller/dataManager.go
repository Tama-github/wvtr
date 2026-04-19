package databasecontroller

import (
	"fmt"
	"log"
	"wvtrserv/databasemodel"
	"wvtrserv/gameexpedition"
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

	// this is to check if we need to populate the db if it is empty with the dev values.
	u := GetUserByID(1)
	if u.ID == 0 {
		// Lets do it once
		createDBDev(db)
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

func createDBDev(db *gorm.DB) {
	db.AutoMigrate(&databasemodel.HeroAttributes{})
	db.AutoMigrate(&databasemodel.Hero{})
	db.AutoMigrate(&databasemodel.Team{})
	db.AutoMigrate(&databasemodel.GameState{})
	db.AutoMigrate(&databasemodel.User{})
	db.AutoMigrate(&databasemodel.ExpeditionStepResolveInfo{})
	db.AutoMigrate(&databasemodel.ExpeditionDB{})

	for _, o := range databasemodel.Heroes {
		result := db.Create(o)
		if result.Error != nil {
			fmt.Printf("Could not insert: ")
			fmt.Printf("%s", result.Error.Error())
		}
	}
	for _, o := range databasemodel.Teams {
		result := db.Create(o)
		if result.Error != nil {
			fmt.Printf("Could not insert: ")
			fmt.Printf("%s", result.Error.Error())
		}
	}
	for _, o := range databasemodel.Users {
		result := db.Create(o)
		if result.Error != nil {
			fmt.Printf("Could not insert: ")
			fmt.Printf("%s", result.Error.Error())
		}
	}
}

// DB requests
func GetHeroByID(id uint) databasemodel.Hero {
	var res databasemodel.Hero
	db.Preload("GameState").Preload("Team").Preload("Hero").Find(&res, id)
	fmt.Printf("Got hero by ID: %d\n", id)
	return res
}

func GetTeamByID(id uint) databasemodel.Team {
	var res databasemodel.Team
	db.Preload("GameState").Preload("Team").Preload("Hero").Find(&res, id)
	return res
}

func GetUserByID(id uint) *databasemodel.User {
	var res *databasemodel.User = nil
	// We get the full user values.
	db.Preload("CurrentTeam").
		Preload("CurrentTeam.Heroes").
		Preload("State").
		Preload("State.CurrentExpedition").
		Preload("State.CurrentExpedition.WhatHappened").
		Preload("State.ETeam").
		Preload("State.ETeam.Heroes").
		Preload("OwnedHeroes").
		Find(&res, id)

	logger.DumpLog.Println("Get user by id ", id)
	logger.DumpLog.Println(res)

	return res
}

func GetUserByDiscordID(did string) *databasemodel.User {
	var res *databasemodel.User = nil
	db.Where("discord_id = ?", did).Find(&res)
	if res != nil {
		logger.DumpLog.Println("GetUserByDiscordID: ", did, " | ", res.Name)
	} else {
		logger.DumpLog.Println("GetUserByDiscordID: ", did, " | user not found.")
	}
	return res
}

func CreateExpeditionDB(edb *databasemodel.ExpeditionDB) {
	db.Save(edb)
}

func UpdateUser(user *databasemodel.User) {
	db.Save(user)
}

func CreateNewUser(user *databasemodel.User) {
	logger.DumpLog.Print("CreateNewUser")
	db.Create(user)
}

func CreateHero(hero *databasemodel.Hero) {
	logger.DumpLog.Print("CreateNewHero")
	db.Create(hero)
}

func UpdateTeam(user *databasemodel.User) {
	db.Model(&user.CurrentTeam).Association("Heroes").Replace(user.CurrentTeam.Heroes)
}

func UpdateGameState(state *databasemodel.GameState) {
	db.Save(state)
}

func LaunchExpedition(user *databasemodel.User, expedition *databasemodel.ExpeditionDB) {
	// TODO: Check expedition integrity
	if UserIsHome(user) && !UserHasAProblem(user) {
		user.State.CurrentExpedition = expedition
		user.State.State = expedition.WhatHappened[0].StepState
		if user.State.State == databasemodel.Fight {
			user.State.ETeam = gameexpedition.GetEnemyTeamForEvent(expedition.Identifier, 0)
		}
		UpdateGameState(user.State)
	} else {
		logger.ErrLog.Printf("requested to launch an expedition on user %s but user is [state: %s].", user.Name, user.State.State)
	}
}
