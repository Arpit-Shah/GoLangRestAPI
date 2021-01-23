package main

import (
	"errors"
	"fmt"
	"strings"
)

// IRobots slices
var IRobots []IRobot

// RobotIDCounter uint64
var RobotIDCounter uint64

// RobotTaskCounter uint64
var RobotTaskCounter uint64

// Robot interface
type Robot interface {
	EnqueueTask(commands string) (taskID string, position chan RobotState, err chan error)

	CancelTask(taskID string) error

	CurrentState() RobotState
}

// IRobot Structure
type IRobot struct {
	ID         uint64
	RobotState RobotState
}

// RobotState Model
type RobotState struct {
	X        uint64
	Y        uint64
	HasCrate bool
}

// EnqueueTask method for IROBOT
func (ir *IRobot) EnqueueTask(commands string) (taskID string, position chan RobotState, err chan error) {

	fmt.Println("Processing Robot ID: ", ir.ID)

	// Split strings using white space
	actions := strings.Fields(commands)

	for i := 0; i < len(actions); i++ {
		if actions[i] == "N" {
			if ir.RobotState.Y != 9 {
				ir.RobotState.Y++
				fmt.Println("Travelling North ", ir.RobotState.Y)
				//StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move North successful"})
			} else {
				err <- errors.New("Can not move to North")
			}
		} else if actions[i] == "S" {
			if ir.RobotState.Y != 0 {
				ir.RobotState.Y--
				fmt.Println("Travelling South")
				//StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move South successful"})
			} else {
				err <- errors.New("Can not move to South")
			}
		} else if actions[i] == "E" {
			if ir.RobotState.X != 9 {
				ir.RobotState.X++
				fmt.Println("Travelling East")
				//StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move East successful"})
			} else {
				err <- errors.New("Can not move to East")
			}
		} else if actions[i] == "W" {
			if ir.RobotState.X != 0 {
				ir.RobotState.X--
				fmt.Println("Travelling West")
				//StatusEvents = append(StatusEvents, mStatusEvent{Error: false, Message: "Move West successful"})
			} else {
				err <- errors.New("Can not move to West")
			}
		}
		fmt.Println("Sending position")
		//position <- ir.RobotState

	}
	incrementRobotTaskCounter()
	return fmt.Sprint(RobotTaskCounter), position, err

}

// CancelTask method for IROBOT
func (ir *IRobot) CancelTask(taskID string) error {
	return errors.New("Something went wrong")
}

// CurrentState method for IROBOT
func (ir *IRobot) CurrentState() RobotState {
	return ir.RobotState
}

func createRobot() IRobot {
	var robot IRobot
	robot.ID = RobotIDCounter
	incrementRobotIDCounter()
	return robot
}

func incrementRobotIDCounter() {
	RobotIDCounter++
}

func incrementRobotTaskCounter() {
	RobotTaskCounter++
}
