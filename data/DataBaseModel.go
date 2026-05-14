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

type Damage struct {
	ModelBase
	SlashDmg     float64 `json:"slashDmg"`
	BluntDmg     float64 `json:"bluntDmg"`
	PierceDmg    float64 `json:"pierceDmg"`
	FireDmg      float64 `json:"fireDmg"`
	FrostDmg     float64 `json:"frostDmg"`
	LightningDmg float64 `json:"lightningDmg"`

	// polumorphism
	OwnerID   uint   `json:"-"`
	OwnerType string `json:"-"`
}

type StatsRange struct {
	ModelBase
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Value float64 `json:"value"`

	// polumorphism
	OwnerID   uint   `json:"-"`
	OwnerType string `json:"-"`
}

type Affix struct {
	ModelBase
	Name   string        `json:"name"`
	Ranges []*StatsRange `json:"ranges" gorm:"polymorphic:Owner;polymorphicId:OwnerID"`
	Type   AffixType     `json:"type"`

	// polumorphism
	OwnerID   uint   `json:"-"`
	OwnerType string `json:"-"`
}

type Storable struct {
	ModelBase
	Name    string `json:"name"`
	IconURL string `json:"iconURL"`

	// fk
	InventoryID uint `json:"-"`
}

type Currency struct {
	Storable
	Type CurrencyType `json:"type"`
}

type CurrencyOwned struct {
	ModelBase
	NumberOwned int       `json:"numberOwned"`
	Currency    *Currency `json:"currency"`

	// fk
	CurrencyID  uint `json:"-"`
	InventoryID uint `json:"-"`
}

type Equipable struct {
	Storable
	RealWeightScore float64  `json:"realWeightScore"`
	Affixes         []*Affix `json:"affixes" gorm:"polymorphic:Owner;polymorphicId:OwnerID"`
}

type Weapon struct {
	Equipable
	BaseDamage      *Damage     `json:"baseDamage" gorm:"polymorphic:Owner;polymorphicId:OwnerID"`
	BaseCritRate    *StatsRange `json:"baseCritRate" gorm:"polymorphic:Owner;foreignKey:BaseCritRateID"`
	BaseAttackSpeed *StatsRange `json:"baseAttackSpeed" gorm:"polymorphic:Owner;foreignKey:BaseAttackSpeedID"`

	// Damage scaling
	StrScaling float64 `json:"strScaling"`
	IntScaling float64 `json:"intScaling"`
	DexScaling float64 `json:"dexScaling"`
	LckScaling float64 `json:"lckScaling"`

	//fk
	BaseCritRateID    uint `json:"-"`
	BaseAttackSpeedID uint `json:"-"`
}

type Armor struct {
	Equipable
	BlockScore           *StatsRange `json:"blockScore" gorm:"polymorphic:Owner;polymorphicId:OwnerID;foreignKey:BlockScoreID"`
	EvadeScore           *StatsRange `json:"evadeScore" gorm:"polymorphic:Owner;polymorphicId:OwnerID;foreignKey:EvadeScoreID"`
	BaseResistancesRange *Damage     `json:"baseResistancesRange" gorm:"polymorphic:Owner;polymorphicId:OwnerID"`

	//fk
	BlockScoreID uint `json:"-"`
	EvadeScoreID uint `json:"-"`
}

type Omamori struct {
	Equipable
}

type HeroEquipment struct {
	ModelBase
	Weapon  *Weapon  `json:"weapon"`
	Armor   *Armor   `json:"armor"`
	Omamori *Omamori `json:"omamori"`

	//fk
	WeaponID  uint `json:"-"`
	ArmorID   uint `json:"-"`
	OmamoriID uint `json:"-"`
}

type Inventory struct {
	ModelBase
	Weapons    []*Weapon        `json:"weapons"`
	Armors     []*Armor         `json:"armors"`
	Omamoris   []*Omamori       `json:"omamoris"`
	Currencies []*CurrencyOwned `json:"currencies"`
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

	//Defense
	BlockScore float64 `json:"blockScore"`
	EvadeScore float64 `json:"evadeScore"`

	// Resistances
	Blunt     float64 `json:"blunt"`
	Pierce    float64 `json:"pierce"`
	Slash     float64 `json:"slash"`
	Fire      float64 `json:"fire"`
	Frost     float64 `json:"frost"`
	Lightning float64 `json:"lighting"`

	// fk
	HeroID uint `json:"-"`
}

type Reward struct {
	ModelBase
	XP   float64    `json:"xp"`
	Loot *Inventory `json:"loot"`

	// fk
	LootID uint `json:"-"`
}

