package databasemodel

type Hero struct {
	ModelBase
	Name          string          `json:"name"`
	ImageUrl      string          `json:"imageUrl"`
	Rank          string          `json:"rank"`
	Level         int             `json:"level"`
	CurrentXP     int             `json:"currentXP"`
	XPBeforeLvlUp int             `json:"xpBeforLvlUp"`
	CurrentHP     int             `json:"currentHP"`
	MaxHP         int             `json:"maxHP"`
	Attributes    *HeroAttributes `json:"attributes"`
	UserID        uint            `gorm:"" json:"-"` // foreign key

	// info that we save to request nanapi if we need to.
	WaifuID        string `json:"id_w"`  // not foreign
	AnilistCharaID uint   `json:"id_al"` // not foreign
}
