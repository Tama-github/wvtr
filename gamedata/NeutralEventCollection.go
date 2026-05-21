package gamedata

import (
	"fmt"
	"time"
	"wvtrserv/data"
	"wvtrserv/gamelogic/expedition"
)

// Nothing events
var nothing10s expedition.ExpeditionEvent = expedition.NewNeutralEvent(time.Second*10, "Neutral", nil,
	expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo, toSpend []data.IStorable) {
	}))

// train events
func trainingEventFactory(duration time.Duration, quantity float64) expedition.ExpeditionEvent {
	return expedition.NewNeutralEvent(duration, "Training", nil,
		expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo, toSpend []data.IStorable) {
			selfEvent.GetReward().AddXP(quantity)
			e.AddNewHappening(time.Now().Add(duration-2*time.Millisecond), fmt.Sprintf("Team has trained and got %fxp", quantity), nil)
		}))
}

var selfTrainingTest expedition.ExpeditionEvent = trainingEventFactory(3*time.Second, 5)

var selfTraining expedition.ExpeditionEvent = trainingEventFactory(2*time.Hour, 5)
var trainWithEquipment expedition.ExpeditionEvent = trainingEventFactory(2*time.Hour, 20)
var trainWithMentor expedition.ExpeditionEvent = trainingEventFactory(2*time.Hour, 100)

// rest events
func restEventFactory(duration time.Duration, quantity float64) expedition.ExpeditionEvent {
	return expedition.NewNeutralEvent(duration, "Resting", nil,
		expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo, toSpend []data.IStorable) {
			for _, h := range t.Heroes {
				healed := h.Rest(quantity)
				e.AddNewHappening(time.Now().Add(duration-2*time.Millisecond), fmt.Sprintf("%s healed for %f hp points.", h.Name, healed), &data.FieldActionDesc{
					FromH:        h,
					FromPVChange: -healed,
				})
			}
		}))
}

var restQualityMultiplicator_Bad float64 = 0.5
var restQualityMultiplicator_Normal float64 = 1.0
var restQualityMultiplicator_Good float64 = 2.0
var restQualityMultiplicator_Excellent float64 = 4.0

var restAmout float64 = 10.0
var healerRestAmout float64 = 50.0
var hospitalRestAmount float64 = 100.0

var sizeMultiplicator_small float64 = 1.0
var sizeMultiplicator_medium float64 = 2.0
var sizeMultiplicator_big float64 = 3.0

var dayTime = 20 * time.Hour
var dayTest = 2 * time.Second

var testsmallRest1 expedition.ExpeditionEvent = restEventFactory(dayTest, restAmout)
var testsmallRest2 expedition.ExpeditionEvent = restEventFactory(dayTest, 15)
var testsmallRest3 expedition.ExpeditionEvent = restEventFactory(dayTest, 20)

// Used only in expedition
var smallRest_Bad expedition.ExpeditionEvent = restEventFactory(1*time.Hour, restAmout*sizeMultiplicator_small*restQualityMultiplicator_Bad)
var smallRest_Normal expedition.ExpeditionEvent = restEventFactory(1*time.Hour, restAmout*sizeMultiplicator_small*restQualityMultiplicator_Normal)
var smallRest_Good expedition.ExpeditionEvent = restEventFactory(1*time.Hour, restAmout*sizeMultiplicator_small*restQualityMultiplicator_Good)
var smallRest_Excellent expedition.ExpeditionEvent = restEventFactory(1*time.Hour, restAmout*sizeMultiplicator_small*restQualityMultiplicator_Excellent)

