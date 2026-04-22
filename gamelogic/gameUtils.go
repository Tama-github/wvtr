package gamelogic

import "math/rand/v2"

type PoolElement struct {
	Obj    interface{}
	Weight float64
}

func RollCheck(rollScore float64, target float64) bool {
	return rollScore >= target
}

func NaturalRoll(min float64, max float64) float64 {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}

	return rand.Float64()
}

// Arrays
func RollInArrayWithRate(randomNumber float64, rollRateArray []float64) int {
	acc := 0.0
	for i := range len(rollRateArray) {
		acc += rollRateArray[i]
		if randomNumber <= acc {
			return i
		}
	}
	return len(rollRateArray) - 1
}

func NormalizeArray(array []float64) []float64 {
	total := 0.0
	res := array
	for i := range len(array) {
		total += array[i]
	}

	for i := range len(array) {
		res[i] = array[i] / total
	}

	return res
}

// Pool
func RollInPool(randomNumber float64, rollRateArray []PoolElement) int {
	acc := 0.0
	for i := range len(rollRateArray) {
		acc += rollRateArray[i].Weight
		if randomNumber <= acc {
			return i
		}
	}
	return len(rollRateArray) - 1
}

func NormalizePool(array []PoolElement) []PoolElement {
	total := 0.0
	res := array
	for i := range len(array) {
		total += array[i].Weight
	}

	for i := range len(array) {
		res[i].Weight = array[i].Weight / total
	}

	return res
}
