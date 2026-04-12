package stypes

import (
	"time"
)

type Waifu struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"imageUrl"`
}

type Team struct {
	Id 	   int		`json:"id"`
	Waifus [3]Waifu `json:"waifus"`
}

type EncounterState int;
const (
    Home EncounterState = iota + 1
    Travel
	Fight
    Neutral
	Error
)

type GameState struct {
    IsBusy  bool           `json:"isBusy"`
    State 	EncounterState `json:"state"`
    WTeam 	Team           `json:"wTeam"`
    ETeam 	Team           `json:"eTeam"`
}

type User struct {
	Id 				int		  `json:"id"`
	Name			string	  `json:"name"`
	State 			GameState `json:"state"`
	CurrentTeam 	Team	  `json:"currentTeam"`
	LastActionTime  time.Time `json:"lastActionTime"`
}
