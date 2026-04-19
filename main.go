package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"wvtrserv/databasecontroller"
	"wvtrserv/databasemodel"
	"wvtrserv/gameexpedition"
	"wvtrserv/logger"
	"wvtrserv/nanapi/client"
	"wvtrserv/utils"
)

const DOMAIN_NAME = "https://tama.rhiobet.sh"
const AUTH_SERVER = "https://auth.japan7.bde.enseeiht.fr"

// Main page ?
// func handler(w http.ResponseWriter, r *http.Request) {
// 	logger.DumpLog.Printf("req.Method: %s\n", r.Method)
// 	logger.DumpLog.Printf("req.URL.Path: %s\n", r.URL.Path)
// 	logger.DumpLog.Printf("req.ContentLength: %d\n", r.ContentLength)

// 	d := http.Dir("./ui/vu/UI/dist")
// 	f, err := d.Open("index.html")
// 	if err != nil {
// 		logger.ErrLog.Println(err)
// 	}

// 	defer f.Close()
// 	io.Copy(w, f)
// }

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
	resp := utils.Fetch(AUTH_SERVER+"/.well-known/openid-configuration", "GET", url.Values{}, []string{"Content-Type", "application/json"})
	if resp == nil {
		return
	}
	//defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(authEndPoints)
	if err != nil {
		logger.ErrLog.Printf("Problem while trying to get the authentificator endoints\n")
	}
}

// Main page
// func handleMainPage(w http.ResponseWriter, r *http.Request) {

// }

type DiscordAccount struct {
	Name      string `json:"name"`
	DiscordID string `json:"discord_id"`
}

// Connexion
func handlerConnexion(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerConnexion]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	// logger.DumpLog.Print(r)
	code := r.URL.Query().Get("code")
	logger.DumpLog.Println(code)
	clientId := "japan7"
	clientSecret := "nihongo"

	// Example usage
	tokenEndpoint := authEndPoints.TokenEndPoint
	logger.DumpLog.Println(tokenEndpoint)
	methode := "POST"
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("code", code)
	params.Add("client_id", clientId)
	params.Add("client_secret", clientSecret)
	params.Add("redirect_uri", DOMAIN_NAME+"/api/oidc/callback")

	header := []string{"Content-Type", "application/x-www-form-urlencoded"}
	logger.DumpLog.Println(params.Encode())
	tokenResp := utils.Fetch(tokenEndpoint, methode, params, header)

	uToken := &UserToken{}
	err := json.NewDecoder(tokenResp.Body).Decode(uToken)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}

	// Read and print response
	utils.ReadResponse(tokenResp)

	userResp := utils.Fetch(authEndPoints.UserInfoEndPoint, "", url.Values{}, []string{"Authorization", "Bearer " + uToken.AccessToken})
	//readResponse(userResp)

	discordAccount := &DiscordAccount{}
	decodError := json.NewDecoder(userResp.Body).Decode(discordAccount)
	if decodError != nil {
		logger.ErrLog.Println(decodError)
		return
	}

	user := databasecontroller.GetUserByDiscordID(discordAccount.DiscordID)

	// this means it's the first time the user arrive here.
	// and we need to create a new user based on the discord account info
	if user.DiscordID == "" {
		user = &databasemodel.User{
			Name: discordAccount.Name,
			CurrentTeam: &databasemodel.Team{
				Heroes: make([]*databasemodel.Hero, 0),
			},
			OwnedHeroes: make([]*databasemodel.Hero, 0),
			State: &databasemodel.GameState{
				State: databasemodel.Home,
			},
			DiscordID: discordAccount.DiscordID,
		}
		databasecontroller.CreateNewUser(user)
		user = databasecontroller.GetUserByDiscordID(discordAccount.DiscordID)
	}

	redirectParams := url.Values{}
	redirectParams.Add("wvtrusrid", fmt.Sprintf("%d", user.ID))

	redReq, err := http.NewRequest("POST", DOMAIN_NAME, strings.NewReader(redirectParams.Encode()))
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	reqPath := DOMAIN_NAME + "?" + redirectParams.Encode()
	logger.DumpLog.Println(reqPath)
	http.Redirect(w, redReq, reqPath, http.StatusSeeOther)
}

func handleGetPlayerWaicolleAscendedWaifus(w http.ResponseWriter, r *http.Request) {
	functionS := "[handleGetPlayerWaicolleAscendedWaifus]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	user := databasecontroller.GetUserByID(uint(id))

	waifusResponse := client.GetAvailableWaifuToSendToWVTR(user.DiscordID)
	if waifusResponse == nil {
		logger.ErrLog.Printf("%s can't get response from nanpi with user [%d]", functionS, user.ID)
		return
	}
	toSend := utils.ReadResponse(waifusResponse)
	fmt.Fprintf(w, "%s", toSend)
}

