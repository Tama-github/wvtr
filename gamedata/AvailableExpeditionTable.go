package gamedata

import (
	"time"
	"wvtrserv/data"
	"wvtrserv/gamelogic/expedition"
	"wvtrserv/logger"
)

var traval30s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*30, "Travel")
var traval40s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*40, "Travel")
var traval10s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*10, "Travel")
var nothing10s expedition.ExpeditionEvent = expedition.NewNeutralEvent(time.Second*10, "Neutral", expedition.HappeningType(func(t *data.Team, e *data.ExpeditionStepResolveInfo) {}))

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
	"travelAndDoNothing": {
		Events: []expedition.ExpeditionEvent{
			traval10s,
			nothing10s,
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

func GetEnemyTeamForEvent(identifier string, idx int) *data.Team {
	fEvent, ok := Expeditions[identifier].Events[idx].(expedition.FightEvent)
	if !ok {
		logger.ErrLog.Printf("can't cast %dth event from %s expedition event into a fight event.", idx, identifier)
		return nil
	}
	return fEvent.ETeam
}
