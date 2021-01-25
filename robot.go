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
	err = make(chan error)
	position = make(chan RobotState)

	go func() {
		// Split strings using white space
		actions := strings.Fields(commands)
		isFail := false

		for i := 0; i < len(actions); i++ {
			if actions[i] == "N" {
				if ir.RobotState.Y != 9 {
					ir.RobotState.Y++
					fmt.Println("Travelling North ", ir.RobotState.Y)
				} else {
					err <- errors.New("Can not move to North")
					isFail = true
					break
				}
			} else if actions[i] == "S" {
				if ir.RobotState.Y != 0 {
					ir.RobotState.Y--
					fmt.Println("Travelling South")
				} else {
					err <- errors.New("Can not move to South")
					isFail = true
					break
				}
			} else if actions[i] == "E" {
				if ir.RobotState.X != 9 {
					ir.RobotState.X++
					fmt.Println("Travelling East", ir.RobotState.Y)
				} else {
					err <- errors.New("Can not move to East")
					isFail = true
					break
				}
			} else if actions[i] == "W" {
				if ir.RobotState.X != 0 {
					ir.RobotState.X--
					fmt.Println("Travelling West")
				} else {
					err <- errors.New("Can not move to West")
					isFail = true
					break
				}
			}
			fmt.Println("Sending position")
			position <- ir.RobotState

		}

		if !isFail {
			err <- errors.New("OK")
		}
	}()
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
