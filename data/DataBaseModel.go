package data

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

type HeroAttributes struct {
	ModelBase

	Level     int     `json:"level"`
	CurrentXP float64 `json:"currentXP"`
	XPToLvlUp float64 `json:"xpBeforLvlUp"`
	CurrentHP float64 `json:"currentHP"`

	//Attributes
	MaxHP        float64 `json:"maxHP"`
	Strength     float64 `json:"strength"`
	Intelligence float64 `json:"intelligence"`
	Dexterity    float64 `json:"dexterity"`
	Luck         float64 `json:"luck"`

	//Growth rate
	HPgt float64 `json:"hpgt"`
	Sgt  float64 `json:"sgt"`
	Igt  float64 `json:"igt"`
	Dgt  float64 `json:"dgt"`
	Lgt  float64 `json:"lgt"`

	// Resistances
	Blunt    float64 `json:"blunt"`
	Pierce   float64 `json:"pierce"`
	Slash    float64 `json:"slash"`
	Fire     float64 `json:"fire"`
	Frost    float64 `json:"frost"`
	Lighting float64 `json:"lighting"`

	// fk
	HeroID uint `json:"-"`
}

type ExpeditionStepTimestamp struct {
	ModelBase
	When                        time.Time `json:"when"`
	What                        string    `json:"what"`
	ExpeditionStepResolveInfoID uint      `json:"-"` // foreign key
}

type ExpeditionStepResolveInfo struct {
	ModelBase
	StepState      EncounterState             `json:"stepState"`
	Timeline       []*ExpeditionStepTimestamp `json:"timeline"`
	ExpeditionDBID uint                       `json:"-"` // foreign key
}

type ExpeditionDB struct {
	ModelBase
	Identifier   string                       `json:"identifier"`
	StartedAt    time.Time                    `json:"startedAt"`
	WhatHappened []*ExpeditionStepResolveInfo `json:"whatHappened"`
	UserID       uint                         `json:"-"` // foreign key
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

type HeroClass struct {
	ModelBase
	Name        string      `json:"name"`
	Identifier  HeroClassID `gorm:"unique" json:"-"`
	Descritpion string      `json:"description"`
	Weight      float64     `json:"-"`

	// Base attribute
	//Attributes
	MaxHP        float64 `json:"-"`
	Strength     float64 `json:"-"`
	Intelligence float64 `json:"-"`
	Dexterity    float64 `json:"-"`
	Luck         float64 `json:"-"`

	//Base Growth rate
	HPgt float64 `json:"-"`
	Sgt  float64 `json:"-"`
	Igt  float64 `json:"-"`
	Dgt  float64 `json:"-"`
	Lgt  float64 `json:"-"`
}

type Skill struct {
	ModelBase
	Identifier  SkillID   `gorm:"unique" json:"identifier"`
	Name        string    `json:"name"`
	Type        SkillType `json:"skill_type"`
	ImageURL    string    `json:"image_url"`
	Description string    `json:"description"`
	Weight      float64   `json:"weight"`
}

type Hero struct {
	ModelBase
	ImageUrl      string          `json:"imageUrl"`
	Name          string          `json:"name"`
	Class         *HeroClass      `gorm:"foreignkey:HeroClassID" json:"heroClass"`
	HeroClassID   uint            `json:"-"`
	Rank          string          `json:"rank"`
	Attributes    *HeroAttributes `json:"attributes"`
	UniqueSkill   *Skill          `gorm:"foreignkey:UniqueSkillID" json:"uniqueSkill"`
	UniqueSkillID uint            `json:"-"`
	ActiveSkill   *Skill          `gorm:"foreignkey:ActiveSkillID" json:"activeSkill"`
	ActiveSkillID uint            `gorm:"" json:"-"`
	UserID        uint            `gorm:"" json:"-"` // foreign key

	// info that we save to request nanapi if we need to.
	WaifuID        string `gorm:"unique" json:"id_w"` // not foreign
	AnilistCharaID uint   `json:"id_al"`              // not foreign
}

type Team struct {
	ModelBase
	Heroes []*Hero `gorm:"many2many:team_heroes;" json:"heroes"`
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

var DBSchema []any = []any{
	&HeroAttributes{},
	&ExpeditionStepTimestamp{},
	&ExpeditionStepResolveInfo{},
	&ExpeditionDB{},
	&GameState{},
	&HeroClass{},
	&Skill{},
	&Hero{},
	&Team{},
	&User{},
}
