package expedition

import (
	"fmt"
	"time"
	"wvtrserv/databasemodel"
)

type Expedition struct {
	StartedAt time.Time
	Events    []ExpeditionEvent
}

func (e Expedition) Solve(identifier string, pTeam *databasemodel.Team) *databasemodel.ExpeditionDB {
	fmt.Printf("Solve expedition :\n")
	var t time.Time = time.Now()
	happened := make([]*databasemodel.ExpeditionStepResolveInfo, 0)
	for _, ev := range e.Events {
		happened = append(happened, ev.Solve(t, pTeam))
	}
	edb := &databasemodel.ExpeditionDB{
		Identifier:   identifier,
		StartedAt:    t.UTC(),
		WhatHappened: happened,
	}
	return edb
}

func (e Expedition) GetMinimumTotalTime() time.Duration {
	var res time.Duration = 0
	for _, ev := range e.Events {
		res += ev.GetDuration()
	}
	return res
}

func (e Expedition) GetEnemyTeamForEvent(idx int) *databasemodel.Team {
	event := e.Events[idx]
	fight := event.(FightEvent)
	return fight.ETeam
}
