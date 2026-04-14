package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Elevator represents an individual elevator in the system
type Elevator struct {
	ID               string
	CurrentFloor     int
	CurrentDirection Direction
	Capacity         int
	CurrentLoad      int

	// Thread safety for adding/removing requests concurrently
	mu sync.RWMutex

	// Pending stops for the elevator (using maps as native Sets to ensure unique stops)
	UpStops   map[int]bool
	DownStops map[int]bool

	// Dedicated requests tracking passenger boarding/alighting
	Requests []*Request
}

func NewElevator(id string, capacity int) *Elevator {
	el := &Elevator{
		ID:               id,
		CurrentFloor:     0,
		CurrentDirection: IDLE,
		Capacity:         capacity,
		CurrentLoad:      0,
		UpStops:          make(map[int]bool),
		DownStops:        make(map[int]bool),
		Requests:         make([]*Request, 0),
	}
	// Start spinning the elevator's internal service loop concurrently
	go el.run()
	return el
}

func (e *Elevator) AddRequest(req *Request) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.CurrentLoad+req.Passengers > e.Capacity {
		return fmt.Errorf("elevator %s capacity exceeded", e.ID)
	}

	e.Requests = append(e.Requests, req)
	e.CurrentLoad += req.Passengers

	// Register source floor mathematically
	if req.SourceFloor > e.CurrentFloor {
		e.UpStops[req.SourceFloor] = true
	} else if req.SourceFloor < e.CurrentFloor {
		e.DownStops[req.SourceFloor] = true
	} else {
		// Currently on the exact floor, skip direct routing and immediately target destination
		if req.DestinationFloor > e.CurrentFloor {
			e.UpStops[req.DestinationFloor] = true
		} else {
			e.DownStops[req.DestinationFloor] = true
		}
	}

	// Register destination floor mathematically
	if req.DestinationFloor > req.SourceFloor {
		e.UpStops[req.DestinationFloor] = true
	} else if req.DestinationFloor < req.SourceFloor {
		e.DownStops[req.DestinationFloor] = true
	}

	// Wake up elevator direction if idle
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
	e.mu.RLock()
	defer e.mu.RUnlock()

	// If capacity is breached, immediately discard it mathematically
	if e.CurrentLoad+req.Passengers > e.Capacity {
		return math.MaxInt32
	}

	dist := int(math.Abs(float64(e.CurrentFloor - req.SourceFloor)))
	
	// Evaluate directional efficiency (Optimal matching algorithm)
	if e.CurrentDirection == IDLE {
		return dist
	}
	if e.CurrentDirection == UP && req.SourceFloor >= e.CurrentFloor {
		return dist
	}
	if e.CurrentDirection == DOWN && req.SourceFloor <= e.CurrentFloor {
		return dist
	}

	// Returning extremely high distance penalty if moving completely opposite to the request
	return dist + 1000 
}

func (e *Elevator) processCurrentFloor() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.CurrentDirection == UP {
		if e.UpStops[e.CurrentFloor] {
			fmt.Printf("Elevator %s 🛑 STOPPED at floor %d (UP)\n", e.ID, e.CurrentFloor)
			delete(e.UpStops, e.CurrentFloor)
			e.dropOffPassengers()
		}
	} else if e.CurrentDirection == DOWN {
		if e.DownStops[e.CurrentFloor] {
			fmt.Printf("Elevator %s 🛑 STOPPED at floor %d (DOWN)\n", e.ID, e.CurrentFloor)
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
			fmt.Printf("   -> %d passengers ALIGHTED at floor %d from elevator %s\n", req.Passengers, e.CurrentFloor, e.ID)
		} else {
			remaining = append(remaining, req)
		}
	}
	e.Requests = remaining
}

func (e *Elevator) getNextDirection() Direction {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if e.CurrentDirection == UP {
		// Keep going UP if there are stops above
		for floor := range e.UpStops {
			if floor > e.CurrentFloor {
				return UP
			}
		}
		// Otherwise switch downward bounds
		if len(e.DownStops) > 0 {
			return DOWN
		}
	} else if e.CurrentDirection == DOWN {
		// Keep going DOWN if there are stops below
		for floor := range e.DownStops {
			if floor < e.CurrentFloor {
				return DOWN
			}
		}
		// Otherwise switch upward bounds
		if len(e.UpStops) > 0 {
			return UP
		}
	}

	// IDLE recovery sweep
	if len(e.UpStops) > 0 {
		return UP
	} else if len(e.DownStops) > 0 {
		return DOWN
	}

	return IDLE
}

func (e *Elevator) run() {
	for {
		time.Sleep(100 * time.Millisecond) // Simulated travel time between floors

		nextDir := e.getNextDirection()

		e.mu.Lock()
		e.CurrentDirection = nextDir
		e.mu.Unlock()

		if e.CurrentDirection == IDLE {
			continue // Spin gracefully
		}

		if e.CurrentDirection == UP {
			e.mu.Lock()
			e.CurrentFloor++
			e.mu.Unlock()
		} else if e.CurrentDirection == DOWN {
			e.mu.Lock()
			e.CurrentFloor--
			e.mu.Unlock()
		}

		// Evaluate limits internally per floor change
		e.processCurrentFloor()
	}
}
