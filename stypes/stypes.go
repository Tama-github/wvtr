package stypes

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ModelBase struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Hero struct {
	ModelBase
	Name         string `json:"name"`
	ImageUrl     string `json:"imageUrl"`
	Level        int    `json:"level"`
	CurrentXP    int    `json:"currentXP"`
	XPBeforLvlUp int    `json:"xpBeforLvlUp"`
	CurrentHP    int    `json:"currentHP"`
	MaxHP        int    `json:"maxHP"`
	UserID       uint   `json:"-"` // foreign key
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

type GameState struct {
	ModelBase
	State               EncounterState `json:"state"`
	CurrentExpedition   *ExpeditionDB  `gorm:"foreignKey:CurrentExpeditionID" json:"currentExpedition"`
	CurrentExpeditionID uint           `json:"-"` // foreign key
	ETeam               *Team          `json:"eTeam"`
	ETeamID             uint           `json:"-"` // foreign key
	UserID              uint           `json:"-"` // foreign key
}

func (g *GameState) ResolveGameState(t *time.Time) *ExpeditionStepResolveInfo {
	idx := -1
	var currentExpStep *ExpeditionStepResolveInfo = nil
	if g.CurrentExpedition != nil {
		idx, currentExpStep = g.CurrentExpedition.GetCurrentStep(t)
	}
	if idx < 0 {
		g.State = Home
		g.CurrentExpedition = nil
		g.ETeam = nil
		return nil
	}

	g.State = currentExpStep.StepState
	if g.State == Fight {
		g.ETeam = Expeditions[g.CurrentExpedition.Identifier].GetEnemyTeamForEvent(idx)
	}
	return currentExpStep
}

func (g *GameState) LaunchExpedition(expIdentifier string, pTeam *Team) {
	newExpedition, ok := Expeditions[expIdentifier]
	if ok {
		g.CurrentExpedition = newExpedition.Solve(expIdentifier, pTeam)
		CreateExpeditionDB(g.CurrentExpedition)
		g.State = g.CurrentExpedition.WhatHappened[0].StepState
		UpdateGameState(g)
	} else {
		fmt.Printf("[%s] is not an existing expedition.\n", expIdentifier)
	}
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

func (e ExpeditionDB) GetCurrentStep(t *time.Time) (int, *ExpeditionStepResolveInfo) {
	for i, step := range e.WhatHappened {
		if step.StepEndAt.After(*t) {
			return i, step
		}
	}
	return -1, nil
}

type User struct {
	ModelBase
	Name           string     `json:"name"`
	State          *GameState `json:"state"`
	CurrentTeam    *Team      `json:"currentTeam"`
	CurrentTeamID  uint       `json:"-"` // foreign key
	LastActionTime time.Time  `json:"lastActionTime"`
	OwnedHeroes    []*Hero    `json:"ownedHeroes"`
}

func (u User) isHome() bool {
	return u.State.State != Error && u.State.State == Home
}

func (u *User) LaunchExpedition(expIdentifier string) {
	if u.isHome() {
		u.State.LaunchExpedition(expIdentifier, u.CurrentTeam)
	}
}

func (u *User) ResolveExpeditionState(t *time.Time) *ExpeditionStepResolveInfo {
	res := u.State.ResolveGameState(t)

	return res
}
