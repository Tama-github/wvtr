package databasecontroller

import (
	"time"
	"wvtrserv/data"
)

func CreateExpeditionDB(edb *data.ExpeditionDB) {
	db.Save(edb)
}

func GetCurrentExpeditionStepIdx(e data.ExpeditionDB, t *time.Time) int {
	for i, step := range e.WhatHappened {
		if step.Timeline[len(step.Timeline)-1].When.After(*t) {
			return i
		}
	}
	return -1
}
