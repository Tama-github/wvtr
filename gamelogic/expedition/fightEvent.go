package expedition

import (
	"fmt"
	"time"
	"wvtrserv/databasemodel"
)

type FightEvent struct {
	EEvent
	ETeam *databasemodel.Team
}

func NewFightEvent(t *databasemodel.Team) *FightEvent {
	return &FightEvent{
		EEvent: EEvent{
			duration: 0,
		},
		ETeam: t,
	}
}

func (e FightEvent) GetEventType() databasemodel.EncounterState {
	return databasemodel.Fight
}

func (fEvent FightEvent) Solve(startAt time.Time, team *databasemodel.Team) *databasemodel.ExpeditionStepResolveInfo {
	resolvInfo := fmt.Sprintf("T: {start: %s, end: %s}", startAt.String(), startAt.Add(fEvent.duration).String())
	fmt.Printf("Solve fight event : %s\n", resolvInfo)
	//fEvent.ETeam.Heroes[0] team.Heroes[0].
	return &databasemodel.ExpeditionStepResolveInfo{
		StepInfos: resolvInfo,
		StepEndAt: startAt.Add(fEvent.duration),
		StepState: fEvent.GetEventType(),
	}
}
