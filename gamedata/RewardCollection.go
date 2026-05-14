package gamedata

import (
	"wvtrserv/data"
	"wvtrserv/gamelogic/expedition"
)

var travelReward1 *expedition.RewardPool = &expedition.RewardPool{
	ItemBasePool: []data.IEquipableBase{},
	CurrencyPool: map[data.CurrencyType]data.StatsRange{
		data.Gold:   {Min: 0, Max: 1.01},
		data.CScrap: {Min: 0, Max: 1.05},
		data.LSCrap: {Min: 0, Max: 1.05},
		data.MScrap: {Min: 0, Max: 1.05},
	},
}

var plainsRewardPool *expedition.RewardPool = &expedition.RewardPool{
	ItemBasePool: []data.IEquipableBase{
		SwordBase,
		DaggerBase,
		HammerBase,
	},
	CurrencyPool: map[data.CurrencyType]data.StatsRange{
		data.Gold:   {Min: 0, Max: 0},
		data.CScrap: {Min: 0, Max: 1.2},
		data.LSCrap: {Min: 0, Max: 1.1},
		data.MScrap: {Min: 0, Max: 1.05},
	},
}

var craftRewardPool *expedition.RewardPool = &expedition.RewardPool{
	ItemBasePool: []data.IEquipableBase{
		SwordBase,
		DaggerBase,
		HammerBase,
	},
	CurrencyPool: map[data.CurrencyType]data.StatsRange{
		data.Gold:   {Min: 0, Max: 0},
		data.CScrap: {Min: 0, Max: 0},
		data.LSCrap: {Min: 0, Max: 0},
		data.MScrap: {Min: 0, Max: 0},
	},
}
