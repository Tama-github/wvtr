package data

import (
	"fmt"
	"math"
	"time"
	"wvtrserv/logger"
)

type HeroTakeDamageStatus byte

const (
	TookDamage HeroTakeDamageStatus = 1 << iota
	Dodged
	Blocked
	Died
	SecondWindUsed
	Crit
)

func NewHero() *Hero {
	return &Hero{}
}

func (h *Hero) ClearAllStatusAndSetToFullLife() {
	h.Attributes.CurrentHP = h.Attributes.MaxHP
	// TODO: Clear special status if there are any
}

func (h *Hero) SetUSkillUsed(used bool) {
	h.UniqueSkill.HaveBeenUsed = used
}

func (h Hero) IsDefeated() bool {
	return h.Attributes.CurrentHP <= 0
}

func (h Hero) HasUniqueSkill(skillId SkillID) bool {
	return h.UniqueSkill.Identifier == skillId
}

func (h *Hero) GoToLevel(level int) {
	if level < h.Attributes.Level {
		return
	}

	for range h.Attributes.Level - level {
		h.GainXP(h.Attributes.XPToLvlUp - h.Attributes.CurrentXP)
	}
}

func (h *Hero) GainXP(amount float64) {
	if h.HasUniqueSkill(FastLearner) && h.UniqueSkill.Use(h) {
		amount = amount * (20.0 / 100.0)
	}

	for amount > 0 {
		thresholdForCurrentLevel := h.Attributes.LevelThreshold()
		xpToGainForLevel := amount

		if amount+h.Attributes.CurrentXP >= thresholdForCurrentLevel {
			xpToGainForLevel = thresholdForCurrentLevel
			h.LevelUp()
		}

		h.Attributes.CurrentXP += xpToGainForLevel
		amount -= xpToGainForLevel
	}
}

func (h *Hero) IncreaseAttributeWithRate() {
	attrs := h.Attributes.GetAttributesArray()
	grs := h.Attributes.GetGRArray()
	for i := range len(grs) {
		// this is in case the proba is above 100%
		// we add the int part and rand on the float part
		gr := grs[i]
		if h.HasUniqueSkill(Prodigy) && h.UniqueSkill.Use(h) {
			gr += (attrs[IntelligenceID] + attrs[LuckID]) / 200
		}

		toadd := float64(int(gr))
		proba := gr - toadd
		if RollCheck(NaturalRoll(0, 1), 1-proba) {
			toadd++
		}
		attrs[i] += toadd
	}
	h.Attributes.SetAttributesWithArray(attrs)
}

func (h *Hero) LevelUp() {
	h.Attributes.Level += 1
	h.Attributes.CurrentXP = 0
	h.Attributes.XPToLvlUp = h.Attributes.LevelThreshold()
	h.IncreaseAttributeWithRate()
}

func GenerateGrowthRateFromRank(rank string) []float64 {
	res := make([]float64, 5)
	min := 0.0
	max := 0.0
	switch rank {
	case "S":
		min = 0.1
		max = 0.3
	case "A":
		min = 0.05
		max = 0.2
	case "B":
		min = 0.05
		max = 0.15
	case "C":
		min = 0
		max = 0.1
	case "D":
		min = 0
		max = 0.05
	}
	res[HPgtID] = NaturalRoll(min, max)
	res[SgtID] = NaturalRoll(min, max)
	res[IgtID] = NaturalRoll(min, max)
	res[DgtID] = NaturalRoll(min, max)
	res[LgtID] = NaturalRoll(min, max)
	return res
}

func (h *Hero) RollNumber(min float64, max float64) float64 {
	resRoll := NaturalRoll(min, max)

	// Lucky
	if h.HasUniqueSkill(Lucky) && h.UniqueSkill.Use(h) {
		skillRoll := NaturalRoll(min, max)
		if resRoll < skillRoll {
			resRoll = skillRoll
		}
	}
	return resRoll
}

func (h *Hero) RollCheck(proba float64) bool {
	return RollCheck(h.RollNumber(0, 1), proba)
}

