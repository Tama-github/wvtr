package gamedata

import "wvtrserv/data"

var armorsImgsPath string = DOMAIN_NAME + "/imgs/equipments/armors/"

var PlateArmorBase *data.ArmorBase = &data.ArmorBase{
	BaseBlockScore: &data.StatsRange{Min: 5.0, Max: 10.0},
	BaseEvadeScore: &data.StatsRange{Min: 0.0, Max: 0.0},
	BaseResistancesRange: map[data.DamageType]*data.StatsRange{
		data.Slash:     {Min: 1.0, Max: 5.0},
		data.Blunt:     {Min: 1.0, Max: 5.0},
		data.Pierce:    {Min: 1.0, Max: 5.0},
		data.Fire:      {Min: 0.0, Max: 0.0},
		data.Frost:     {Min: 0.0, Max: 0.0},
		data.Lightning: {Min: 0.0, Max: 0.0},
	},
	EquipableBase: data.EquipableBase{
		BaseName: "Plate armor",
		IconURL:  armorsImgsPath + "plate_armor_icon.png",
		AffixesPool: []*data.Affix{
			PhysicalRes[1],
			ElementalRes[1],
			FlatLife[1],
			PercentLife[1],
			FlatDex[1],
			FlatInt[1],
			FlatStr[1],
			FlatLck[1],
		},
	},
}

var LeatherArmorBase *data.ArmorBase = &data.ArmorBase{
	BaseBlockScore: &data.StatsRange{Min: 0.0, Max: 0.0},
	BaseEvadeScore: &data.StatsRange{Min: 5.0, Max: 10.0},
	BaseResistancesRange: map[data.DamageType]*data.StatsRange{
		data.Slash:     {Min: 1.0, Max: 5.0},
		data.Blunt:     {Min: 1.0, Max: 5.0},
		data.Pierce:    {Min: 1.0, Max: 5.0},
		data.Fire:      {Min: 0.0, Max: 0.0},
		data.Frost:     {Min: 0.0, Max: 0.0},
		data.Lightning: {Min: 0.0, Max: 0.0},
	},
	EquipableBase: data.EquipableBase{
		BaseName: "Leather armor",
		IconURL:  armorsImgsPath + "leather_armor_icon.png",
		AffixesPool: []*data.Affix{
			PhysicalRes[1],
			ElementalRes[1],
			FlatLife[1],
			PercentLife[1],
			FlatDex[1],
			FlatInt[1],
			FlatStr[1],
			FlatLck[1],
		},
	},
}

var ClothArmorBase *data.ArmorBase = &data.ArmorBase{
	BaseBlockScore: &data.StatsRange{Min: 0.0, Max: 0.0},
	BaseEvadeScore: &data.StatsRange{Min: 1.0, Max: 5.0},
	BaseResistancesRange: map[data.DamageType]*data.StatsRange{
		data.Slash:     {Min: 0.0, Max: 0.0},
		data.Blunt:     {Min: 0.0, Max: 0.0},
		data.Pierce:    {Min: 0.0, Max: 0.0},
		data.Fire:      {Min: 1.0, Max: 5.0},
		data.Frost:     {Min: 1.0, Max: 5.0},
		data.Lightning: {Min: 1.0, Max: 5.0},
	},
	EquipableBase: data.EquipableBase{
		BaseName: "Cloth armor",
		IconURL:  armorsImgsPath + "cloth_armor_icon.png",
		AffixesPool: []*data.Affix{
			PhysicalRes[1],
			ElementalRes[1],
			FlatLife[1],
			PercentLife[1],
			FlatDex[1],
			FlatInt[1],
			FlatStr[1],
			FlatLck[1],
		},
	},
}
