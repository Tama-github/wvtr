package expedition

import (
	"wvtrserv/data"
)

type RewardPool struct {
	ItemBasePool []data.IEquipableBase
	CurrencyPool map[data.CurrencyType]data.StatsRange
}

type Reward struct {
	XP         float64
	Pool       *RewardPool
	LootChance float64                   /* Chance to roll one loot from the pool */
	Loot       []data.IStorable          /* Loot that have droped */
	Currencies map[data.CurrencyType]int /* [Currenty type] quantity */
}

func NewReward(rewardPool *RewardPool) *Reward {
	res := &Reward{
		XP:         0,
		Loot:       make([]data.IStorable, 0),
		Pool:       rewardPool,
		Currencies: make(map[data.CurrencyType]int, 0),
	}

	return res
}

func (r *Reward) MergeReward(toAdd *Reward) {
	r.AddXP(toAdd.XP)
	for _, s := range toAdd.Loot {
		r.AddStorable(s)
	}
	for k, v := range toAdd.Currencies {

		r.Currencies[k] += v
	}
}

func (r *Reward) AddXP(xp float64) {
	r.XP += xp
}

func (r *Reward) AddStorable(s data.IStorable) {
	r.Loot = append(r.Loot, s)
}

func (r *Reward) GenRandomReward() {
	if r.Pool == nil {
		return
	}
	if len(r.Pool.ItemBasePool) > 0 {
		rates := data.MakeUniformArrayRates(r.Pool.ItemBasePool)
		idx := data.RollInArrayWithRate(data.NaturalRoll(0, float64(len(rates))), rates)
		equipable := r.Pool.ItemBasePool[idx].GenEquipment()
		switch e := equipable.(type) {
		case *data.Weapon:
			r.AddStorable(e)
		case *data.Armor:
			r.AddStorable(e)
		case *data.Omamori:
			r.AddStorable(e)
		}
	}

	if len(r.Pool.CurrencyPool) > 0 {
		for k, v := range r.Pool.CurrencyPool {
			r.Currencies[k] = int(data.NaturalRoll(v.Min, v.Max))
		}
	}
}

func (r *Reward) GetCopy() *Reward {
	return &Reward{
		XP:         0,
		Pool:       r.Pool,
		LootChance: r.LootChance,
		Loot:       make([]data.IStorable, 0),
		Currencies: make(map[data.CurrencyType]int, 0),
	}
}
