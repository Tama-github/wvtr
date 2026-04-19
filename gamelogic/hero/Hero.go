package hero

import (
	"wvtrserv/databasemodel"
)

type Hero databasemodel.Hero

// func test(dbh *databasemodel.Hero) {
// 	h := gamelogic.CreateLogicObjectFromDBObject[Hero, databasemodel.Hero](dbh)
// }

func (h Hero) Attack(target Hero) string {
	logResult := ""
	// get attack value

	// check miss

	// check critic

	// damage target

	// check if target dodged

	// chack if target blocked

	// add leach if there is

	// check if target is dead

	return logResult
}

func (h *Hero) takeFlatDamage(dmg int) {
	h.CurrentHP -= dmg
}

func (h *Hero) TakeDamage(dmg int, takeFrom Hero) (int, int) {
	// check dodge

	// check blocked

	// check resistances

	// get total tamage taken

	h.takeFlatDamage(dmg)

	// send reflected damage if there are any

	// check if dead

	return 0, dmg
}
