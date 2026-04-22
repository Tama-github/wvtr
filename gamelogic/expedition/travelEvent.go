package expedition

import (
	"time"
	"wvtrserv/data"
)

type TravelEvent struct {
	EEvent
}

func NewTravelEvent(duration time.Duration, name string) *TravelEvent {
	return &TravelEvent{
		EEvent{
			duration: duration,
			name:     name,
		},
	}
}

func (e TravelEvent) GetEventType() data.EncounterState {
	return data.Travel
}

func (e TravelEvent) Solve(startAt time.Time, t *data.Team) *data.ExpeditionStepResolveInfo {
	resExp := data.NewExpeditionResolveInfo(e.GetEventType())

	resExp.AddNewHappening(startAt, "Traveling Start")
	resExp.AddNewHappening(startAt.Add(e.duration), "Traveling End")

	return resExp
}
