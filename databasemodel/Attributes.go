package databasemodel

type HeroAttributes struct {
	ModelBase

	//Attributes
	Strength     int `json:"strength"`
	Intelligence int `json:"intelligence"`
	Dexterity    int `json:"dexterity"`
	Luck         int `json:"luck"`

	//Growth rate
	Sgt float32 `json:"sgt"`
	Igt float32 `json:"igt"`
	Dgt float32 `json:"dgt"`
	Lgt float32 `json:"lgt"`

	// Resistances
	Blunt    int `json:"blunt"`
	Pierce   int `json:"pierce"`
	Slash    int `json:"slash"`
	Fire     int `json:"fire"`
	Frost    int `json:"frost"`
	Lighting int `json:"lighting"`

	// fk
	HeroID uint `json:"-"`
}
