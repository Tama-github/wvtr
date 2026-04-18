package expedition

import (
	"fmt"
	"time"
	"wvtrserv/databasemodel"
)

type TravelEvent struct {
	EEvent
}

func NewTravelEvent(duration time.Duration) *TravelEvent {
	return &TravelEvent{
		EEvent{
			duration: duration,
		},
	}
}

func (e TravelEvent) GetEventType() databasemodel.EncounterState {
	return databasemodel.Travel
}

func (e TravelEvent) Solve(startAt time.Time, t *databasemodel.Team) *databasemodel.ExpeditionStepResolveInfo {
	resolvInfo := fmt.Sprintf("T: {start: %s, end: %s}", startAt.String(), startAt.Add(e.duration).String())
	return &databasemodel.ExpeditionStepResolveInfo{
		StepInfos: resolvInfo,
		StepEndAt: startAt.Add(e.duration),
		StepState: e.GetEventType(),
	}
}