func handleCreateHeroForPlayer(w http.ResponseWriter, r *http.Request) {
	functionS := "[handleGetPlayerWaicolleAscendedWaifus]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)
	//user := databasecontroller.GetUserByID(uint(id))
	waifu := &client.Waifu{}
	err := json.NewDecoder(r.Body).Decode(waifu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	databasecontroller.CreateHero(&databasemodel.Hero{
		Name:           waifu.NameUserPreferred,
		ImageUrl:       waifu.ImageLarge,
		Level:          1,
		CurrentXP:      0,
		CurrentHP:      0,
		MaxHP:          0,
		XPBeforeLvlUp:  0,
		Attributes:     &databasemodel.HeroAttributes{},
		UserID:         uint(id),
		WaifuID:        waifu.ID,
		AnilistCharaID: waifu.IdAl,
	})

}

// Getters
func handlerHero(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerHero]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	hero := databasecontroller.GetHeroByID(uint(id))

	b, err := json.Marshal(hero)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerTeam(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerTeam]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	team := databasecontroller.GetTeamByID(uint(id))

	b, err := json.Marshal(team)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerAvailableExpeditions(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerAvailableExpeditions]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)

	expeditions := gameexpedition.GetAvailableExpeditions()

	b, err := json.Marshal(expeditions)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerUser]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	user := databasecontroller.GetUserByID(uint(id))

	b, err := json.Marshal(user)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
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
	user := &databasemodel.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		logger.ErrLog.Printf("%s Error when trying to get the user from the request body, got : %s", functionS, r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	databasecontroller.UpdateUser(user)
	w.WriteHeader(http.StatusCreated)
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
	user := &databasemodel.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		logger.ErrLog.Printf("%s Error when trying to get the user from the request body, got : %s", functionS, r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	databasecontroller.UpdateTeam(user)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{}")
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
	var data CurrentStepRequestMessage
	err := json.NewDecoder(r.Body).Decode(&data)
	var t time.Time = time.Unix(0, data.Time*int64(time.Millisecond))

	user := databasecontroller.GetUserByID(uint(data.Uid))

	if err != nil {
		logger.ErrLog.Printf("%s Error when trying to get the time : %s", functionS, r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := databasecontroller.UpdateGameStateWithTime(user.State, &t)
	databasecontroller.UpdateGameState(user.State)

	resS := "{}"
	if res != nil {
		b, err := json.Marshal(res)

		if err != nil {
			logger.ErrLog.Println(err)
			return
		}
		resS = string(b)
	}

	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, resS)
	fmt.Fprintf(w, "%s", resS)
}

func handlerLaunchExpedition(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerLaunchExpedition]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("usr")
	id, _ := strconv.Atoi(ids)

	expIdentifier := r.PathValue("expId")

	user := databasecontroller.GetUserByID(uint(id))

	databasecontroller.LaunchExpedition(user, gameexpedition.Expeditions[expIdentifier].Solve(expIdentifier, user.CurrentTeam))
	b, err := json.Marshal(user.State.CurrentExpedition.WhatHappened[0])
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func main() {
	fetchAuthEndpoints()
	databasecontroller.DBLogIn()

	// Get main page
	fs := http.FileServer(http.Dir("./ui/wvtr-front/dist"))
	http.Handle("/", fs)
	//http.HandleFunc("/", handleMainPage)

	// Connexion
	http.HandleFunc("/api/oidc/callback", handlerConnexion)

	// Request object by ID.
	//get
	http.HandleFunc("/hero/{id}", handlerHero)
	http.HandleFunc("/teams/{id}", handlerTeam)
	http.HandleFunc("/user/{id}", handlerUser)
	http.HandleFunc("/availableexpeditions/", handlerAvailableExpeditions)

	//post
	http.HandleFunc("/currentexpeditionstep/", handlerCurrentExpeditionStep)

	//Update existing values
	//get
	http.HandleFunc("/launchExpedition/{usr}/{expId}", handlerLaunchExpedition)

	//post
	http.HandleFunc("/updateTeam/", handlerUpdateTeam)

	// Images handler
	http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("imgs/"))))

	logger.DumpLog.Println("Listening on :4210...")
	err := http.ListenAndServe(":4210", nil)
	if err != nil {
		log.Fatal(err)
	}
}
