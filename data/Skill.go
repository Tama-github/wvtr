package data

type TargetType int

const (
	Self TargetType = iota
	Enemy
	Friends
)

type SkillType int

const (
	Unique SkillType = iota
	Active
)

type SkillID int

const (
	// Unique
	Lucky SkillID = iota
	GoodRest
	SecondWind
	Prodigy
	Berserk
	Trickster
	FastLearner
	ElementalCursed
	PhysicalCursed

	// Active
	Spit
	Charge
	Tackle
	Sweep
	FireBolt
	IceBolt
	LightningBolt

	// Weapon attack
	WeaponAttack
)

func IsLuckyActivate(x float64) bool {
	roll := NaturalRoll(0, 1)
	target := 1 - ((x * x) / (200 * x))
	return RollCheck(roll, target)
}

func (s *Skill) UseLucky(from *Hero) bool {
	if s.Identifier != Lucky {
		return false
	}
	return IsLuckyActivate(float64(from.Attributes.GetLuck()))
}

func (s *Skill) UseSecondWind(from *Hero) bool {
	if s.Identifier != SecondWind {
		return false
	}
	if s.HaveBeenUsed {
		return false
	}
	s.HaveBeenUsed = true
	return true
}

func (s *Skill) UseTrickster(from *Hero) bool {
	if s.Identifier != Trickster {
		return false
	}
	lck := float64(from.Attributes.GetLuck())
	dex := float64(from.Attributes.GetDexterity())
	x := (dex + lck) / 2
	roll := NaturalRoll(0, 1)
	target := 1 - ((x * x) / (200 * x))
	return RollCheck(roll, target)
}

func (s *Skill) Use(from *Hero) bool {
	switch s.Identifier {
	case Lucky:
		return s.UseLucky(from)
	case GoodRest:
		return true
	case SecondWind:
		return s.UseSecondWind(from)
	case Prodigy:
		return true
	case Berserk:
		return true
	case Trickster:
		return s.UseTrickster(from)
	case FastLearner:
		return true
		// case ElementalCursed:
		// 	return true
		// case PhysicalCursed:
		// 	return true
	}
	return false
}

func WeaponAttackAction(from *Hero, target *Hero, fad *FieldActionDesc) *FieldActionDesc {
	var status HeroTakeDamageStatus = 0
	// get attack value
	dmg := from.GetAttackDamageForWeapon()

	// check miss

	// check critic
	if from.IsCrit() {
		dmg = dmg.GetDmgFunc(func(x float64) float64 {
			return x * from.GetTotalCritMulti()
		})
		status = status | Crit
	}

	// damage target
	target.TakeDamage(dmg, from, fad)
	fad.TargetStatus = fad.TargetStatus | status
	return fad
}

func SpitAction(from *Hero, target *Hero, fad *FieldActionDesc) *FieldActionDesc {
	// target.takeFlatDamage(1)
	// report := fmt.Sprintf("%s inficted %d dmg to %s.\n", from.Name, 1, target.Name)
	// if target.IsDefeated() {
	// 	report += fmt.Sprintf("%s defeated %s.\n", from.Name, target.Name)
	// }
	return fad
}

func ChargeAction(from *Hero, target *Hero, fad *FieldActionDesc) *FieldActionDesc {
	// target.takeFlatDamage(1)
	// report := fmt.Sprintf("%s inficted %d dmg to %s.\n", from.Name, 1, target.Name)
	// if target.IsDefeated() {
	// 	report += fmt.Sprintf("%s defeated %s.\n", from.Name, target.Name)
	// }
	return fad
}

func TackleAction(from *Hero, target *Hero, fad *FieldActionDesc) *FieldActionDesc {
	// target.takeFlatDamage(1)
	// report := fmt.Sprintf("%s inficted %d dmg to %s.\n", from.Name, 1, target.Name)
	// if target.IsDefeated() {
	// 	report += fmt.Sprintf("%s defeated %s.\n", from.Name, target.Name)
	// }
	return fad
}

func SweepAction(from *Hero, target *Hero, fad *FieldActionDesc) *FieldActionDesc {
	// target.takeFlatDamage(1)
	// report := fmt.Sprintf("%s inficted %d dmg to %s.\n", from.Name, 1, target.Name)
	// if target.IsDefeated() {
	// 	report += fmt.Sprintf("%s defeated %s.\n", from.Name, target.Name)
	// }
	return fad
}

func FireBoltAction(from *Hero, target *Hero, fad *FieldActionDesc) *FieldActionDesc {
	// target.takeFlatDamage(1)
	// report := fmt.Sprintf("%s inficted %d dmg to %s.\n", from.Name, 1, target.Name)
	// if target.IsDefeated() {
	// 	report += fmt.Sprintf("%s defeated %s.\n", from.Name, target.Name)
	// }
	return fad
}

func IceBoltAction(from *Hero, target *Hero, fad *FieldActionDesc) *FieldActionDesc {
	// target.takeFlatDamage(1)
	// report := fmt.Sprintf("%s inficted %d dmg to %s.\n", from.Name, 1, target.Name)
	// if target.IsDefeated() {
	// 	report += fmt.Sprintf("%s defeated %s.\n", from.Name, target.Name)
	// }
	return fad
}

func LightningBoltAction(from *Hero, target *Hero, fad *FieldActionDesc) *FieldActionDesc {
	// target.takeFlatDamage(1)
	// report := fmt.Sprintf("%s inficted %d dmg to %s.\n", from.Name, 1, target.Name)
	// if target.IsDefeated() {
	// 	report += fmt.Sprintf("%s defeated %s.\n", from.Name, target.Name)
	// }
	return fad
}

func (s Skill) UseActive(from *Hero, target *Hero) *FieldActionDesc {
	fad := &FieldActionDesc{
		FromH:          from,
		TargetH:        target,
		TargetPVChange: 0.0,
		FromPVChange:   0.0,
		UsedSkill:      &s,
	}
	switch s.Identifier {
	case WeaponAttack:
		return WeaponAttackAction(from, target, fad)
	case Spit:
		return SpitAction(from, target, fad)
	case Charge:
		return ChargeAction(from, target, fad)
	case Tackle:
		return TackleAction(from, target, fad)
	case Sweep:
		return SweepAction(from, target, fad)
	case FireBolt:
		return FireBoltAction(from, target, fad)
	case IceBolt:
		return IceBoltAction(from, target, fad)
	case LightningBolt:
		return LightningBoltAction(from, target, fad)
	}
	return fad
}
