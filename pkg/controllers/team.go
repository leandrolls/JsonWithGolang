package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leandrolls/JsonWithGo/pkg/db"
	"net/http"

)

func GetTimeHandler (w http.ResponseWriter, r *http.Request) {

	path := mux.Vars(r)

	t := path["id"]
	team := db.GetTeam(t)

	teamBytes, err := json.Marshal(team)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(teamBytes))


}

func GetTimesHandler (w http.ResponseWriter, r *http.Request) {

	team := db.GetTeams()

	teamBytes, err := json.Marshal(team)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(teamBytes))


}
