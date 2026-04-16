package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	"wvtrserv/stypes"
)

// Main page
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("req.Method: %s\n", r.Method)
	fmt.Printf("req.URL.Path: %s\n", r.URL.Path)
	fmt.Printf("req.ContentLength: %d\n", r.ContentLength)

	d := http.Dir("./ui/vu/UI/dist")
	f, err := d.Open("index.html")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	io.Copy(w, f)
}

// Getters
func handlerHero(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("call for API hadler hero\n")
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	hero := stypes.GetHeroByID(uint(id))

	b, err := json.Marshal(hero)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Giving : %s\n", string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("call for API hadler team\n")
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	team := stypes.GetTeamByID(uint(id))

	b, err := json.Marshal(team)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Giving : %s\n", string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerAvailableExpeditions(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("call for API hadler Available Expeditions\n")

	expeditions := stypes.GetExpeditions()

	b, err := json.Marshal(expeditions)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Giving : %s\n", string(b))
	fmt.Fprintf(w, "%s", string(b))
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("call for API hadler user\n")
	ids := r.PathValue("id")
	id, _ := strconv.Atoi(ids)

	user := stypes.GetUserByID(uint(id))

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Giving : %s\n", string(b))
	fmt.Fprintf(w, "%s", string(b))
}

// Updaters
func handlerSaveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("call for API hadler Save User\n")
	if r.Method != http.MethodPost {
		s := fmt.Sprintf("Method not allowed (%s) POST expected.", r.Method)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}
	user := &stypes.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Printf("Error when trying to get the user from the request body, got : %s", r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Giving : %s\n", string(b))
	stypes.UpdateUser(user)
	w.WriteHeader(http.StatusCreated)
}

func handlerUpdateTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("call for API hadler Update Team\n")
	if r.Method != http.MethodPost {
		s := fmt.Sprintf("Method not allowed (%s) POST expected.", r.Method)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}
	user := &stypes.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Printf("Error when trying to get the user from the request body, got : %s", r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Giving : %s\n", string(b))
	stypes.UpdateTeam(user)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{}")
}

type CurrentStepRequestMessage struct {
	Uid  int   `json:"id"`
	Time int64 `json:"time"`
}

func handlerCurrentExpeditionStep(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("call for API Current Expedition Step\n")

	fmt.Printf("method check\n")
	if r.Method != http.MethodPost {
		s := fmt.Sprintf("Method not allowed (%s) POST expected.", r.Method)
		fmt.Println(s)
		http.Error(w, s, http.StatusMethodNotAllowed)
		return
	}
	var data CurrentStepRequestMessage
	err := json.NewDecoder(r.Body).Decode(&data)
	fmt.Printf("time ? %d\n", data.Time)
	var t time.Time = time.Unix(0, data.Time*int64(time.Millisecond))

	user := stypes.GetUserByID(uint(data.Uid))

	fmt.Printf("decode check\n")
	if err != nil {
		fmt.Printf("Error when trying to get the time : %s", r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("resovle expedition\n")
	res := user.ResolveExpeditionState(&t)
	fmt.Printf("update game state\n")
	stypes.UpdateGameState(user.State)

	resS := "{}"
	if res != nil {
		fmt.Printf("marshal du res\n")
		b, err := json.Marshal(res)

		fmt.Printf("encode check\n")
		if err != nil {
			fmt.Println(err)
			return
		}
		resS = string(b)
	}

	fmt.Printf("Giving: %s\n", resS)
	fmt.Fprintf(w, "%s", resS)
}

func handlerLaunchExpedition(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("call for API hadler Launch expedition\n")
	ids := r.PathValue("usr")
	id, _ := strconv.Atoi(ids)

	expIdentifier := r.PathValue("expId")

	user := stypes.GetUserByID(uint(id))
	fmt.Printf("user Ok\n")

	user.LaunchExpedition(expIdentifier)
	fmt.Printf("after launch\n")
	b, err := json.Marshal(user.State.CurrentExpedition.WhatHappened[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Giving: %s\n", string(b))
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

	log.Println("Listening on :4210...")
	err := http.ListenAndServe(":4210", nil)
	if err != nil {
		log.Fatal(err)
	}
}
