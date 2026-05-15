package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wvtrserv/data"
	"wvtrserv/databasecontroller"
	"wvtrserv/gamedata"
	"wvtrserv/gamelogic/expedition"
	"wvtrserv/logger"
	"wvtrserv/nanapi/client"
	"wvtrserv/nanapi/config"
	"wvtrserv/utils"
)

type AuthEndpoints struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndPoint         string `json:"token_endpoint"`
	UserInfoEndPoint      string `json:"userinfo_endpoint"`
}

type UserToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
}

var authEndPoints *AuthEndpoints = &AuthEndpoints{}

func fetchAuthEndpoints() {
	resp := utils.Fetch(config.GetNanapiConfig().OIDCURL+"/.well-known/openid-configuration", "GET", "", []string{"Content-Type", "application/json"})
	if resp == nil {
		return
	}

	authEndPoints = utils.DecodeJson(authEndPoints, resp.Body)
}

type DiscordAccount struct {
	Name      string `json:"name"`
	DiscordID string `json:"discord_id"`
}

func handlerAuth(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerAuth]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	tokenEndpoint := authEndPoints.AuthorizationEndpoint
	logger.DumpLog.Println(tokenEndpoint)
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", config.GetNanapiConfig().OIDCCLientId)
	params.Add("redirect_uri", config.GetNanapiConfig().DomainName+"/api/oidc/callback")
	params.Add("scope", "openid profile discord_id")
	http.Redirect(w, r, tokenEndpoint+"?"+params.Encode(), 302)
}

// Connexion
func handlerConnexion(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerConnexion]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	code := r.URL.Query().Get("code")
	logger.DumpLog.Println(code)
	clientId := config.GetNanapiConfig().OIDCCLientId
	clientSecret := config.GetNanapiConfig().OIDCCLientSecret

	tokenEndpoint := authEndPoints.TokenEndPoint
	logger.DumpLog.Println(tokenEndpoint)
	methode := "POST"
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("code", code)
	params.Add("client_id", clientId)
	params.Add("client_secret", clientSecret)
	params.Add("redirect_uri", config.GetNanapiConfig().DomainName+"/api/oidc/callback")

	header := []string{"Content-Type", "application/x-www-form-urlencoded"}
	logger.DumpLog.Println(params.Encode())
	tokenResp := utils.Fetch(tokenEndpoint, methode, params.Encode(), header)

	uToken := &UserToken{}

	uToken = utils.DecodeJson(uToken, tokenResp.Body)
	if uToken == nil {
		return
	}

	userResp := utils.Fetch(authEndPoints.UserInfoEndPoint, "", "", []string{"Authorization", "Bearer " + uToken.AccessToken})

	discordAccount := &DiscordAccount{}

	discordAccount = utils.DecodeJson(discordAccount, userResp.Body)
	if discordAccount == nil {
		return
	}

	user := databasecontroller.GetUserByDiscordID(discordAccount.DiscordID)
	user = databasecontroller.GetUserGameState(user)

	// this means it's the first time the user arrive here.
	// and we need to create a new user based on the discord account info
	if user.DiscordID == "" {
		c := data.NewCurrencyOwned(databasecontroller.GetAllCurrencies())
		user = &data.User{
			Name: discordAccount.Name,
			CurrentTeam: &data.Team{
				Heroes: make([]*data.Hero, 0),
			},
			Inventory:   data.NewInventory(c),
			OwnedHeroes: make([]*data.Hero, 0),
			State: &data.GameState{
				State: data.Home,
			},
			DiscordID: discordAccount.DiscordID,
		}
		databasecontroller.CreateNewUser(user)
		logger.DumpLog.Println("Test")
		user = databasecontroller.GetUserByDiscordID(discordAccount.DiscordID)
	}

	redirectParams := url.Values{}
	redirectParams.Add("wvtrusrid", fmt.Sprintf("%d", user.ID))

	redReq, err := http.NewRequest("POST", config.GetNanapiConfig().DomainName, strings.NewReader(redirectParams.Encode()))
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	reqPath := config.GetNanapiConfig().DomainName + "?" + redirectParams.Encode()
	logger.DumpLog.Println(reqPath)
	http.Redirect(w, redReq, reqPath, http.StatusSeeOther)
}

func handleGetPlayerWaicolleAscendedWaifus(w http.ResponseWriter, r *http.Request) {
	functionS := "[handleGetPlayerWaicolleAscendedWaifus]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	id := utils.GetParamInt("id", r)

	user := databasecontroller.GetUserByID(uint(id))

	toSend := "{}"

	waifus := client.GetAvailableWaifuToSendToWVTR(user)
	if waifus == nil {
		logger.ErrLog.Printf("%s can't get response from nanpi with user [%d]", functionS, user.ID)
		fmt.Fprintf(w, "%s", toSend)
		return
	}

	utils.Give(waifus, w, false)
}

