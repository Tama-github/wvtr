package databasecontroller

import (
	"time"
	"wvtrserv/databasemodel"
)

func GetCurrentExpeditionStepIdx(e databasemodel.ExpeditionDB, t *time.Time) int {
	for i, step := range e.WhatHappened {
		if step.StepEndAt.After(*t) {
			return i
		}
	}
	return -1
}
