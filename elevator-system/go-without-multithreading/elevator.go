package main

import (
	"fmt"
)

// Elevator represents a single physical elevator in the system
type Elevator struct {
	ID               string
	CurrentFloor     int
	CurrentDirection Direction
	Capacity         int
	CurrentLoad      int

	UpStops   map[int]bool // Floors where the elevator needs to stop while moving UP
	DownStops map[int]bool // Floors where the elevator needs to stop while moving DOWN

	Requests []*Request // List of currently assigned requests
}

// NewElevator creates a new elevator initialized to the ground floor (0) and state IDLE
func NewElevator(id string, capacity int) *Elevator {
	return &Elevator{
		ID:               id,
		CurrentFloor:     0,
		CurrentDirection: IDLE,
		Capacity:         capacity,
		CurrentLoad:      0,
		UpStops:          make(map[int]bool),
		DownStops:        make(map[int]bool),
		Requests:         make([]*Request, 0),
	}
}

// AddRequest assigns a new passenger request to this elevator and updates its planned stops
func (e *Elevator) AddRequest(req *Request) error {
	if e.CurrentLoad+req.Passengers > e.Capacity {
		return fmt.Errorf("elevator %s capacity exceeded", e.ID)
	}

	e.Requests = append(e.Requests, req)
	e.CurrentLoad += req.Passengers

	if req.SourceFloor > e.CurrentFloor {
		e.UpStops[req.SourceFloor] = true
	} else if req.SourceFloor < e.CurrentFloor {
		e.DownStops[req.SourceFloor] = true
	} else {
		if req.DestinationFloor > e.CurrentFloor {
			e.UpStops[req.DestinationFloor] = true
		} else {
			e.DownStops[req.DestinationFloor] = true
		}
	}

	if req.DestinationFloor > req.SourceFloor {
		e.UpStops[req.DestinationFloor] = true
	} else if req.DestinationFloor < req.SourceFloor {
		e.DownStops[req.DestinationFloor] = true
	}

	if e.CurrentDirection == IDLE {
		if req.SourceFloor > e.CurrentFloor || req.DestinationFloor > e.CurrentFloor {
			e.CurrentDirection = UP
		} else if req.SourceFloor < e.CurrentFloor || req.DestinationFloor < e.CurrentFloor {
			e.CurrentDirection = DOWN
		}
	}

	return nil
}

// GetDistanceIfAssigned calculates a heuristic "distance" cost for this elevator to handle the given request
// Returns int(1e9) if the elevator is currently too full
func (e *Elevator) GetDistanceIfAssigned(req *Request) int {
	if e.CurrentLoad+req.Passengers > e.Capacity {
		return int(1e9)
	}

	dist := e.CurrentFloor - req.SourceFloor
	if dist < 0 {
		dist = -dist
	}
	
	if e.CurrentDirection == IDLE {
		return dist
	}
	if e.CurrentDirection == UP && req.SourceFloor >= e.CurrentFloor {
		return dist
	}
	if e.CurrentDirection == DOWN && req.SourceFloor <= e.CurrentFloor {
		return dist
	}

	// Penalizes the elevator with a higher cost if moving in the opposite direction
	return dist + 1000 
}

// processCurrentFloor handles stopping and dropping off passengers at the current floor
func (e *Elevator) processCurrentFloor() {
	if e.CurrentDirection == UP {
		if e.UpStops[e.CurrentFloor] {
			fmt.Printf("   -> Elevator %s 🛑 STOPPED at floor %d (UP)\n", e.ID, e.CurrentFloor)
			delete(e.UpStops, e.CurrentFloor)
			e.dropOffPassengers()
		}
	} else if e.CurrentDirection == DOWN {
		if e.DownStops[e.CurrentFloor] {
			fmt.Printf("   -> Elevator %s 🛑 STOPPED at floor %d (DOWN)\n", e.ID, e.CurrentFloor)
			delete(e.DownStops, e.CurrentFloor)
			e.dropOffPassengers()
		}
	}
}

// dropOffPassengers removes completed requests and updates the elevator's current load
func (e *Elevator) dropOffPassengers() {
	var remaining []*Request
	for _, req := range e.Requests {
		if req.DestinationFloor == e.CurrentFloor {
			e.CurrentLoad -= req.Passengers
			fmt.Printf("      -> %d passengers ALIGHTED at floor %d from elevator %s\n", req.Passengers, e.CurrentFloor, e.ID)
		} else {
			remaining = append(remaining, req)
		}
	}
	e.Requests = remaining
}

// getNextDirection determines whether the elevator should continue in its current direction, reverse, or turn IDLE
func (e *Elevator) getNextDirection() Direction {
	if e.CurrentDirection == UP {
		for floor := range e.UpStops {
			if floor > e.CurrentFloor {
				return UP
			}
		}
		if len(e.DownStops) > 0 {
			return DOWN
		}
	} else if e.CurrentDirection == DOWN {
		for floor := range e.DownStops {
			if floor < e.CurrentFloor {
				return DOWN
			}
		}
		if len(e.UpStops) > 0 {
			return UP
		}
	}

	if len(e.UpStops) > 0 {
		return UP
	} else if len(e.DownStops) > 0 {
		return DOWN
	}

	return IDLE
}

// Tick manually steps the exact sequence once per cycle without background threads
// It evaluates direction, moves by one floor, and processes any stops at the new floor
func (e *Elevator) Tick() {
	nextDir := e.getNextDirection()
	e.CurrentDirection = nextDir

	if e.CurrentDirection == IDLE {
		return
	}

	if e.CurrentDirection == UP {
		e.CurrentFloor++
	} else if e.CurrentDirection == DOWN {
		e.CurrentFloor--
	}

	e.processCurrentFloor()
}
