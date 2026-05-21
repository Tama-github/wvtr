package gamedata

import "wvtrserv/data"

var lifeForceCharm *data.OmamoriBase = &data.OmamoriBase{
	EquipableBase: data.EquipableBase{
		BaseName: "Life force charm",
		AffixesPool: []*data.Affix{
			FlatLife[1],
			PercentLife[1],
		},
	},
}

var statsCham *data.OmamoriBase = &data.OmamoriBase{
	EquipableBase: data.EquipableBase{
		BaseName: "Stat charm",
		AffixesPool: []*data.Affix{
			FlatStr[1],
			FlatInt[1],
			FlatDex[1],
			FlatLck[1],
		},
	},
}

var elemProtectionCham *data.OmamoriBase = &data.OmamoriBase{
	EquipableBase: data.EquipableBase{
		BaseName: "Elemental Protection charm",
		AffixesPool: []*data.Affix{
			SlashRes[1],
			BluntRes[1],
			PierceRes[1],
		},
	},
}

var physProtectionCham *data.OmamoriBase = &data.OmamoriBase{
	EquipableBase: data.EquipableBase{
		BaseName: "Physical Protection charm",
		AffixesPool: []*data.Affix{
			FireRes[1],
			FrostRes[1],
			LightningRes[1],
		},
	},
}
