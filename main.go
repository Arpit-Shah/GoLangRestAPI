package main

// create main.go
// Install Go plugin
// Install router from "go get -u github.com/gorilla/mux"

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var robots []Robot

// Models
type Robot struct {
	TaskID     string      `json:"taskid"`
	RobotState *RobotState `json:"robotstate"`
}

// Models
type RobotState struct {
	X        uint `json:"x"`
	Y        uint `json:"y"`
	HasCrate bool `json:"hascrate"`
}

// Get Robot Info
func getRobots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(robots)
}

func main() {
	fmt.Println("Hello World")

	// Initialise router
	r := mux.NewRouter()

	robots = append(robots, Robot{TaskID: "1", RobotState: &RobotState{X: 0, Y: 0, HasCrate: true}})
	// Route Handlers / Endpoints
	r.HandleFunc("/api/wh/robot", getRobots).Methods("GET")

	// Start Server
	log.Fatal(http.ListenAndServe(":8080", r))
}
