package expedition

import (
	"time"
	"wvtrserv/data"
)

type FightEvent struct {
	EEvent
	EnemyPool []*data.Hero
}

func NewFightEvent(areaEnemyPool []*data.Hero, name string, rewardPool *RewardPool) *FightEvent {
	res := &FightEvent{
		EEvent: EEvent{
			duration: 0,
			Name:     name,
			Reward:   NewReward(rewardPool),
		},
		EnemyPool: areaEnemyPool,
	}
	res.EventSolvingInfo = data.NewExpeditionResolveInfo(res.GetEventType())

	return res
}

func (e FightEvent) GetEventType() data.EncounterState {
	return data.Fight
}

func (e FightEvent) GenerateTeamToFightFromAreaPool() *data.Team {
	var teamHeroes []*data.Hero = make([]*data.Hero, 0)
	numberOfEnemies := int(data.NaturalRoll(1, 4))
	for range numberOfEnemies {
		ene := e.EnemyPool[int(data.NaturalRoll(0, float64(len(e.EnemyPool))))]
		ene.Equipment.Weapon.RollBaseStats()
		ene.ClearAllStatusAndSetToFullLife()
		teamHeroes = append(teamHeroes, ene)
	}
	return &data.Team{
		Heroes: teamHeroes,
	}
}

func (e *FightEvent) Solve(startAt time.Time, heroTeam *data.Team, toSpend []data.IStorable) *data.ExpeditionStepResolveInfo {
	resExp := data.NewExpeditionResolveInfo(e.GetEventType())

	resExp.AddNewHappening(startAt, "Fight start", nil)
	resExp.ETeam = e.GenerateTeamToFightFromAreaPool()

	heroTeam.Fight(resExp.ETeam, resExp)
	e.duration = resExp.GetDuration()
	resExp.AddNewHappening(startAt.Add(e.GetDuration()), "Fight End", nil)
	if !heroTeam.IsDefeated() {
		e.Reward.GenRandomReward()
		for _, h := range resExp.ETeam.Heroes {
			e.Reward.XP += xpToGainFromBeatingEnemy(h)
		}
	}
	return resExp
}

func (e FightEvent) CopyEvent() ExpeditionEvent {
	return &FightEvent{
		EEvent: EEvent{
			duration:         e.duration,
			EventSolvingInfo: &data.ExpeditionStepResolveInfo{},
			Name:             e.Name,
			Reward:           e.Reward.GetCopy(),
		},
		EnemyPool: e.EnemyPool,
	}
}

func xpToGainFromBeatingEnemy(e *data.Hero) float64 {
	x := float64(e.Attributes.Level)
	return x + (e.RollNumber(0, x/2))
}

func (e FightEvent) GenRewards(eTeam *data.Team) {
	for _, h := range eTeam.Heroes {
		e.Reward.AddXP(xpToGainFromBeatingEnemy(h))
	}
	for data.RollCheck(data.NaturalRoll(0, 1), e.Reward.LootChance) {
		e.Reward.GenRandomReward()
	}
}
