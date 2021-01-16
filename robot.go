package main

// Robots Fake Database slice
var Robots []Robot

// RobotIDCounter uint64
var RobotIDCounter uint64

// Robot Model
type Robot struct {
	ID    uint64
	Jobs  []Job
	State RobotState
}

// RobotState Model
type RobotState struct {
	X        uint64 `json:"x"`
	Y        uint64 `json:"y"`
	HasCrate bool   `json:"hascrate"`
}

func createRobot() Robot {
	var robot Robot
	robot.ID = RobotIDCounter
	incrementRobotIDCounter()
	return robot
}

func incrementRobotIDCounter() {
	RobotIDCounter++
}