func (h *Hero) ChooseTarget(t *Team) *Hero {
	rate := make([]float64, len(t.Heroes))
	for i, h := range t.Heroes {
		if !h.IsDefeated() {
			rate[i] = 1
		} else {
			rate[i] = 0
		}
	}
	return t.Heroes[RollInArrayWithRate(NaturalRoll(0, 1), rate)]
}

func (h *Hero) ChooseAction(friends *Team, enemies *Team) (*Skill, *Hero) {
	action := h.WeaponAttack

	if h.ActiveSkill != nil && RollCheck(NaturalRoll(0, 1), 0.5) {
		action = h.ActiveSkill
	}
	if action == nil {
		return nil, nil
	}
	var target *Hero = nil
	switch action.Targeting {
	case Self:
		target = h
	case Enemy:
		target = h.ChooseTarget(enemies)
	case Friends:
		target = h.ChooseTarget(friends)
	}
	return action, target
}

func (h *Hero) Play(when time.Time, friends *Team, enemies *Team, fightReport *ExpeditionStepResolveInfo) time.Duration {
	what, target := h.ChooseAction(friends, enemies)
	if what == nil || target == nil {
		return 30 * time.Second
	}
	report := fmt.Sprintf("%s use %s on %s.", h.Name, what.Name, target.Name)

	fad := what.UseActive(h, target)
	logger.DumpLog.Println(fad.String())
	fightReport.AddNewHappening(when, report, fad)
	res := what.RecuperationDuration
	if h.HasUniqueSkill(Trickster) && h.UniqueSkill.Use(h) {
		res -= res * time.Duration(math.Min(float64(h.Attributes.GetDexterity()/100), 0.5)) * time.Second
	}
	return res
}

func (h *Hero) takeFlatDamage(dmg float64) {
	h.Attributes.CurrentHP -= dmg
}

func (h *Hero) Dodge(from *Hero) bool {
	// TODO
	return false
}

func (h *Hero) Block(from *Hero) bool {
	// TODO
	return false
}

func (h *Hero) IsCrit() bool {
	return h.RollCheck(h.Equipment.Weapon.BaseCritRate.Value)
}

func (h *Hero) Rest(heal float64) float64 {
	actualHeal := heal
	if h.HasUniqueSkill(GoodRest) {
		actualHeal = heal * 1.33
	}
	if h.Attributes.CurrentHP+actualHeal >= h.Attributes.MaxHP {
		actualHeal = h.Attributes.MaxHP - h.Attributes.CurrentHP
	}
	h.Attributes.CurrentHP += actualHeal
	return actualHeal
}

func (h *Hero) TakeDamage(dmg *Damage, takeFrom *Hero, fad *FieldActionDesc) *FieldActionDesc {
	if takeFrom.HasUniqueSkill(Berserk) && takeFrom.UniqueSkill.Use(takeFrom) {
		dmg = dmg.GetDamageWithBerserk()
	}

	// check dodge
	if h.Dodge(takeFrom) {
		fad.TargetStatus |= Dodged
		return fad
	}

	// check blocked
	if h.Block(takeFrom) {
		fad.TargetStatus |= Blocked
		return fad
	}

	// check resistances
	actualDamage := dmg.ApplyRes(h.GetTotalRes())
	dmgSum := 0.0
	// get total damage taken
	for _, d := range actualDamage.GetDamageArray() {
		dmgSum += d
	}

	h.takeFlatDamage(dmgSum)
	fad.TargetStatus |= TookDamage
	fad.TargetPVChange = dmgSum

	// send reflected damage if there are any

	// check if dead
	if h.IsDefeated() {
		if h.HasUniqueSkill(SecondWind) && h.UniqueSkill.Use(h) {
			h.ClearAllStatusAndSetToFullLife()
			fad.TargetStatus |= SecondWindUsed
		} else {
			fad.TargetStatus |= Died
			return fad
		}
	}

	return fad
}

// return reation time at the bagining of the combat
func (h *Hero) InitiativeRoll() time.Duration {
	dex := h.Attributes.Dexterity

	initiativeRoll := h.RollNumber(0.5, dex)

	secInitiative := time.Duration((5.0/initiativeRoll)+2.0) * time.Second
	resInit := secInitiative + time.Duration(NaturalRoll(-100000, 100000))*time.Nanosecond
	if h.HasUniqueSkill(Trickster) && h.UniqueSkill.Use(h) {
		secInitiative -= secInitiative * time.Duration(math.Min(float64(dex/100), 0.5)) * time.Second
	}
	return resInit
}

