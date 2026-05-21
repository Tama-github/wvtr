package expedition

import (
	"time"
	"wvtrserv/data"
)

type ExpeditionEvent interface {
	EndAt(startAt time.Time) time.Time
	GetEventType() data.EncounterState
	Solve(startAt time.Time, t *data.Team, toSpend []data.IStorable) *data.ExpeditionStepResolveInfo
	GetDuration() time.Duration
	GetName() string
	CopyEvent() ExpeditionEvent
	GetReward() *Reward
}

type EEvent struct {
	duration         time.Duration
	EventSolvingInfo *data.ExpeditionStepResolveInfo
	Name             string
	Reward           *Reward
}

func (e EEvent) EndAt(startAt time.Time) time.Time {
	return startAt.Add(e.duration)
}

func (e EEvent) GetDuration() time.Duration {
	return e.duration
}

func (e EEvent) GetName() string {
	return e.Name
}

func (e EEvent) GetReward() *Reward {
	return e.Reward
}
