package controllers

import (
	"fmt"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is the contact page!")
}