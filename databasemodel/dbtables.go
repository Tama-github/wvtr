package databasemodel

import (
	"time"

	"gorm.io/gorm"
)

type ModelBase struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Team struct {
	ModelBase
	Heroes []*Hero `gorm:"many2many:team_heroes;" json:"heroes"`
}

type EncounterState int

const (
	Home EncounterState = iota + 1
	Travel
	Fight
	Neutral
	Error
)

func (e EncounterState) String() string {
	res := "Error"
	switch e {
	case Home:
		res = "Home"
	case Travel:
		res = "Travel"
	case Fight:
		res = "Fight"
	case Neutral:
		res = "Neutral"
	case Error:
		res = "Error"
	}
	return res
}

type GameState struct {
	ModelBase
	State               EncounterState `json:"state"`
	CurrentExpedition   *ExpeditionDB  `gorm:"foreignKey:CurrentExpeditionID" json:"currentExpedition"`
	CurrentExpeditionID uint           `json:"-"` // foreign key
	ETeam               *Team          `json:"eTeam"`
	ETeamID             uint           `json:"-"` // foreign key
	UserID              uint           `json:"-"` // foreign key
}

type ExpeditionStepResolveInfo struct {
	ModelBase
	StepInfos      string         `json:"stepInfos"`
	StepEndAt      time.Time      `json:"stepEndAt"`
	StepState      EncounterState `json:"stepState"`
	ExpeditionDBID uint           `json:"-"` // foreign key
}

type ExpeditionDB struct {
	ModelBase
	Identifier   string                       `json:"identifier"`
	StartedAt    time.Time                    `json:"startedAt"`
	WhatHappened []*ExpeditionStepResolveInfo `json:"whatHappened"`
	UserID       uint                         `json:"-"` // foreign key
}

type User struct {
	ModelBase
	Name           string     `json:"name"`
	State          *GameState `json:"state"`
	CurrentTeam    *Team      `json:"currentTeam"`
	CurrentTeamID  uint       `json:"-"` // foreign key
	LastActionTime time.Time  `json:"lastActionTime"`
	OwnedHeroes    []*Hero    `json:"ownedHeroes"`
	DiscordID      string     `gorm:"unique" json:"discord_id"`
}
