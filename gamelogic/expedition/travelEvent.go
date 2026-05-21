package expedition

import (
	"time"
	"wvtrserv/data"
)

type TravelEvent struct {
	EEvent
}

func NewTravelEvent(duration time.Duration, name string, rewardPool *RewardPool) *TravelEvent {
	return &TravelEvent{
		EEvent{
			duration: duration,
			Name:     name,
			Reward:   NewReward(rewardPool),
		},
	}
}

func (e TravelEvent) GetEventType() data.EncounterState {
	return data.Travel
}

func (e TravelEvent) Solve(startAt time.Time, t *data.Team, toSpend []data.IStorable) *data.ExpeditionStepResolveInfo {
	resExp := data.NewExpeditionResolveInfo(e.GetEventType())

	resExp.AddNewHappening(startAt, "Traveling Start", nil)
	resExp.AddNewHappening(startAt.Add(e.duration), "Traveling End", nil)
	e.Reward.GenRandomReward()
	return resExp
}

func (e TravelEvent) CopyEvent() ExpeditionEvent {
	return &TravelEvent{
		EEvent: EEvent{
			duration:         e.duration,
			EventSolvingInfo: &data.ExpeditionStepResolveInfo{},
			Name:             e.Name,
			Reward:           e.Reward.GetCopy(),
		},
	}
}
