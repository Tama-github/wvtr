package data

type DamageType int

const (
	Slash DamageType = iota
	Blunt
	Pierce
	Fire
	Frost
	Lightning
)

func CreateDamageMapRange(ranges map[DamageType]*StatsRange) *Damage {
	return &Damage{
		SlashDmg:     ranges[Slash].RollValue().Value,
		BluntDmg:     ranges[Blunt].RollValue().Value,
		PierceDmg:    ranges[Pierce].RollValue().Value,
		FireDmg:      ranges[Fire].RollValue().Value,
		FrostDmg:     ranges[Frost].RollValue().Value,
		LightningDmg: ranges[Lightning].RollValue().Value,
	}
}

func CreateDamageFromArray(arr []float64) *Damage {
	d := &Damage{}
	d.SlashDmg = arr[Slash]
	d.BluntDmg = arr[Blunt]
	d.PierceDmg = arr[Pierce]
	d.FireDmg = arr[Fire]
	d.FrostDmg = arr[Frost]
	d.LightningDmg = arr[Lightning]
	return d
}

func (d *Damage) GetDamageArray() []float64 {
	res := make([]float64, 6)
	res[Slash] = d.SlashDmg
	res[Blunt] = d.BluntDmg
	res[Pierce] = d.PierceDmg
	res[Fire] = d.FireDmg
	res[Frost] = d.FrostDmg
	res[Lightning] = d.LightningDmg

	return res
}

func (d *Damage) GetDmgFunc(f func(x float64) float64) *Damage {
	dmg := d.GetDamageArray()
	for i := range dmg {
		if dmg[i] > 0 {
			dmg[i] = f(dmg[i])
		}
	}
	return CreateDamageFromArray(dmg)
}

func (d *Damage) ApplyRes(hRes *Damage) *Damage {
	res := &Damage{}
	res.BluntDmg = d.BluntDmg * (1 - (hRes.BluntDmg / 100))
	res.PierceDmg = d.PierceDmg * (1 - (hRes.PierceDmg / 100))
	res.SlashDmg = d.SlashDmg * (1 - (hRes.SlashDmg / 100))
	res.FireDmg = d.FireDmg * (1 - (hRes.FireDmg / 100))
	res.FrostDmg = d.FrostDmg * (1 - (hRes.FrostDmg / 100))
	res.LightningDmg = d.LightningDmg * (1 - (hRes.LightningDmg / 100))

	return res
}

func (d *Damage) GetDamageWithBerserk() *Damage {
	res := &Damage{}
	res.BluntDmg = d.BluntDmg * (1.2)
	res.PierceDmg = d.PierceDmg * (1.2)
	res.SlashDmg = d.SlashDmg * (1.2)
	res.FireDmg = d.FireDmg
	res.FrostDmg = d.FrostDmg
	res.LightningDmg = d.LightningDmg

	return res
}
