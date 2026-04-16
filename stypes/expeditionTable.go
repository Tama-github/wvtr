package stypes

import (
	"time"
)

var traval30s ExpeditionEvent = NewTravelEvent(time.Second * 30)
var traval40s ExpeditionEvent = NewTravelEvent(time.Second * 40)

var Expeditions = map[string]Expedition{
	"travel30s": {
		Events: []ExpeditionEvent{
			traval30s,
		},
	},
	"travel40s": {
		Events: []ExpeditionEvent{
			traval40s,
		},
	},
}

func GetExpeditions() map[string]time.Duration {
	res := make(map[string]time.Duration)
	for k, v := range Expeditions {
		res[k] = v.GetMinimumTotalTime()
	}
	return res
}
