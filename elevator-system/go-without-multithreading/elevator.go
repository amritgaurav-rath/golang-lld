package main

import (
	"fmt"
	"math"
)

type Elevator struct {
	ID               string
	CurrentFloor     int
	CurrentDirection Direction
	Capacity         int
	CurrentLoad      int

	UpStops   map[int]bool
	DownStops map[int]bool

	Requests []*Request
}

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

func (e *Elevator) GetDistanceIfAssigned(req *Request) int {
	if e.CurrentLoad+req.Passengers > e.Capacity {
		return math.MaxInt32
	}

	dist := int(math.Abs(float64(e.CurrentFloor - req.SourceFloor)))
	
	if e.CurrentDirection == IDLE {
		return dist
	}
	if e.CurrentDirection == UP && req.SourceFloor >= e.CurrentFloor {
		return dist
	}
	if e.CurrentDirection == DOWN && req.SourceFloor <= e.CurrentFloor {
		return dist
	}

	return dist + 1000 
}

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
