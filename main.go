package main

import (
	"github.com/gorilla/mux"
	"github.com/leandrolls/JsonWithGo/pkg/controllers"
	"log"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/contact", controllers.ContactHandler).Methods("GET")
	r.HandleFunc("/about", controllers.AboutHandler).Methods("GET")
	r.HandleFunc("/times", controllers.GetTimesHandler).Methods("GET")
	r.HandleFunc("/time/{id}", controllers.GetTimeHandler).Methods("GET")
	r.HandleFunc("/", controllers.IndexHandler).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}


