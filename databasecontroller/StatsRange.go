package databasecontroller

import "wvtrserv/data"

func GetStatsRangeByID(id uint) *data.StatsRange {
	var rng *data.StatsRange = &data.StatsRange{}
	db.Find(&rng, id)
	return rng
}

func SaveStatsRange(sr *data.StatsRange) {
	db.Save(sr)
}

func CreateStatsRange(sr *data.StatsRange) {
	db.Create(sr)
}
