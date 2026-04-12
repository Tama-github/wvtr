package dcole

import (
	"wvtrserv/stypes"
	"time"
)

var waifus = []stypes.Waifu {
	{
		Id: -1,
		ImageUrl: "/imgs/noimage.jpg",
	},
	{
		Id: 1,
		ImageUrl: "/imgs/otaku.png",
	},
	{
		Id: 2,
		ImageUrl: "/imgs/otaku2.png",
	},
	{
		Id: 3,
		ImageUrl: "/imgs/emilia.jpg",
	},
	{
		Id: 4,
		ImageUrl: "/imgs/eren.jpg",
	},
	{
		Id: 5,
		ImageUrl: "/imgs/satoru.png",
	},
}

var teams = []stypes.Team {
	{
		Id: -1,
		Waifus: [3]stypes.Waifu{},
	},
	{
		Id: 1,
		Waifus: [3]stypes.Waifu{waifus[1], waifus[0], waifus[0]},
	},
	{
		Id: 2,
		Waifus: [3]stypes.Waifu{waifus[1], waifus[0], waifus[2]},
	},
	{
		Id: 3,
		Waifus: [3]stypes.Waifu{waifus[2], waifus[2], waifus[2]},
	},
	{
		Id: 4,
		Waifus: [3]stypes.Waifu{waifus[3], waifus[4], waifus[4]},
	},

}

var users = []stypes.User {
	{
		Id: -1,
		Name: "Error User",
		State: stypes.GameState{
			IsBusy: false,
			State: stypes.Error,
		},
		CurrentTeam: teams[0],
		LastActionTime: time.Now(),
	},
	{
		Id: 1,
		Name: "Tama",
		State: stypes.GameState{
			IsBusy: false,
			State: stypes.Home,
		},
		CurrentTeam: teams[4],
		LastActionTime: time.Now(),
	},
	{
		Id: 2,
		Name: "Tama",
		State: stypes.GameState{
			IsBusy: true,
			State: stypes.Travel,
		},
		CurrentTeam: teams[4],
		LastActionTime: time.Now(),
	},
	{
		Id: 3,
		Name: "Tama",
		State: stypes.GameState{
			IsBusy: true,
			State: stypes.Fight,
		},
		CurrentTeam: teams[4],
		LastActionTime: time.Now(),
	},
	{
		Id: 4,
		Name: "Tama",
		State: stypes.GameState{
			IsBusy: false,
			State: stypes.Neutral,
		},
		CurrentTeam: teams[4],
		LastActionTime: time.Now(),
	},
}

func GetWaifuByID(id int) stypes.Waifu {
	for _,w := range waifus {
		if id == w.Id {
			return w
		}
	}
	return waifus[0]
}

func GetTeamByID(id int) stypes.Team {
	for _,t := range teams {
		if id == t.Id {
			return t
		}
	}
	return teams[0]
}

func GetUserByID(id int) stypes.User {
	for _,u := range users {
		if id == u.Id {
			return u
		}
	}
	return users[0]
}
