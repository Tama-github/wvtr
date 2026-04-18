package databasemodel

import (
	"time"
)

var Heroes = []*Hero{
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

var Teams = []*Team{
	{
		Heroes: []*Hero{},
	},
	{
		Heroes: []*Hero{Heroes[1], Heroes[0], Heroes[0]},
	},
	{
		Heroes: []*Hero{Heroes[1], Heroes[0], Heroes[2]},
	},
	{
		Heroes: []*Hero{Heroes[2], Heroes[2], Heroes[2]},
	},
	{
		Heroes: []*Hero{Heroes[3], Heroes[4], Heroes[5]},
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

var Users = []*User{
	{
		Name: "Tama",
		State: &GameState{
			State: Home,
		},
		CurrentTeam:    Teams[0],
		LastActionTime: time.Now(),
		OwnedHeroes:    []*Hero{Heroes[3], Heroes[4], Heroes[5], Heroes[1]},
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
