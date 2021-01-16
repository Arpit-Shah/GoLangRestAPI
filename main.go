package main

// create main.go
// Install Go plugin
// Install router from "go get -u github.com/gorilla/mux"

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Robots Fake Database slice
var Robots []Robot

// IDCounter counter
var IDCounter uint64

// Robot Model
type Robot struct {
	ID         uint64      `json:"id"`
	RobotState *RobotState `json:"robotstate"`
}

// RobotState Model
type RobotState struct {
	X        uint64 `json:"x"`
	Y        uint64 `json:"y"`
	HasCrate bool   `json:"hascrate"`
}

// StatusEvent for errors or success messages
type StatusEvent struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

/*
	POST: localhost:8080/create will create robot with (X,Y) location to (0,0)
	POST: localhost:8080/create?X=1&Y=1 will create robot with (X,Y) location to (1, 1)
*/
func createRobot(w http.ResponseWriter, r *http.Request) {

	var xPosuint uint64
	var yPosuint uint64

	// Returns a url.Values, which is a map[string][]string
	vals := r.URL.Query()

	// Get Query Parameters
	xPos, ok := vals["X"]
	if ok {
		// Convert query parameters to uInt64
		xPosuint, _ = strconv.ParseUint(xPos[0], 10, 64)
		if xPosuint > 10 {
			xPosuint = 10
		}
	}
	yPos, ok := vals["Y"]
	if ok {
		// Convert query parameters to uInt64
		yPosuint, _ = strconv.ParseUint(yPos[0], 10, 64)
		if yPosuint > 10 {
			yPosuint = 10
		}
	}

	// Create RobotState
	state := new(RobotState)
	state.X = xPosuint
	state.Y = yPosuint
	state.HasCrate = false

	// Create Robot
	robot := new(Robot)
	robot.ID = IDCounter
	IDCounter++
	robot.RobotState = state

	// Add Robot to databse
	Robots = append(Robots, *robot)

	// Return created Robot
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(Robots)

	fmt.Fprintf(w, "Created Robot with ID %v", robot.ID)
}

// Get Robot Info
func getRobots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Robots)
}

// Delete Robot
func deleteRobot(w http.ResponseWriter, r *http.Request) {

	robotID := mux.Vars(r)["id"]
	// Convert to uInt64
	robotIDuint, _ := strconv.ParseUint(robotID, 10, 64)

	for i, singleRobot := range Robots {
		if singleRobot.ID == robotIDuint {
			Robots = append(Robots[:i], Robots[i+1:]...)
			fmt.Fprintf(w, "The robot with ID %v has been deleted successfully", robotID)
		}
	}
}

// Get Robot Current State
func getRobotState(w http.ResponseWriter, r *http.Request) {

	robotID := mux.Vars(r)["id"]
	// Convert to uInt64
	robotIDuint, _ := strconv.ParseUint(robotID, 10, 64)

	for i, singleRobot := range Robots {
		if singleRobot.ID == robotIDuint {
			singleRobotState := Robots[i].RobotState
			// Return created Robot
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(singleRobotState)
		}
	}
}

// Move Robot
func moveRobot(w http.ResponseWriter, r *http.Request) {

	robotID := mux.Vars(r)["id"]
	command := mux.Vars(r)["command"]
	fmt.Println(robotID)
	fmt.Println(command)

	// Split strings using white space
	actions := strings.Fields(command)

	// Validate if its valid command
	for i := 0; i < len(actions); i++ {
		if len(actions[i]) != 1 {
			fmt.Fprintf(w, "Invalid Command Length: %v", actions[i])
		}
		if actions[i] != "N" && actions[i] != "E" && actions[i] != "S" && actions[i] != "W" {
			fmt.Fprintf(w, "Invalid Command: %v", actions[i])
		}
		fmt.Println(actions[i])
	}

	// Convert to uInt64
	robotIDuint, _ := strconv.ParseUint(robotID, 10, 64)

	for i, singleRobot := range Robots {

		// create a slice for the errors
		var StatusEvents []StatusEvent

		if singleRobot.ID == robotIDuint {
			singleRobotState := Robots[i].RobotState

			for i := 0; i < len(actions); i++ {
				if actions[i] == "N" {
					if singleRobotState.Y != 9 {
						singleRobotState.Y++
						fmt.Println("Travelling North")
						StatusEvents = append(StatusEvents, StatusEvent{Error: false, Message: "Move North successful"})
					} else {
						StatusEvents = append(StatusEvents, StatusEvent{Error: true, Message: "Can not move to North"})
					}
				}
				if actions[i] == "S" {
					if singleRobotState.Y != 0 {
						singleRobotState.Y--
						fmt.Println("Travelling South")
						StatusEvents = append(StatusEvents, StatusEvent{Error: false, Message: "Move South successful"})
					} else {
						StatusEvents = append(StatusEvents, StatusEvent{Error: true, Message: "Can not move to South"})
					}
				}
				if actions[i] == "E" {
					if singleRobotState.X != 9 {
						singleRobotState.X++
						fmt.Println("Travelling East")
						StatusEvents = append(StatusEvents, StatusEvent{Error: false, Message: "Move East successful"})
					} else {
						StatusEvents = append(StatusEvents, StatusEvent{Error: true, Message: "Can not move to East"})
					}
				}
				if actions[i] == "W" {
					if singleRobotState.X != 0 {
						singleRobotState.X--
						fmt.Println("Travelling East")
						StatusEvents = append(StatusEvents, StatusEvent{Error: false, Message: "Move West successful"})
					} else {
						StatusEvents = append(StatusEvents, StatusEvent{Error: true, Message: "Can not move to West"})
					}
				}

			}

			// Return Status Event
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(StatusEvents)

		}
	}
}

// Home Page
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!"+
		"\nFollowing end points can be used"+
		"\nlocalhost:8080/robots : will return all robots information"+
		"\nlocalhost:8080/create : will create new robot with (0,0) location and return created robot ID"+
		"\nlocalhost:8080/create?X=1&Y=1 : will create new robot with (1,1) location and return all robots information"+
		"\nlocalhost:8080/delete/1 : will delete robot with id=1 (if present)"+
		"\nlocalhost:8080/state/1 : will return json object with current state of robot with id=1 (if present)"+
		"\nlocalhost:8080/move/1/N N E : will move robot with id=1 two steps in north direction and one step in E direction. Robot will be in (1, 2) position")
}

func main() {
	fmt.Println("Welcome to Robot SDK")

	// Initialise router
	r := mux.NewRouter().StrictSlash(true)

	// Route Handlers / Endpoints
	r.HandleFunc("/", homeLink)
	r.HandleFunc("/robots", getRobots).Methods("GET")
	r.HandleFunc("/create", createRobot).Methods("POST")
	r.HandleFunc("/delete/{id}", deleteRobot).Methods("POST")
	r.HandleFunc("/state/{id}", getRobotState).Methods("GET")
	r.HandleFunc("/move/{id}/{command}", moveRobot).Methods("POST")

	// Start Server
	log.Fatal(http.ListenAndServe(":8080", r))
}
