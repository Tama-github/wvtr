package databasecontroller

import "wvtrserv/data"

func UpdateTeam(user *data.User) {
	db.Model(&user.CurrentTeam).
		Association("Heroes").
		Replace(user.CurrentTeam.Heroes)
}

func GetTeamByID(id uint) *data.Team {
	var res *data.Team
	db.Preload("Heroes").
		Preload("Heroes.Class").
		Preload("Heroes.UniqueSkill").
		Preload("Heroes.ActiveSkill").
		Preload("Heroes.Attributes").
		Preload("Team").
		Find(&res, id)
	return res
}