var mediumRest_Bad expedition.ExpeditionEvent = restEventFactory(8*time.Hour, restAmout*sizeMultiplicator_medium*restQualityMultiplicator_Bad)
var mediumRest_Normal expedition.ExpeditionEvent = restEventFactory(8*time.Hour, restAmout*sizeMultiplicator_medium*restQualityMultiplicator_Normal)
var mediumRest_Good expedition.ExpeditionEvent = restEventFactory(8*time.Hour, restAmout*sizeMultiplicator_medium*restQualityMultiplicator_Good)
var mediumRest_Excellent expedition.ExpeditionEvent = restEventFactory(8*time.Hour, restAmout*sizeMultiplicator_medium*restQualityMultiplicator_Excellent)

var bigRest_Bad expedition.ExpeditionEvent = restEventFactory(12*time.Hour, restAmout*sizeMultiplicator_big*restQualityMultiplicator_Bad)
var bigRest_Normal expedition.ExpeditionEvent = restEventFactory(12*time.Hour, restAmout*sizeMultiplicator_big*restQualityMultiplicator_Normal)
var bigRest_Good expedition.ExpeditionEvent = restEventFactory(12*time.Hour, restAmout*sizeMultiplicator_big*restQualityMultiplicator_Good)
var bigRest_Excellent expedition.ExpeditionEvent = restEventFactory(12*time.Hour, restAmout*sizeMultiplicator_big*restQualityMultiplicator_Excellent)

// For healing outside expeditions
var rest expedition.ExpeditionEvent = restEventFactory(20*time.Hour, restAmout*restQualityMultiplicator_Good)
var healerRest expedition.ExpeditionEvent = restEventFactory(20*time.Hour, healerRestAmout*restQualityMultiplicator_Good)
var hospitalRest expedition.ExpeditionEvent = restEventFactory(20*time.Hour, hospitalRestAmount*restQualityMultiplicator_Good)

// Work events
func workEventFactory(duration time.Duration, quantity float64) expedition.ExpeditionEvent {
	return expedition.NewNeutralEvent(duration, "Working",
		&expedition.RewardPool{
			CurrencyPool: map[data.CurrencyType]data.StatsRange{
				data.Gold: {Min: quantity, Max: quantity},
			},
		},
		expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo, toSpend []data.IStorable) {
			e.AddNewHappening(time.Now().Add(duration-2*time.Millisecond), fmt.Sprintf("Gained %f gold", quantity), nil)
		}))
}

var workShortTest expedition.ExpeditionEvent = workEventFactory(3*time.Second, 5)
var workGold expedition.ExpeditionEvent = workEventFactory(20*time.Hour, 20)

// Crafting
func craftEventFactory(duration time.Duration, bases *expedition.RewardPool, proba float64) expedition.ExpeditionEvent {
	res := expedition.NewNeutralEvent(duration, "Crafting",
		bases,
		expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo, toSpend []data.IStorable) {
			e.AddNewHappening(time.Now().Add(duration-2*time.Millisecond), "Crafted", nil)
		}))

	res.EEvent.Reward.LootChance = proba
	return res
}

var craftWeapon expedition.ExpeditionEvent = craftEventFactory(dayTest, craftWeaponsRewardPool, 1)
var craftArmor expedition.ExpeditionEvent = craftEventFactory(dayTest, craftArmorsRewardPool, 1)
var craftOmamori expedition.ExpeditionEvent = craftEventFactory(dayTest, craftOmamorisRewardPool, 1)

// Scrap Weapon
func scrapEventFactory(duration time.Duration) expedition.ExpeditionEvent {
	return expedition.NewNeutralEvent(duration, "Working",
		&expedition.RewardPool{
			CurrencyPool: map[data.CurrencyType]data.StatsRange{
				data.CScrap: {Min: 1, Max: 6},
				data.LSCrap: {Min: 1, Max: 6},
				data.MScrap: {Min: 1, Max: 6},
			},
		},
		expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo, toSpend []data.IStorable) {
			// TODO: Make the function give reward depending on the spent resources
			for _, eq := range toSpend {
				e.AddNewHappening(time.Now().Add(duration-2*time.Millisecond), fmt.Sprintf("Scraped %s", eq.GetName()), nil)
			}
		}))
}
