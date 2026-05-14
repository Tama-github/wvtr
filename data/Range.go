package data

func (r *StatsRange) RollValue() StatsRange {
	if r == nil {
		return StatsRange{}
	}
	r.Value = NaturalRoll(r.Min, r.Max)
	r.ID = 0
	return *r
}