type FieldActionDesc struct {
	ModelBase
	FromH          *Hero                `json:"fromH"`
	UsedSkill      *Skill               `json:"usedSKill"`
	TargetH        *Hero                `json:"targetH"`
	TargetStatus   HeroTakeDamageStatus `json:"targetStatus"`
	FromPVChange   float64              `json:"fromPVChange"`
	TargetPVChange float64              `json:"targetPVChange"`

	// fk
	FromHID     uint `json:"-"`
	UsedSkillID uint `json:"-"`
	TargetHID   uint `json:"-"`
}

type ExpeditionStepTimestamp struct {
	ModelBase
	When       time.Time        `json:"when"`
	What       string           `json:"what"`
	WhatAction *FieldActionDesc `json:"whatAction"`

	// fk
	ExpeditionStepResolveInfoID uint `json:"-"`
	WhatActionID                uint `json:"-"`
}

type ExpeditionStepResolveInfo struct {
	ModelBase
	StepState EncounterState             `json:"stepState"`
	Timeline  []*ExpeditionStepTimestamp `json:"timeline"`
	ETeam     *Team                      `json:"eTeam"`

	// fk
	ETeamID        uint `json:"-"`
	ExpeditionDBID uint `json:"-"`
}

type ExpeditionDB struct {
	ModelBase
	Identifier        string                       `json:"identifier"`
	StartedAt         time.Time                    `json:"startedAt"`
	WhatHappened      []*ExpeditionStepResolveInfo `json:"whatHappened"`
	ExpeditionRewards *Reward                      `json:"ExpeditionRewards"`

	// fk
	ExpeditionRewardsID uint `json:"-"`
}

type GameState struct {
	ModelBase
	State             EncounterState `json:"state"`
	CurrentExpedition *ExpeditionDB  `json:"currentExpedition"`

	// fk
	CurrentExpeditionID uint `json:"-"`
}

type HeroClass struct {
	ModelBase
	Name         string      `json:"name"`
	Identifier   HeroClassID `gorm:"unique" json:"-"`
	Descritpion  string      `json:"description"`
	ClassIconURL string      `json:"class_icon_url"`
	Weight       float64     `json:"-"`

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
	Identifier           SkillID       `gorm:"unique" json:"identifier"`
	Name                 string        `json:"name"`
	Type                 SkillType     `json:"skill_type"`
	Targeting            TargetType    `json:"target_type"`
	RecuperationDuration time.Duration `json:"recuperation_duration"`
	ImageURL             string        `json:"image_url"`
	Description          string        `json:"description"`
	Weight               float64       `json:"-"`
}

type Hero struct {
	ModelBase
	ImageUrl   string          `json:"imageUrl"`
	Name       string          `json:"name"`
	Class      *HeroClass      `gorm:"foreignkey:HeroClassID" json:"heroClass"`
	Rank       string          `json:"rank"`
	Attributes *HeroAttributes `json:"attributes"`

	// Skills
	WeaponAttack *Skill `gorm:"foreignkey:WeaponAttackID" json:"weaponAttack"`
	UniqueSkill  *Skill `gorm:"foreignkey:UniqueSkillID" json:"uniqueSkill"`
	ActiveSkill  *Skill `gorm:"foreignkey:ActiveSkillID" json:"activeSkill"`

	// Items
	Equipment *HeroEquipment `json:"equipment"`

	// info that we save to request nanapi if we need to.
	WaifuID        *string `gorm:"unique" json:"id_w"` // not foreign
	AnilistCharaID uint    `json:"id_al"`              // not foreign

	// fk
	UserID         uint `json:"-"`
	WeaponAttackID uint `json:"-"`
	ActiveSkillID  uint `json:"-"`
	UniqueSkillID  uint `json:"-"`
	HeroClassID    uint `json:"-"`
	EquipmentID    uint `json:"-"`
}

type Team struct {
	ModelBase
	Heroes []*Hero `gorm:"many2many:team_heroes;" json:"heroes"`
}

type User struct {
	ModelBase
	Name           string     `json:"name"`
	Inventory      *Inventory `json:"inventory"`
	State          *GameState `json:"state"`
	CurrentTeam    *Team      `json:"currentTeam"`
	LastActionTime time.Time  `json:"lastActionTime"`
	OwnedHeroes    []*Hero    `json:"ownedHeroes"`
	DiscordID      string     `gorm:"unique" json:"discord_id"`

	// fk
	InventoryID   uint `json:"-"`
	StateID       uint `json:"-"`
	CurrentTeamID uint `json:"-"`
}

var DBSchema []any = []any{
	&Damage{},
	&StatsRange{},
	&Affix{},
	&Currency{},
	&CurrencyOwned{},
	&Weapon{},
	&Armor{},
	&Omamori{},
	&HeroEquipment{},
	&Inventory{},
	&HeroAttributes{},
	&Reward{},
	&FieldActionDesc{},
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
