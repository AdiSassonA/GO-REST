package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Init router
var router = mux.NewRouter()
var counter int

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func requests(w http.ResponseWriter, r *http.Request) {
	counter = counter + 1
	fmt.Fprintf(w, "This is requests %v", counter)
}

func drop(w http.ResponseWriter, r *http.Request) {
	counter = 0
	fmt.Fprintf(w, "This is drop %v", counter)
}

func update(w http.ResponseWriter, r *http.Request) {
	reqBody, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		fmt.Println(err)
		return
	}
	counter = reqBody
	fmt.Fprintf(w, "Counter is now %v", counter)
}

func handleRequests() {
	router.HandleFunc("/", homePage)
	router.HandleFunc("/requests", requests).Methods("GET")
	router.HandleFunc("/drop", drop).Methods("DELETE")
	router.HandleFunc("/update", update).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
