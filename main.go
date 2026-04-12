package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"wvtrserv/dcole"
)

// import "wvtrserv/stypes"

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
    io.Copy(w, f);
}

func handlerWaifu(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("API Waifu req\n");
    id, err := strconv.Atoi(r.URL.Path[len("/waifus/"):])
    fmt.Printf("id: %s", id)

    waifu := dcole.GetWaifuByID(id)
    
    b, err := json.Marshal(waifu)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
    fmt.Fprintf(w, "%s", string(b))
}

func handlerTeam(w http.ResponseWriter, r *http.Request) {
    
    // ids, err := r.URL.Path[len("/waifus/"):]
    // id := strconv.Atoi(ids)
    ids := r.PathValue("id")
    id,_ := strconv.Atoi(ids)
    fmt.Printf("API Team req : %d\n", ids);

    team := dcole.GetTeamByID(id)

    b, err := json.Marshal(team)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
    fmt.Fprintf(w, "%s", string(b))
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
    ids := r.PathValue("id")
    id,_ := strconv.Atoi(ids)
    fmt.Printf("API Team req : %d\n", ids);

    user := dcole.GetUserByID(id)

    b, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
    fmt.Fprintf(w, "%s", string(b))
}

func main() {
    fs := http.FileServer(http.Dir("./ui/wvtr-front/dist"))
    http.Handle("/", fs)

    //waifusRouter := http.NewServeMux()

    //waifusRouter.HandleFunc("GET /waifus/{id}", handlerWaifu)
    http.HandleFunc("/waifus/", handlerWaifu)
    http.HandleFunc("/teams/{id}", handlerTeam)
    //http.Handle("/imgs/", handlerImages)
    http.Handle("/imgs/", http.StripPrefix("/imgs/", http.FileServer(http.Dir("imgs/"))))
    http.HandleFunc("/user/{id}", handlerUser) // Return the state of the user for the front to know what to display

    log.Println("Listening on :4210...")
    err := http.ListenAndServe(":4210", nil)
    if err != nil {
        log.Fatal(err)
    }
}
