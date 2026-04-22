package databasecontroller

import "wvtrserv/data"

func CreateSkill(skill *data.Skill) {
	db.Create(skill)
}

func GetSkills() []*data.Skill {
	res := []*data.Skill{}
	db.Find(&res)
	return res
}
