package expedition

import (
	"time"
	"wvtrserv/data"
)

type HappeningType func(*data.Team, *data.ExpeditionStepResolveInfo)

/***********************/
/***  Neutral Event  ***/
/***********************/
type NeutralEvent struct {
	EEvent
	Happening HappeningType
}

func NewNeutralEvent(duration time.Duration, name string, h HappeningType) *NeutralEvent {
	return &NeutralEvent{
		EEvent: EEvent{
			duration: duration,
			name:     name,
		},
		Happening: h,
	}
}

func (e NeutralEvent) GetEventType() data.EncounterState {
	return data.Neutral
}

func (e NeutralEvent) Solve(startAt time.Time, t *data.Team) *data.ExpeditionStepResolveInfo {
	resExp := data.NewExpeditionResolveInfo(e.GetEventType())

	resExp.AddNewHappening(startAt, "Traveling Start")
	e.Happening(t, resExp)
	resExp.AddNewHappening(startAt.Add(e.duration), "Traveling End")

	return resExp
}
