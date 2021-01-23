package main

// Warehouse interface
type Warehouse interface {
	Robots() []Robot
}

// RocosWarehouse Structure
type RocosWarehouse struct {
	ID       uint64
	Location string
}

// Robots return slice of Robot interfaces
func (rw *RocosWarehouse) Robots() []Robot {

	var Robots []Robot

	// Take current IRobot slice and cast it into Robot slice
	for _, Robot := range Robots {
		Robots = append(Robots, Robot)
	}

	return Robots
}
