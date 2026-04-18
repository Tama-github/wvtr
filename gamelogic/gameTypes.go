package gamelogic

// import (
// 	"fmt"
// 	"time"
// )

// type EncounterState int
// const (
// 	Home EncounterState = iota + 1
// 	Travel
// 	Fight
// 	Neutral
// 	Error
// )

// type ExpeditionEvent interface {
// 	EndAt(startAt time.Time) time.Time
// 	GetEventType() EncounterState
// 	Solve(startAt time.Time, t *Team) *ExpeditionStepResolveInfo
// 	GetDuration() time.Duration
// }

// type EEvent struct {
// 	duration   time.Duration
// 	solveTrace string
// }

// func (e EEvent) EndAt(startAt time.Time) time.Time {
// 	return startAt.Add(e.duration)
// }

// func (e EEvent) GetDuration() time.Duration {
// 	return e.duration
// }

// type HappeningType func(*Team) string

// /***********************/
// /***  Neutral Event  ***/
// /***********************/
// type NeutralEvent struct {
// 	EEvent
// 	Happening HappeningType
// }

// func NewNeutralEvent(duration time.Duration, h HappeningType) *NeutralEvent {
// 	return &NeutralEvent{
// 		EEvent: EEvent{
// 			duration: duration,
// 		},
// 		Happening: h,
// 	}
// }

// func (e NeutralEvent) GetEventType() EncounterState {
// 	return Neutral
// }

// func (e NeutralEvent) Solve(startAt time.Time, t *Team) *ExpeditionStepResolveInfo {

// 	resolvInfo := fmt.Sprintf("T: {start: %s, end: %s, %s}", startAt.String(), startAt.Add(e.duration).String(), e.Happening(t))
// 	fmt.Printf("Solve neutral event : %s\n", resolvInfo)
// 	return &ExpeditionStepResolveInfo{
// 		StepInfos: resolvInfo,
// 		StepEndAt: startAt.Add(e.duration),
// 		StepState: e.GetEventType(),
// 	}
// }

// /***********************/
// /***  Travel Event   ***/
// /***********************/
// type TravelEvent struct {
// 	EEvent
// }

// func NewTravelEvent(duration time.Duration) *TravelEvent {
// 	return &TravelEvent{
// 		EEvent{
// 			duration: duration,
// 		},
// 	}
// }

// func (e TravelEvent) GetEventType() EncounterState {
// 	return Travel
// }

// func (e TravelEvent) Solve(startAt time.Time, t *Team) *ExpeditionStepResolveInfo {
// 	resolvInfo := fmt.Sprintf("T: {start: %s, end: %s}", startAt.String(), startAt.Add(e.duration).String())
// 	return &ExpeditionStepResolveInfo{
// 		StepInfos: resolvInfo,
// 		StepEndAt: startAt.Add(e.duration),
// 		StepState: e.GetEventType(),
// 	}
// }

// /***********************/
// /***   Fight Event   ***/
// /***********************/
// type FightEvent struct {
// 	EEvent
// 	ETeam *Team
// }

// func NewFightEvent(t *Team) *FightEvent {
// 	return &FightEvent{
// 		EEvent: EEvent{
// 			duration: 0,
// 		},
// 		ETeam: t,
// 	}
// }

// func (e FightEvent) GetEventType() EncounterState {
// 	return Fight
// }

// func (e FightEvent) Solve(startAt time.Time, t *Team) *ExpeditionStepResolveInfo {
// 	resolvInfo := fmt.Sprintf("T: {start: %s, end: %s}", startAt.String(), startAt.Add(e.duration).String())
// 	fmt.Printf("Solve fight event : %s\n", resolvInfo)
// 	return &ExpeditionStepResolveInfo{
// 		StepInfos: resolvInfo,
// 		StepEndAt: startAt.Add(e.duration),
// 		StepState: e.GetEventType(),
// 	}
// }

// type Expedition struct {
// 	StartedAt time.Time
// 	Events    []ExpeditionEvent
// }

// func (e Expedition) Solve(identifier string, pTeam *Team) *ExpeditionDB {
// 	fmt.Printf("Solve expedition :\n")
// 	var t time.Time = time.Now()
// 	happened := make([]*ExpeditionStepResolveInfo, 0)
// 	for _, ev := range e.Events {
// 		happened = append(happened, ev.Solve(t, pTeam))
// 	}
// 	edb := &ExpeditionDB{
// 		Identifier:   identifier,
// 		StartedAt:    t.UTC(),
// 		WhatHappened: happened,
// 	}
// 	return edb
// }

// func (e Expedition) GetMinimumTotalTime() time.Duration {
// 	var res time.Duration = 0
// 	for _, ev := range e.Events {
// 		res += ev.GetDuration()
// 	}
// 	return res
// }

// func (e Expedition) GetEnemyTeamForEvent(idx int) *Team {
// 	event := e.Events[idx]
// 	fight := event.(FightEvent)
// 	return fight.ETeam
// }

// func test() {
// 	f := TravelEvent{
// 		EEvent{
// 			duration: time.Second * 30,
// 		},
// 	}
// 	t := NewTravelEvent(time.Second * 40)

// 	endat := f.EndAt(time.Now())
// 	fmt.Printf("this travel event will end at : ")
// 	fmt.Println(endat)

// 	endat2 := t.EndAt(time.Now())
// 	fmt.Printf("this travel event will end at : ")
// 	fmt.Println(endat2)
// }
