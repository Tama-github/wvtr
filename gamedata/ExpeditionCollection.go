package gamedata

import (
	"slices"
	"wvtrserv/data"
	"wvtrserv/gamelogic/expedition"
)

var ExpeditionsJobs = map[string]expedition.Expedition{
	"Work at the tavern": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			workGold,
		},
		Order: 0,
	},
}

var ExpeditionsHeal = map[string]expedition.Expedition{
	"Heal in a hospital": {
		ImgURL:     DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Cost:       allCurrencies[data.Gold],
		CostNumber: 500,
		Events: []expedition.ExpeditionEvent{
			hospitalRest,
		},
		Order: 2,
	},
	"Heal at a healer's hut": {
		ImgURL:     DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Cost:       allCurrencies[data.Gold],
		CostNumber: 100,
		Events: []expedition.ExpeditionEvent{
			healerRest,
		},
		Order: 1,
	},
	"Rest": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			rest,
		},
		Order: 0,
	},
}

var ExpeditionsQuests = map[string]expedition.Expedition{
	"Plain quest": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			goingToPlains,
			plainFight,
			goingToPlains,
		},
		Order: 0,
	},
}

var ExpeditionsCrafts = map[string]expedition.Expedition{
	"Plain quest": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			goingToPlains,
			plainFight,
			goingToPlains,
		},
		Order: 0,
	},
}

var ExpeditionsTest = map[string]expedition.Expedition{
	"Traveling 10 sec": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			traval10s,
		},
		Order: 0,
	},
	"Work": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			workShortTest,
		},
		Order: 6,
	},
	"Training": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/self_training.png",
		Events: []expedition.ExpeditionEvent{
			selfTrainingTest,
		},
		Order: 1,
	},
	"Craft Weapon": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			craftWeapon,
		},
		Order: 2,
	},
	"Scrap Weapon": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			craftWeapon,
		},
		Cost:       &data.Weapon{Equipable: data.Equipable{Storable: data.Storable{Name: "Weapon"}}},
		CostNumber: 1,
		Order:      3,
	},
	"Scrap Armor": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			craftWeapon,
		},
		Cost:       &data.Armor{Equipable: data.Equipable{Storable: data.Storable{Name: "Armor"}}},
		CostNumber: 1,
		Order:      4,
	},
	"Scrap Omamori": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			craftWeapon,
		},
		Cost:       &data.Omamori{Equipable: data.Equipable{Storable: data.Storable{Name: "Omamori"}}},
		CostNumber: 1,
		Order:      5,
	},
}

var Expeditions = map[string]map[string]expedition.Expedition{
	"Jobs":   ExpeditionsJobs,
	"Heal":   ExpeditionsHeal,
	"Quests": ExpeditionsQuests,
	"Tests":  ExpeditionsTest,
}

func GetAvailableExpeditions(user *data.User) []*expedition.ExpToSendToFront {
	res := make([]*expedition.ExpToSendToFront, 0)

	for category, exps := range Expeditions {
		for k, v := range exps {
			name := ""
			cbl := true
			if v.Cost != nil {
				name = v.Cost.GetName()
				cbl = user.Inventory.HasEnought(v.Cost, v.CostNumber)
			}
			res = append(res, &expedition.ExpToSendToFront{
				Category:      category,
				Key:           k,
				ImgURL:        v.ImgURL,
				Duration:      v.GetMinimumTotalTime(),
				CostName:      name,
				CostNumber:    v.CostNumber,
				CanBeLaunched: cbl,
				Order:         v.Order,
			})
		}
	}
	// order by Order value
	slices.SortFunc(res, func(a, b *expedition.ExpToSendToFront) int {
		if a.Order < b.Order {
			return -1
		}
		return 1
	})
	return res
}
