package gameexpedition

import (
	"time"
	"wvtrserv/databasemodel"
	"wvtrserv/gamelogic/expedition"
	"wvtrserv/logger"
)

var traval30s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second * 30)
var traval40s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second * 40)

var Expeditions = map[string]expedition.Expedition{
	"travel30s": {
		Events: []expedition.ExpeditionEvent{
			traval30s,
		},
	},
	"travel40s": {
		Events: []expedition.ExpeditionEvent{
			traval40s,
		},
	},
}

func GetAvailableExpeditions() map[string]time.Duration {
	res := make(map[string]time.Duration)
	for k, v := range Expeditions {
		res[k] = v.GetMinimumTotalTime()
	}
	return res
}

func GetEnemyTeamForEvent(identifier string, idx int) *databasemodel.Team {
	fEvent, ok := Expeditions[identifier].Events[idx].(expedition.FightEvent)
	if !ok {
		logger.ErrLog.Printf("can't cast %dth event from %s expedition event into a fight event.", idx, identifier)
		return nil
	}
	return fEvent.ETeam
}

// func ResolveGameState(t *time.Time, g *databasemodel.GameState, exp *databasemodel.ExpeditionDB) *databasemodel.ExpeditionStepResolveInfo {
// 	idx := -1
// 	var currentExpStep *databasemodel.ExpeditionStepResolveInfo = nil
// 	if g.CurrentExpedition != nil {
// 		idx, currentExpStep = g.CurrentExpedition.GetCurrentStep(t)
// 	}
// 	if idx < 0 {
// 		g.State = databasemodel.Home
// 		g.CurrentExpedition = nil
// 		g.ETeam = nil
// 		return nil
// 	}

// 	g.State = currentExpStep.StepState
// 	if g.State == databasemodel.Fight {
// 		event := Expeditions[g.CurrentExpedition.Identifier].Events[idx]
// 		fEvent := event.(expedition.FightEvent)
// 		g.ETeam = fEvent.ETeam
// 	}
// 	return currentExpStep
// }
