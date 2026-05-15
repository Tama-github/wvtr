package data

import (
	"slices"
	"time"
	"wvtrserv/logger"
)

func (t *Team) IsDefeated() bool {
	for _, h := range t.Heroes {
		if !h.IsDefeated() {
			return false
		}
	}
	return true
}

type Action struct {
	When time.Time
	Who  *Hero
}

type FightTurnOrderTimeline []*Action

func NewFightTurnOrderTimeline(a Team, b Team, start time.Time) FightTurnOrderTimeline {
	res := FightTurnOrderTimeline(make([]*Action, 0))
	for _, h := range a.Heroes {
		initScore := h.InitiativeRoll()
		res = res.addAction(&Action{When: start.Add(initScore), Who: h})
	}
	for _, h := range b.Heroes {
		initScore := h.InitiativeRoll()
		res = res.addAction(&Action{When: start.Add(initScore), Who: h})
	}

	return res
}

func (f FightTurnOrderTimeline) fightTotalDuration(timestart time.Time) time.Duration {
	if len(f) == 0 {
		return 0
	}
	return f[len(f)-1].When.Sub(timestart)
}

func (f FightTurnOrderTimeline) addAction(ac *Action) FightTurnOrderTimeline {
	f = append(f, ac)

	// order by time
	slices.SortFunc(f, func(a, b *Action) int {
		if a.When.Compare(b.When) == 0 {
			b.When = b.When.Add(2)
		}
		return a.When.Compare(b.When)
	})
	return f
}

func (f FightTurnOrderTimeline) getNextAction(t time.Time) *Action {
	for _, a := range f {
		if a.When.After(t) {
			return a
		}
	}
	return nil
}

func (team *Team) ResetSkills() {
	for _, h := range team.Heroes {
		h.SetUSkillUsed(false)
	}
}

func (team *Team) Fight(oponent *Team, fightReport *ExpeditionStepResolveInfo) {
	startTime := fightReport.Timeline[0].When

	logger.DumpLog.Println("begin fight")
	turnOrder := NewFightTurnOrderTimeline(*team, *oponent, startTime)
	escapeTime := 30 * time.Minute
	time := startTime
	for !team.IsDefeated() && !oponent.IsDefeated() {
		if turnOrder.fightTotalDuration(startTime) > escapeTime {
			fightReport.AddNewHappening(time, "The fight was too long everyone fly away.", nil)
			return
		}

		a := turnOrder.getNextAction(time)
		logger.DumpLog.Printf("%s to play", a.Who.Name)

		// Check who are friends who are enemies
		ft := oponent
		et := team
		if slices.Contains(team.Heroes, a.Who) {
			ft = team
			et = oponent
		}

		logger.DumpLog.Printf("%s to play", a.Who.Name)
		recupTime := a.Who.Play(a.When, ft, et, fightReport)
		newAction := &Action{
			When: a.When.Add(recupTime),
			Who:  a.Who,
		}

		turnOrder = turnOrder.addAction(newAction)
		time = a.When
	}
	if team.IsDefeated() {
		fightReport.AddNewHappening(time, "The team have been defeated by their oponents.", nil)
	} else if oponent.IsDefeated() {
		fightReport.AddNewHappening(time, "The oponents have been defeated by the team.", nil)
	}
}

func (team *Team) getHero(h *Hero) *Hero {
	if h == nil {
		return nil
	}
	for _, he := range team.Heroes {
		if he.ID == h.ID {
			return he
		}
	}
	return nil
}

func (team *Team) ApplyESRI(esri *ExpeditionStepResolveInfo, start time.Time, until time.Time) *Team {
	if esri == nil || len(esri.Timeline) == 0 {
		return team
	}

	for _, tl := range esri.Timeline {

		if tl.When.After(until) {
			return team
		}
		if tl.When.Before(start) {
			continue
		}
		if tl.WhatAction == nil {
			continue
		}
		from := team.getHero(tl.WhatAction.FromH)
		if from != nil {
			from.Attributes.CurrentHP -= tl.WhatAction.FromPVChange
		}
		target := team.getHero(tl.WhatAction.TargetH)
		if target != nil {
			target.Attributes.CurrentHP -= tl.WhatAction.TargetPVChange
		}
	}
	return team
}
