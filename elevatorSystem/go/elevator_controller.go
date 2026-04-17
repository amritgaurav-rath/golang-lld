package main

import (
	"fmt"
	"math"
)

// ElevatorController manages multiple elevators and handles grouping
type ElevatorController struct {
	Elevators []*Elevator
}

// NewElevatorController seeds a fixed array of Elevators globally processing
func NewElevatorController(numElevators int, capacity int) *ElevatorController {
	elevators := make([]*Elevator, numElevators)
	for i := 0; i < numElevators; i++ {
		elevators[i] = NewElevator(fmt.Sprintf("E%d", i+1), capacity)
	}
	return &ElevatorController{
		Elevators: elevators,
	}
}

// RequestElevator scans available objects to find the optimal assignment 
func (ec *ElevatorController) RequestElevator(req *Request) error {
	var optimalElevator *Elevator
	minDistance := math.MaxInt32

	for _, el := range ec.Elevators {
		dist := el.GetDistanceIfAssigned(req)
		
		// Prioritizes shortest distance and direction
		if dist < minDistance {
			minDistance = dist
			optimalElevator = el
		}
	}

	// Exhausted memory assignment mapping
	if optimalElevator == nil {
		return fmt.Errorf("all elevators are currently heavily over capacity to handle %d passengers", req.Passengers)
	}

	fmt.Printf("[Dispatcher] 🟢 Routing (Src: %d -> Dest: %d, Pax: %d) into Optimal => %s (Weight %d)\n", req.SourceFloor, req.DestinationFloor, req.Passengers, optimalElevator.ID, minDistance)
	return optimalElevator.AddRequest(req)
}
