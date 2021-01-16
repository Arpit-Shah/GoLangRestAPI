package main

// Robot Model
type mRobot struct {
	ID         uint64       `json:"id"`
	RobotState *mRobotState `json:"robotstate"`
}

// RobotState Model
type mRobotState struct {
	X        uint64 `json:"x"`
	Y        uint64 `json:"y"`
	HasCrate bool   `json:"hascrate"`
}

// StatusEvent for errors or success messages
type mStatusEvent struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
