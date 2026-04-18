package databasecontroller

import (
	"time"
	"wvtrserv/databasemodel"
	"wvtrserv/gameexpedition"
)

func UpdateGameStateWithTime(g *databasemodel.GameState, t *time.Time) *databasemodel.ExpeditionStepResolveInfo {
	idx := -1
	var currentExpStep *databasemodel.ExpeditionStepResolveInfo = nil
	if g.CurrentExpedition != nil {
		idx = GetCurrentExpeditionStepIdx(*g.CurrentExpedition, t)
		currentExpStep = UpdateGameStateWithIndex(g, idx)
	}
	return currentExpStep
}

func UpdateGameStateWithIndex(g *databasemodel.GameState, idx int) *databasemodel.ExpeditionStepResolveInfo {
	var currentExpStep *databasemodel.ExpeditionStepResolveInfo = nil

	if idx < 0 || idx >= len(g.CurrentExpedition.WhatHappened) {
		g.State = databasemodel.Home
		g.CurrentExpedition = nil
		g.ETeam = nil
		return nil
	}

	currentExpStep = g.CurrentExpedition.WhatHappened[idx]

	g.State = currentExpStep.StepState
	if g.State == databasemodel.Fight {
		g.ETeam = gameexpedition.GetEnemyTeamForEvent(g.CurrentExpedition.Identifier, idx)
	}
	return currentExpStep
}
