package expedition

import (
	"fmt"
	"time"
	"wvtrserv/databasemodel"
)

type HappeningType func(*databasemodel.Team) string

/***********************/
/***  Neutral Event  ***/
/***********************/
type NeutralEvent struct {
	EEvent
	Happening HappeningType
}

func NewNeutralEvent(duration time.Duration, h HappeningType) *NeutralEvent {
	return &NeutralEvent{
		EEvent: EEvent{
			duration: duration,
		},
		Happening: h,
	}
}

func (e NeutralEvent) GetEventType() databasemodel.EncounterState {
	return databasemodel.Neutral
}

func (e NeutralEvent) Solve(startAt time.Time, t *databasemodel.Team) *databasemodel.ExpeditionStepResolveInfo {

	resolvInfo := fmt.Sprintf("T: {start: %s, end: %s, %s}", startAt.String(), startAt.Add(e.duration).String(), e.Happening(t))
	fmt.Printf("Solve neutral event : %s\n", resolvInfo)
	return &databasemodel.ExpeditionStepResolveInfo{
		StepInfos: resolvInfo,
		StepEndAt: startAt.Add(e.duration),
		StepState: e.GetEventType(),
	}
}
