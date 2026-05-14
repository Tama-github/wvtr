package databasecontroller

import (
	"time"
	"wvtrserv/data"
)

func CreateExpeditionDB(edb *data.ExpeditionDB) {
	CreateInventory(edb.ExpeditionRewards.Loot)
	db.Create(edb)
}

func GetExpeditionDBByID(id uint) *data.ExpeditionDB {
	var ex *data.ExpeditionDB = &data.ExpeditionDB{}
	db.Preload("WhatHappened").
		Preload("ExpeditionRewards").
		Find(&ex, id)

	for i := range ex.WhatHappened {
		ex.WhatHappened[i] = GetExpeditionStepResolveInfoByID(ex.WhatHappened[i].ID)
	}

	if ex.ExpeditionRewards != nil {
		ex.ExpeditionRewards = GetGetRewardByID(ex.ExpeditionRewards.ID)
	}
	return ex
}

func GetCurrentExpedition(exp *data.ExpeditionDB) *data.ExpeditionDB {
	db.Preload("WhatHappened").
		Find(&exp)
	return exp
}

func GetCurrentExpeditionStepIdx(e data.ExpeditionDB, t *time.Time) int {
	for i, step := range e.WhatHappened {
		if step.Timeline[len(step.Timeline)-1].When.After(*t) {
			return i
		}
	}
	return len(e.WhatHappened)
}

func DeleteExpeditionDB(exp *data.ExpeditionDB) {
	for _, eri := range exp.WhatHappened {
		db.Delete(eri)
	}
	db.Delete(exp.ExpeditionRewards)
	db.Delete(exp)
}
