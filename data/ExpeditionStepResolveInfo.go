package data

import "time"

func NewExpeditionResolveInfo(state EncounterState) *ExpeditionStepResolveInfo {
	return &ExpeditionStepResolveInfo{
		StepState: state,
		Timeline:  make([]*ExpeditionStepTimestamp, 0),
	}
}

func (e *ExpeditionStepResolveInfo) AddNewHappening(when time.Time, what string) {
	e.Timeline = append(e.Timeline, &ExpeditionStepTimestamp{When: when, What: what})
}