type EquipmentType int

const (
	WeaponType EquipmentType = iota
	ArmorType
	OmamoriType
)

func handleEquip(w http.ResponseWriter, r *http.Request) {
	functionS := "[handleEquip]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	uid := utils.GetParamInt("uid", r)
	hid := utils.GetParamInt("hid", r)
	var etype EquipmentType = EquipmentType(utils.GetParamInt("etype", r))
	eid := utils.GetParamInt("eid", r)

	user := databasecontroller.GetUserByID(uint(uid))
	hero := databasecontroller.GetHeroByID(uint(hid))

	switch etype {
	case WeaponType:
		wtoe := databasecontroller.GetWeaponByID(uint(eid))
		user.Inventory.Store(hero.Equip(wtoe).(*data.Weapon), 1)
	case ArmorType:
		atoe := databasecontroller.GetArmorByID(uint(eid))
		user.Inventory.Store(hero.Equip(atoe).(*data.Armor), 1)
	case OmamoriType:
		otoe := databasecontroller.GetOmamoriByID(uint(eid))
		user.Inventory.Store(hero.Equip(otoe).(*data.Omamori), 1)
	}

	databasecontroller.SaveInventory(user.Inventory)
	databasecontroller.SaveHero(hero)

	toSend := "{}"

	utils.Give(toSend, w, false)
}

func handlerCreateHeroForPlayer(w http.ResponseWriter, r *http.Request) {
	functionS := "[handleGetPlayerWaicolleAscendedWaifus]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	id := utils.GetParamInt("id", r)

	waifu := &client.JoinWC{}
	waifu = utils.DecodeJson(waifu, r.Body)
	if waifu == nil {
		return
	}

	logger.DumpLog.Println(waifu)

	newH := gamedata.CreateNewHeroFromDBWaifuInfos(waifu, databasecontroller.GetHeroClasses(), databasecontroller.GetSkills())
	newH.UserID = uint(id)
	databasecontroller.CreateWeapon(newH.Equipment.Weapon)
	errReq := databasecontroller.CreateHero(newH)
	if errReq != nil {
		logger.ErrLog.Println("Can't Create new hero: ", errReq)
		fmt.Fprintf(w, "%s is already a hero", waifu.NameUserPreferred)
		return
	}

	utils.Give(newH, w, true)
}

// Getters
func handlerHero(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerHero]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	id := utils.GetParamInt("id", r)

	hero := databasecontroller.GetHeroByID(uint(id))

	utils.Give(hero, w, true)
}

func handlerTeam(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerTeam]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	id := utils.GetParamInt("id", r)

	team := databasecontroller.GetTeamByID(uint(id))

	utils.Give(team, w, true)
}

func handlerInventory(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerInventory]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	id := utils.GetParamInt("id", r)

	inv := databasecontroller.GetInventoryByID(uint(id))

	utils.Give(inv, w, true)
}

func handlerAvailableExpeditions(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerAvailableExpeditions]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	id := utils.GetParamInt("id", r)

	if id == 0 {
		fmt.Fprintf(w, "{}")
		return
	}
	user := databasecontroller.GetUserByID(uint(id))

	expeditions := gamedata.GetAvailableExpeditions(user)

	utils.Give(expeditions, w, true)
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerUser]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	id := utils.GetParamInt("id", r)

	user := databasecontroller.GetUserByID(uint(id))

	utils.Give(user, w, true)
}

// Updaters
func handlerSaveUser(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerSaveUser]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	if r.Method != http.MethodPost {
		s := fmt.Sprintf("Method not allowed (%s) POST expected.", r.Method)
		logger.ErrLog.Println(s)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}

	user := &data.User{}
	user = utils.DecodeJson(user, r.Body)
	if user == nil {
		return
	}

	databasecontroller.UpdateUser(user)
	w.WriteHeader(http.StatusCreated)
}

func handlerSaveGameState(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerSaveGameState]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	if r.Method != http.MethodPost {
		s := fmt.Sprintf("Method not allowed (%s) POST expected.", r.Method)
		logger.ErrLog.Println(s)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}

	state := &data.GameState{}

	state = utils.DecodeJson(state, r.Body)
	if state == nil {
		utils.Give("{}", w, true)
		return
	}

	databasecontroller.UpdateGameState(state)

	utils.Give(state, w, true)
}

