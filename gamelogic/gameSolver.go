package gamelogic

// func (g *databasemanager.GameState) ResolveGameState(t *time.Time) *ExpeditionStepResolveInfo {
// 	idx := -1
// 	var currentExpStep *ExpeditionStepResolveInfo = nil
// 	if g.CurrentExpedition != nil {
// 		idx, currentExpStep = g.CurrentExpedition.GetCurrentStep(t)
// 	}
// 	if idx < 0 {
// 		g.State = Home
// 		g.CurrentExpedition = nil
// 		g.ETeam = nil
// 		return nil
// 	}

// 	g.State = currentExpStep.StepState
// 	if g.State == Fight {
// 		g.ETeam = Expeditions[g.CurrentExpedition.Identifier].GetEnemyTeamForEvent(idx)
// 	}
// 	return currentExpStep
// }

// func (g *GameState) LaunchExpedition(expIdentifier string, pTeam *Team) {
// 	newExpedition, ok := Expeditions[expIdentifier]
// 	if ok {
// 		g.CurrentExpedition = newExpedition.Solve(expIdentifier, pTeam)
// 		CreateExpeditionDB(g.CurrentExpedition)
// 		g.State = g.CurrentExpedition.WhatHappened[0].StepState
// 		UpdateGameState(g)
// 	} else {
// 		fmt.Printf("[%s] is not an existing expedition.\n", expIdentifier)
// 	}
// }

// func (e ExpeditionDB) GetCurrentStep(t *time.Time) (int, *ExpeditionStepResolveInfo) {
// 	for i, step := range e.WhatHappened {
// 		if step.StepEndAt.After(*t) {
// 			return i, step
// 		}
// 	}
// 	return -1, nil
// }

// func (u User) isHome() bool {
// 	return u.State.State != Error && u.State.State == Home
// }

// func (u *User) LaunchExpedition(expIdentifier string) {
// 	if u.isHome() {
// 		u.State.LaunchExpedition(expIdentifier, u.CurrentTeam)
// 	}
// }

// func (u *User) ResolveExpeditionState(t *time.Time) *ExpeditionStepResolveInfo {
// 	res := u.State.ResolveGameState(t)

// 	return res
// }