func (h *Hero) GetAttackDamageForWeapon() *Damage {
	// Basicaly a fist
	if h.Equipment.Weapon == nil {
		return &Damage{
			BluntDmg: 1,
		}
	}
	w := h.Equipment.Weapon
	attr := h.Attributes
	// we apply attribute scaling to all damage types
	dmg := w.BaseDamage.GetDmgFunc(func(x float64) float64 {
		return x + w.StrScaling*attr.Strength + w.IntScaling*attr.Intelligence + w.DexScaling*attr.Dexterity + w.LckScaling*attr.Luck
	})

	return dmg
}

func (h *Hero) GetTotalCritMulti() float64 {
	return 1.5
}

func (h *Hero) GetBaseStr() float64 {
	return h.Attributes.Strength
}

func (h *Hero) GetBaseInt() float64 {
	return h.Attributes.Strength
}

func (h *Hero) GetBaseDex() float64 {
	return h.Attributes.Strength
}

func (h *Hero) GetBaseLck() float64 {
	return h.Attributes.Strength
}

func (h *Hero) GetBaseRes() *Damage {
	return CreateDamageFromArray(h.Attributes.GetResArray())
}

func (h *Hero) GetTotalStr() float64 {
	res := h.GetBaseStr()

	// Affixes modifier
	res += h.Equipment.GetTotalValueOfAffixInEquipment(FlatStr).Ranges[0].Value

	// Affixes Multiplicator
	res *= h.Equipment.GetTotalValueOfAffixInEquipment(PercentStr).Ranges[0].Value

	return res
}

func (h *Hero) GetTotalDex() float64 {
	res := h.GetBaseDex()

	// Affixes modifier
	res += h.Equipment.GetTotalValueOfAffixInEquipment(FlatDex).Ranges[0].Value

	// Affixes Multiplicator
	res *= h.Equipment.GetTotalValueOfAffixInEquipment(PercentDex).Ranges[0].Value

	return res
}

func (h *Hero) GetTotalInt() float64 {
	res := h.GetBaseInt()

	// Affixes modifier
	res += h.Equipment.GetTotalValueOfAffixInEquipment(FlatInt).Ranges[0].Value

	// Affixes Multiplicator
	res *= h.Equipment.GetTotalValueOfAffixInEquipment(PercentInt).Ranges[0].Value

	return res
}

func (h *Hero) GetTotalLck() float64 {
	res := h.GetBaseLck()

	// Affixes modifier
	res += h.Equipment.GetTotalValueOfAffixInEquipment(FlatLck).Ranges[0].Value

	// Affixes Multiplicator
	res *= h.Equipment.GetTotalValueOfAffixInEquipment(PercentLck).Ranges[0].Value

	return res
}

func (h *Hero) GetTotalRes() *Damage {
	res := h.GetBaseRes()

	// Affixes modifier
	physRes := h.Equipment.GetTotalValueOfAffixInEquipment(PhysRes).Ranges[0].Value

	res.SlashDmg = h.Equipment.GetTotalValueOfAffixInEquipment(SlashRes).Ranges[0].Value + res.SlashDmg + physRes
	res.BluntDmg = h.Equipment.GetTotalValueOfAffixInEquipment(BluntRes).Ranges[0].Value + res.BluntDmg + physRes
	res.PierceDmg = h.Equipment.GetTotalValueOfAffixInEquipment(PierceRes).Ranges[0].Value + res.PierceDmg + physRes

	elemRes := h.Equipment.GetTotalValueOfAffixInEquipment(ElemRes).Ranges[0].Value

	res.FireDmg = h.Equipment.GetTotalValueOfAffixInEquipment(FireRes).Ranges[0].Value + res.FireDmg + elemRes
	res.FrostDmg = h.Equipment.GetTotalValueOfAffixInEquipment(FrostRes).Ranges[0].Value + res.FrostDmg + elemRes
	res.LightningDmg = h.Equipment.GetTotalValueOfAffixInEquipment(LightningRes).Ranges[0].Value + res.LightningDmg + elemRes

	return res
}
