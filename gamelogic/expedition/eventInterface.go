package expedition

import (
	"time"
	"wvtrserv/databasemodel"
)

type ExpeditionEvent interface {
	EndAt(startAt time.Time) time.Time
	GetEventType() databasemodel.EncounterState
	Solve(startAt time.Time, t *databasemodel.Team) *databasemodel.ExpeditionStepResolveInfo
	GetDuration() time.Duration
}

type EEvent struct {
	duration   time.Duration
	solveTrace string
}

func (e EEvent) EndAt(startAt time.Time) time.Time {
	return startAt.Add(e.duration)
}

func (e EEvent) GetDuration() time.Duration {
	return e.duration
}
