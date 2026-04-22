package data

func (u User) UserHasAProblem() bool {
	return u.State.State == Error
}

func (u User) UserIsHome() bool {
	return u.State.State == Home
}
