package main

// Direction indicates the current movement state of the elevator
type Direction int

const (
	IDLE Direction = iota // Elevator is stopped and has no active requests
	UP                   // Elevator is currently moving upwards
	DOWN                 // Elevator is currently moving downwards
)

func (d Direction) String() string {
	switch d {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	default:
		return "IDLE"
	}
}
