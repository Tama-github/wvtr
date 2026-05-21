package expedition

import (
	"time"
	"wvtrserv/data"
)

type HappeningType func(ExpeditionEvent, *data.Team, *data.ExpeditionStepResolveInfo, []data.IStorable)

/***********************/
/***  Neutral Event  ***/
/***********************/
type NeutralEvent struct {
	EEvent
	Happening HappeningType
}

func NewNeutralEvent(duration time.Duration, name string, rewardPool *RewardPool, h HappeningType) *NeutralEvent {
	return &NeutralEvent{
		EEvent: EEvent{
			duration: duration,
			Name:     name,
			Reward:   NewReward(rewardPool),
		},
		Happening: h,
	}
}

func (e NeutralEvent) GetEventType() data.EncounterState {
	return data.Neutral
}

func (e *NeutralEvent) Solve(startAt time.Time, t *data.Team, toSpend []data.IStorable) *data.ExpeditionStepResolveInfo {
	resExp := data.NewExpeditionResolveInfo(e.GetEventType())

	resExp.AddNewHappening(startAt, "Traveling Start", nil)
	e.Happening(e, t, resExp, toSpend)
	resExp.AddNewHappening(startAt.Add(e.duration), "Traveling End", nil)
	e.Reward.GenRandomReward()
	return resExp
}

func (e NeutralEvent) CopyEvent() ExpeditionEvent {
	return &NeutralEvent{
		EEvent: EEvent{
			duration:         e.duration,
			EventSolvingInfo: &data.ExpeditionStepResolveInfo{},
			Name:             e.Name,
			Reward:           e.Reward.GetCopy(),
		},
		Happening: e.Happening,
	}
}
