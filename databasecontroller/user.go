package databasecontroller

import (
	"wvtrserv/databasemodel"
)

func UserHasAProblem(u *databasemodel.User) bool {
	return u.State.State == databasemodel.Error
}

func UserIsHome(u *databasemodel.User) bool {
	return u.State.State == databasemodel.Home
}
