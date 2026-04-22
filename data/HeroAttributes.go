package data

type AttributesID int

const (
	MaxHPID AttributesID = iota
	StrengthID
	IntelligenceID
	DexterityID
	LuckID
)

type ResistancesID int

const (
	BluntID ResistancesID = iota
	SlashID
	PierceID
	FireID
	FrostID
	LightingID
)

type GrowthRateID int

const (
	HPgtID GrowthRateID = iota
	SgtID
	IgtID
	DgtID
	LgtID
)

func NewHeroAttribute(class *HeroClass, grModifier []float64) *HeroAttributes {
	return &HeroAttributes{
		Level:     0,
		CurrentXP: 0,
		CurrentHP: 0,

		MaxHP:        class.MaxHP,
		Strength:     class.Strength,
		Intelligence: class.Intelligence,
		Dexterity:    class.Dexterity,
		Luck:         class.Luck,
		HPgt:         class.Sgt + grModifier[HPgtID],
		Sgt:          class.Sgt + grModifier[SgtID],
		Igt:          class.Igt + grModifier[IgtID],
		Dgt:          class.Dgt + grModifier[DgtID],
		Lgt:          class.Lgt + grModifier[LgtID],
	}
}

func (a HeroAttributes) GetMaxHP() float64 {
	return a.MaxHP
}

func (a HeroAttributes) GetStrength() float64 {
	return a.Strength
}

func (a HeroAttributes) GetIntelligence() float64 {
	return a.Intelligence
}

func (a HeroAttributes) GetDexterity() float64 {
	return a.Dexterity
}

func (a HeroAttributes) GetLuck() float64 {
	return a.Luck
}

func (a HeroAttributes) GetBluntRes() float64 {
	return a.Blunt
}

func (a HeroAttributes) GetSlashRes() float64 {
	return a.Slash
}

func (a HeroAttributes) GetPierceRes() float64 {
	return a.Pierce
}

func (a HeroAttributes) GetFireRes() float64 {
	return a.Fire
}

func (a HeroAttributes) GetFrostRes() float64 {
	return a.Frost
}

func (a HeroAttributes) GetLightingRes() float64 {
	return a.Lighting
}

func (a HeroAttributes) GetAttributesArray() []float64 {
	res := make([]float64, 5)
	res[MaxHPID] = a.GetMaxHP()
	res[StrengthID] = a.GetStrength()
	res[IntelligenceID] = a.GetIntelligence()
	res[DexterityID] = a.GetDexterity()
	res[LuckID] = a.GetLuck()
	return res
}

func (a HeroAttributes) GetGRArray() []float64 {
	res := make([]float64, 5)
	res[HPgtID] = a.HPgt
	res[SgtID] = a.Sgt
	res[IgtID] = a.Igt
	res[DgtID] = a.Dgt
	res[LgtID] = a.Lgt
	return res
}

func (a *HeroAttributes) SetAttributesWithArray(arr []float64) {
	a.MaxHP = arr[MaxHPID]
	a.Strength = arr[StrengthID]
	a.Intelligence = arr[IntelligenceID]
	a.Dexterity = arr[DexterityID]
	a.Luck = arr[LuckID]
}

func (a HeroAttributes) LevelThreshold() float64 {
	x := (float64(a.Level) * 5.0) / 2.0
	return x * x
}
