package stypes

import (
	"fmt"
	"log"
	"time"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

var heroes = []*Hero{
	{
		Name:         "Error",
		ImageUrl:     "/imgs/noimage.jpg",
		Level:        0,
		CurrentXP:    0,
		XPBeforLvlUp: 0,
		CurrentHP:    0,
		MaxHP:        0,
	},
	{
		Name:         "Otaku",
		ImageUrl:     "/imgs/otaku.png",
		Level:        1,
		CurrentXP:    0,
		XPBeforLvlUp: 0,
		CurrentHP:    20,
		MaxHP:        20,
	},
	{
		Name:         "Otaku 2",
		ImageUrl:     "/imgs/otaku2.png",
		Level:        2,
		CurrentXP:    0,
		XPBeforLvlUp: 0,
		CurrentHP:    30,
		MaxHP:        30,
	},
	{
		Name:         "Emilia",
		ImageUrl:     "/imgs/emilia.jpg",
		Level:        3,
		CurrentXP:    21,
		XPBeforLvlUp: 50,
		CurrentHP:    32,
		MaxHP:        52,
	},
	{
		Name:         "Eren",
		ImageUrl:     "/imgs/eren.jpg",
		Level:        4,
		CurrentXP:    5,
		XPBeforLvlUp: 60,
		CurrentHP:    2,
		MaxHP:        40,
	},
	{
		Name:         "Satoru",
		ImageUrl:     "/imgs/satoru.png",
		Level:        3,
		CurrentXP:    42,
		XPBeforLvlUp: 50,
		CurrentHP:    45,
		MaxHP:        45,
	},
}

var teams = []*Team{
	{
		Heroes: []*Hero{},
	},
	{
		Heroes: []*Hero{heroes[1], heroes[0], heroes[0]},
	},
	{
		Heroes: []*Hero{heroes[1], heroes[0], heroes[2]},
	},
	{
		Heroes: []*Hero{heroes[2], heroes[2], heroes[2]},
	},
	{
		Heroes: []*Hero{heroes[3], heroes[4], heroes[5]},
	},
}

// var states = []GameState{
// 	{
// 		IsBusy: false,
// 		State:  Error,
// 	},
// }

// var expeditions = []*ExpeditionDB{
// 	{
// 		Identifier: "travel30s",
// 	},
// 	{
// 		Identifier: "travel40s",
// 	},
// }

var users = []*User{
	{
		Name: "Tama",
		State: &GameState{
			State: Home,
		},
		CurrentTeam:    teams[0],
		LastActionTime: time.Now(),
		OwnedHeroes:    []*Hero{heroes[3], heroes[4], heroes[5], heroes[1]},
	},
	// {
	// 	Name: "Uroja",
	// 	State: &GameState{
	// 		IsBusy: true,
	// 		State:  Travel,
	// 	},
	// 	CurrentTeam:    teams[4],
	// 	LastActionTime: time.Now(),
	// },
	// {
	// 	Name: "Robinet",
	// 	State: &GameState{
	// 		IsBusy: true,
	// 		State:  Fight,
	// 	},
	// 	CurrentTeam:    teams[4],
	// 	LastActionTime: time.Now(),
	// },
	// {
	// 	Name: "TestUser123",
	// 	State: &GameState{
	// 		IsBusy: false,
	// 		State:  Neutral,
	// 	},
	// 	CurrentTeam:    teams[4],
	// 	LastActionTime: time.Now(),
	// },
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

func createDBBase(db *gorm.DB) {
	db.AutoMigrate(&Hero{})
	db.AutoMigrate(&Team{})
	db.AutoMigrate(&GameState{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&ExpeditionStepResolveInfo{})
	db.AutoMigrate(&ExpeditionDB{})

	for _, o := range heroes {
		result := db.Create(o)
		if result.Error != nil {
			fmt.Printf("Could not insert: ")
			fmt.Printf("%s", result.Error.Error())
		}
	}
	for _, o := range teams {
		result := db.Create(o)
		if result.Error != nil {
			fmt.Printf("Could not insert: ")
			fmt.Printf("%s", result.Error.Error())
		}
	}
	for _, o := range users {
		result := db.Create(o)
		if result.Error != nil {
			fmt.Printf("Could not insert: ")
			fmt.Printf("%s", result.Error.Error())
		}
	}
}

var db *gorm.DB

func DBLogIn() *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open("./madb.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database!")
	}
	u := GetUserByID(1)

	if u.ID == 0 {
		// Lets do it once
		createDBBase(db)
	}

	return db
}

func GetHeroByID(id uint) Hero {
	var res Hero
	db.Preload("GameState").Preload("Team").Preload("Hero").Find(&res, id)
	fmt.Printf("Got hero by ID: %d\n", id)
	return res
}

func GetTeamByID(id uint) Team {
	var res Team
	db.Preload("GameState").Preload("Team").Preload("Hero").Find(&res, id)
	return res
}

func GetUserByID(id uint) *User {
	var res *User = nil
	db.Preload("CurrentTeam").
		Preload("CurrentTeam.Heroes").
		Preload("State").
		Preload("State.CurrentExpedition").
		Preload("State.CurrentExpedition.WhatHappened").
		Preload("State.ETeam").
		Preload("State.ETeam.Heroes").
		Preload("OwnedHeroes").
		Find(&res, id)
	//db.Preload("GameState").Preload("Team").Preload("Hero").Find(&res, id)
	return res
}

func CreateExpeditionDB(edb *ExpeditionDB) {
	db.Save(edb)
}

func UpdateUser(user *User) {
	db.Save(user)
}

func UpdateTeam(user *User) {
	db.Model(&user.CurrentTeam).Association("Heroes").Replace(user.CurrentTeam.Heroes)
}

func UpdateGameState(state *GameState) {
	db.Save(state)
}
