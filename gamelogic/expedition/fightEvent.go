package expedition

import (
	"time"
	"wvtrserv/data"
)

type FightEvent struct {
	EEvent
	ETeam *data.Team
}

func NewFightEvent(t *data.Team, name string) *FightEvent {
	return &FightEvent{
		EEvent: EEvent{
			duration: 0,
			name:     name,
		},
		ETeam: (*data.Team)(t),
	}
}

func (e FightEvent) GetEventType() data.EncounterState {
	return data.Fight
}

func (e FightEvent) Solve(startAt time.Time, heroTeam *data.Team) *data.ExpeditionStepResolveInfo {
	resExp := data.NewExpeditionResolveInfo(e.GetEventType())

	resExp.AddNewHappening(startAt, "Traveling Start")
	Fight(heroTeam, e.ETeam, resExp)
	resExp.AddNewHappening(startAt.Add(e.duration), "Traveling End")

	return resExp
}

func Fight(heroTeam *data.Team, enemyTeam *data.Team, infos *data.ExpeditionStepResolveInfo) {

}
