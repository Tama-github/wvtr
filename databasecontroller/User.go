package databasecontroller

import (
	"wvtrserv/data"
	"wvtrserv/logger"
)

func CreateNewUser(user *data.User) {
	logger.DumpLog.Print("CreateNewUser")
	db.Create(user)
}

func GetUserByID(id uint) *data.User {
	var res *data.User = nil
	// We get the full user values.
	db.Preload("CurrentTeam").
		Preload("CurrentTeam.Heroes").
		Preload("CurrentTeam.Heroes.Class").
		Preload("CurrentTeam.Heroes.UniqueSkill").
		Preload("CurrentTeam.Heroes.ActiveSkill").
		Preload("CurrentTeam.Heroes.Attributes").
		Preload("State").
		Preload("State.CurrentExpedition").
		Preload("State.CurrentExpedition.WhatHappened").
		Preload("State.ETeam").
		Preload("State.ETeam.Heroes").
		Preload("OwnedHeroes").
		Preload("OwnedHeroes.Class").
		Preload("OwnedHeroes.UniqueSkill").
		Preload("OwnedHeroes.ActiveSkill").
		Preload("OwnedHeroes.Attributes").
		Find(&res, id)

	logger.DumpLog.Println("Get user by id ", id)
	logger.DumpLog.Println(res)

	return res
}

func GetUserByDiscordID(did string) *data.User {
	var res *data.User = nil
	db.Where("discord_id = ?", did).Find(&res)
	if res != nil {
		logger.DumpLog.Println("GetUserByDiscordID: ", did, " | ", res.Name)
	} else {
		logger.DumpLog.Println("GetUserByDiscordID: ", did, " | user not found.")
	}
	return res
}

func UpdateUser(user *data.User) {
	db.Save(user)
}
