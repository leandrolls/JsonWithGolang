package controllers

import (
	"fmt"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is the about page!")
}
