package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Create 1 Robot by Default
	Robots = append(Robots, createRobot())

	// Initialise router
	r := mux.NewRouter().StrictSlash(true)

	// Route Handlers / Endpoints
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/robots", getRobotsHandler).Methods("GET")
	r.HandleFunc("/create", createRobotHandler).Methods("POST")
	r.HandleFunc("/delete/{id}", deleteRobotHandler).Methods("POST")
	r.HandleFunc("/state/{id}", getRobotStateHandler).Methods("GET")
	r.HandleFunc("/move/{id}/{command}", moveRobotHandler).Methods("POST")
	r.HandleFunc("/status/{id}/{jobid}", jobStatusHandler).Methods("GET")
	r.HandleFunc("/cancel/{id}/{jobid}", jobCancelHandler).Methods("GET")

	// Start Server
	log.Fatal(http.ListenAndServe(":8080", r))
}
