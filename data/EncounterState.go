package data

type EncounterState int

const (
	Home EncounterState = iota + 1
	Travel
	Fight
	Neutral
	Error
)

func (e EncounterState) String() string {
	res := "Error"
	switch e {
	case Home:
		res = "Home"
	case Travel:
		res = "Travel"
	case Fight:
		res = "Fight"
	case Neutral:
		res = "Neutral"
	case Error:
		res = "Error"
	}
	return res
}
