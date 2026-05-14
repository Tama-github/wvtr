package data

func (u User) UserHasAProblem() bool {
	return u.State.State == Error
}

func (u User) UserIsHome() bool {
	return u.State.State == Home
}

func (u *User) GetReward(expReward *Reward) {
	xpGained := expReward.XP / float64(len(u.CurrentTeam.Heroes))
	for _, h := range u.CurrentTeam.Heroes {
		h.GainXP(xpGained)
	}
	u.Inventory.Merge(expReward.Loot)
	expReward.Loot = nil
	expReward.LootID = 0
}

func (u *User) GetOwnedHeroByWaifuID(id string) *Hero {
	for _, oh := range u.OwnedHeroes {
		if *oh.WaifuID == id {
			return oh
		}
	}
	return nil
}
