package main

// Jobs Fake Database slice
var Jobs []Job

// JobIDCounter uint64
var JobIDCounter uint64

// Job Model
type Job struct {
	ID      uint64
	Command string
	Active  bool
	Error   bool
	Message string
}

func incrementJobIDCounter() {
	JobIDCounter++
}