func handlerUpdateTeam(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerUpdateTeam]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	if r.Method != http.MethodPost {
		s := fmt.Sprintf("%s Method not allowed (%s) POST expected.", functionS, r.Method)
		logger.ErrLog.Println(s)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}
	user := &data.User{}
	user = utils.DecodeJson(user, r.Body)
	if user == nil {
		return
	}

	databasecontroller.UpdateTeam(user.CurrentTeam)
	user = databasecontroller.GetUserByID(user.ID)

	w.WriteHeader(http.StatusCreated)
	utils.Give(user.CurrentTeam, w, true)
}

type CurrentStepRequestMessage struct {
	Uid  int   `json:"id"`
	Time int64 `json:"time"`
}

func handlerCurrentExpeditionStep(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerCurrentExpeditionStep]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	if r.Method != http.MethodPost {
		s := fmt.Sprintf("%s Method not allowed (%s) POST expected.", functionS, r.Method)
		logger.ErrLog.Println(s)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}
	var data *CurrentStepRequestMessage = &CurrentStepRequestMessage{}
	data = utils.DecodeJson(data, r.Body)
	if data == nil {
		return
	}

	var t time.Time = time.Unix(0, data.Time*int64(time.Millisecond))

	user := databasecontroller.GetUserByID(uint(data.Uid))
	res := databasecontroller.UpdateGameStateWithTime(user.State, &t)
	databasecontroller.UpdateTeamWithExpAndTime(user.CurrentTeam, *user.State.CurrentExpedition, t)
	databasecontroller.UpdateGameState(user.State)

	utils.Give(res, w, true)
}

func handlerLaunchExpedition(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerLaunchExpedition]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	id := utils.GetParamInt("usr", r)
	expCathegory := r.PathValue("expCat")
	expIdentifier := r.PathValue("expId")

	user := databasecontroller.GetUserByID(uint(id))
	var exp expedition.Expedition = gamedata.Expeditions[expCathegory][expIdentifier].GetCopy()
	if exp.CanEnter(user) && user.Inventory.Remove(exp.Cost, exp.CostNumber) {
		databasecontroller.SaveInventory(user.Inventory)
		c := data.NewCurrencyOwned(databasecontroller.GetAllCurrencies())
		databasecontroller.LaunchExpedition(user, exp.Solve(expIdentifier, user.CurrentTeam, c))
		user.CurrentTeam.ResetSkills()

		utils.Give(user.State.CurrentExpedition.WhatHappened[0], w, true)
		return
	}

	fmt.Fprintf(w, "Can't launch expedition")
}

func handlerExpeditionReport(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerExpeditionReport]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	id := utils.GetParamInt("uid", r)

	user := databasecontroller.GetUserByID(uint(id))
	exp := user.State.CurrentExpedition
	user.GetReward(exp.ExpeditionRewards)
	//databasecontroller.Delete(exp)
	databasecontroller.SaveInventory(user.Inventory)
	databasecontroller.UpdateUser(user)
	databasecontroller.SaveTeam(user.CurrentTeam)
	utils.Give(user.State.CurrentExpedition, w, true)
}

func main() {
	fetchAuthEndpoints()
	databasecontroller.DBLogIn()

	// Get main page
	fs := http.FileServer(http.Dir("./ui/wvtr-front/dist"))
	http.Handle("/", fs)

	// Connexion
	http.HandleFunc("/api/oidc/auth", handlerAuth)
	http.HandleFunc("/api/oidc/callback", handlerConnexion)

	// Request object by ID.
	//get
	http.HandleFunc("/api/hero/{id}", handlerHero)
	http.HandleFunc("/api/teams/{id}", handlerTeam)
	http.HandleFunc("/api/inventory/{id}", handlerInventory)
	http.HandleFunc("/api/expeditionReport/{uid}", handlerExpeditionReport)
	http.HandleFunc("/api/user/{id}", handlerUser)
	http.HandleFunc("/api/availableexpeditions/{id}", handlerAvailableExpeditions)
	http.HandleFunc("/api/userwaifus/{id}", handleGetPlayerWaicolleAscendedWaifus)
	http.HandleFunc("/api/equip/{uid}/{hid}/{etype}/{eid}", handleEquip)

	//post
	http.HandleFunc("/api/currentexpeditionstep/", handlerCurrentExpeditionStep)

	//Modify db
	//get
	http.HandleFunc("/api/launchExpedition/{usr}/{expCat}/{expId}", handlerLaunchExpedition)
	http.HandleFunc("/api/createherofromwaifu/{id}", handlerCreateHeroForPlayer)

	//post
	http.HandleFunc("/api/updateTeam/", handlerUpdateTeam)
	http.HandleFunc("/api/saveUser/", handlerSaveUser)
	http.HandleFunc("/api/saveGameState/", handlerSaveGameState)

	logger.DumpLog.Println("Listening on :4210...")
	err := http.ListenAndServe(":4210", nil)

	if err != nil {
		log.Fatal(err)
	}
}
