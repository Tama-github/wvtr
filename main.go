package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	"wvtrserv/logger"
	"wvtrserv/stypes"
)

// Main page ?
func handler(w http.ResponseWriter, r *http.Request) {
	logger.DumpLog.Printf("req.Method: %s\n", r.Method)
	logger.DumpLog.Printf("req.URL.Path: %s\n", r.URL.Path)
	logger.DumpLog.Printf("req.ContentLength: %d\n", r.ContentLength)

	d := http.Dir("./ui/vu/UI/dist")
	f, err := d.Open("index.html")
	if err != nil {
		logger.ErrLog.Println(err)
	}

	defer f.Close()
	io.Copy(w, f)
}

// Getters
func handlerHero(w http.ResponseWriter, r *http.Request) {
	functionS := "[handlerHero]"
	logger.DumpLog.Printf("%s call for API hadler\n", functionS)
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	hero := stypes.GetHeroByID(uint(id))

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

	team := stypes.GetTeamByID(uint(id))

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

	expeditions := stypes.GetExpeditions()

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

	user := stypes.GetUserByID(uint(id))

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
	user := &stypes.User{}
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
	stypes.UpdateUser(user)
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
	user := &stypes.User{}
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
	stypes.UpdateTeam(user)
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

	user := stypes.GetUserByID(uint(data.Uid))

	if err != nil {
		logger.ErrLog.Printf("%s Error when trying to get the time : %s", functionS, r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := user.ResolveExpeditionState(&t)
	stypes.UpdateGameState(user.State)

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

	user := stypes.GetUserByID(uint(id))

	user.LaunchExpedition(expIdentifier)
	b, err := json.Marshal(user.State.CurrentExpedition.WhatHappened[0])
	if err != nil {
		logger.ErrLog.Println(err)
		return
	}
	logger.DumpLog.Printf("%s Giving :\n %s\n", functionS, string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func main() {
	stypes.DBLogIn()
	// Show main page
	fs := http.FileServer(http.Dir("./ui/wvtr-front/dist"))
	http.Handle("/", fs)

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
