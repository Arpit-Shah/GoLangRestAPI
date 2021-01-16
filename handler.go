package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Home Page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!"+
		"\nFollowing end points can be used"+
		"\nlocalhost:8080/robots : will return all robots information"+
		"\nlocalhost:8080/create : will create new robot with (0,0) location and return created robot ID"+
		"\nlocalhost:8080/delete/1 : will delete robot with id=1 (if present)"+
		"\nlocalhost:8080/state/1 : will return json object with current state of robot with id=1 (if present)"+
		"\nlocalhost:8080/move/1/N N E : will move robot with id=1 two steps in north direction and one step in E direction. Robot will be in (1, 2) position and return jobID"
		"\nlocalhost:8080/status/10 : will return job status of the job with jobID=10"
		"\nlocalhost:8080/cancel/10 : will cancel job with jobID=10 if not in process")
}

func createRobotHandler(w http.ResponseWriter, r *http.Request) {
	var robot = createRobot()
	Robots = append(Robots, robot)
	fmt.Fprintf(w, "The robot with ID %v has been created successfully", robot.ID)
}

func getRobotsHandler(w http.ResponseWriter, r *http.Request) {
	var mRobots []mRobot
	for i := 0; i < len(Robots); i++ {

		// Create JSON model
		var modelRobotState mRobotState
		modelRobotState.X = Robots[i].State.X
		modelRobotState.Y = Robots[i].State.Y
		modelRobotState.HasCrate = Robots[i].State.HasCrate
		var modelRobot mRobot
		modelRobot.ID = Robots[i].ID
		modelRobot.RobotState = &modelRobotState

		// Add each model to slice
		mRobots = append(mRobots, modelRobot)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mRobots)
}

// Delete Robot
func deleteRobotHandler(w http.ResponseWriter, r *http.Request) {

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
func getRobotStateHandler(w http.ResponseWriter, r *http.Request) {

	robotID := mux.Vars(r)["id"]
	// Convert to uInt64
	robotIDuint, _ := strconv.ParseUint(robotID, 10, 64)

	for i := 0; i < len(Robots); i++ {

		if Robots[i].ID == robotIDuint {
			// Create JSON Model
			var modelRobotState mRobotState
			modelRobotState.X = Robots[i].State.X
			modelRobotState.Y = Robots[i].State.Y
			modelRobotState.HasCrate = Robots[i].State.HasCrate

			// Return created Robot
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(modelRobotState)
		}
	}
}

func jobStatusHandler(w http.ResponseWriter, r *http.Request) {
	robotID := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Under Development %v", robotID)
}

func jobCancelHandler(w http.ResponseWriter, r *http.Request) {
	robotID := mux.Vars(r)["id"]
	fmt.Fprintf(w, "Under Development %v", robotID)
}

// Move Robot
func moveRobotHandler(w http.ResponseWriter, r *http.Request) {

	robotID := mux.Vars(r)["id"]
	command := mux.Vars(r)["command"]
	// fmt.Println(robotID)
	// fmt.Println(command)

	if !validateCommand(command) {
		http.Error(w, "BAD Request", http.StatusBadRequest)
		return
	}

	// Convert to uInt64
	robotIDuint, _ := strconv.ParseUint(robotID, 10, 64)

	for j := 0; j < len(Robots); j++ {

		fmt.Println(Robots[j].ID)
		if Robots[j].ID == robotIDuint {
			// var StatusEvents = processJob(&singleRobot.State, command)

			// create a slice for the errors
			var StatusEvents []mStatusEvent

			// Split strings using white space
			actions := strings.Fields(command)

			for i := 0; i < len(actions); i++ {
				if actions[i] == "N" {
					if Robots[j].State.Y != 9 {
						Robots[j].State.Y++
						fmt.Println("Travelling North")
						StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move North successful"})
					} else {
						StatusEvents = append(StatusEvents, mStatusEvent{Error: true, Message: "Can not move to North"})
					}
				} else if actions[i] == "S" {
					if Robots[j].State.Y != 0 {
						Robots[j].State.Y--
						fmt.Println("Travelling South")
						StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move South successful"})
					} else {
						StatusEvents = append(StatusEvents, mStatusEvent{Error: true, Message: "Can not move to South"})
					}
				} else if actions[i] == "E" {
					if Robots[j].State.X != 9 {
						fmt.Println("I am here")
						fmt.Print(Robots[i].State.X)
						Robots[j].State.X++
						fmt.Print(Robots[i].State.X)
						fmt.Println("Travelling East")
						StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move East successful"})
					} else {
						StatusEvents = append(StatusEvents, mStatusEvent{Error: true, Message: "Can not move to East"})
					}
				} else if actions[i] == "W" {
					if Robots[j].State.X != 0 {
						Robots[j].State.X--
						fmt.Println("Travelling West")
						StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move West successful"})
					} else {
						StatusEvents = append(StatusEvents, mStatusEvent{Error: true, Message: "Can not move to West"})
					}
				}

			}

			// Return Status Event
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(StatusEvents)
			break
		}
	}

}

func processJob(singleRobotState *RobotState, command string) []mStatusEvent {

	// create a slice for the errors
	var StatusEvents []mStatusEvent

	// Split strings using white space
	actions := strings.Fields(command)

	for i := 0; i < len(actions); i++ {
		if actions[i] == "N" {
			if singleRobotState.Y != 9 {
				singleRobotState.Y++
				fmt.Println("Travelling North")
				StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move North successful"})
			} else {
				StatusEvents = append(StatusEvents, mStatusEvent{Error: true, Message: "Can not move to North"})
			}
		}
		if actions[i] == "S" {
			if singleRobotState.Y != 0 {
				singleRobotState.Y--
				fmt.Println("Travelling South")
				StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move South successful"})
			} else {
				StatusEvents = append(StatusEvents, mStatusEvent{Error: true, Message: "Can not move to South"})
			}
		}
		if actions[i] == "E" {
			if singleRobotState.X != 9 {
				singleRobotState.X++
				fmt.Println("Travelling East")
				StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move East successful"})
			} else {
				StatusEvents = append(StatusEvents, mStatusEvent{Error: true, Message: "Can not move to East"})
			}
		}
		if actions[i] == "W" {
			if singleRobotState.X != 0 {
				singleRobotState.X--
				fmt.Println("Travelling East")
				StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move West successful"})
			} else {
				StatusEvents = append(StatusEvents, mStatusEvent{Error: true, Message: "Can not move to West"})
			}
		}

	}

	fmt.Print(singleRobotState.X)
	fmt.Print(singleRobotState.Y)
	fmt.Println()

	return StatusEvents
}

func createJob(command string) Job {
	var job Job
	job.ID = JobIDCounter
	incrementJobIDCounter()
	job.Command = command
	job.Active = false
	job.Error = false
	job.Message = ""

	return job
}

func validateCommand(command string) bool {
	// Split strings using white space
	actions := strings.Fields(command)

	// Validate if its valid command
	for i := 0; i < len(actions); i++ {
		if len(actions[i]) != 1 {
			return false
		}
		if actions[i] != "N" && actions[i] != "E" && actions[i] != "S" && actions[i] != "W" {
			return false
		}
		// fmt.Println(actions[i])
	}

	return true
}
