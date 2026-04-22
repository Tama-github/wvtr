package expedition

import (
	"fmt"
	"time"
	"wvtrserv/data"
	"wvtrserv/logger"
)

type Expedition struct {
	StartedAt time.Time
	Events    []ExpeditionEvent
}

func (e Expedition) Solve(identifier string, pTeam *data.Team) *data.ExpeditionDB {
	fmt.Printf("Solve expedition :\n")
	var t time.Time = time.Now()
	happened := make([]*data.ExpeditionStepResolveInfo, 0)
	logger.DumpLog.Println("solving expedition: ")
	for _, ev := range e.Events {
		happened = append(happened, ev.Solve(t, pTeam))
		t = t.Add(ev.GetDuration())
	}
	edb := &data.ExpeditionDB{
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

func (e Expedition) GetEnemyTeamForEvent(idx int) *data.Team {
	event := e.Events[idx]
	fight := event.(FightEvent)
	return fight.ETeam
}
