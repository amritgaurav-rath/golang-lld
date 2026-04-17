package main

import (
	"fmt"
)

// ElevatorController manages a fleet of elevators and dispatches incoming requests
type ElevatorController struct {
	Elevators []*Elevator // array of all managed elevators
}

// NewElevatorController initializes the controller with a specified number of elevators and capacity
func NewElevatorController(numElevators int, capacity int) *ElevatorController {
	elevators := make([]*Elevator, numElevators)
	for i := 0; i < numElevators; i++ {
		elevators[i] = NewElevator(fmt.Sprintf("E%d", i+1), capacity)
	}
	return &ElevatorController{
		Elevators: elevators,
	}
}

// RequestElevator assigns a passenger request to the most optimal elevator based on distance and capacity
func (ec *ElevatorController) RequestElevator(req *Request) error {
	var optimalElevator *Elevator
	minDistance := int(1e9)

	// Loop through all elevators to find the closest one capable of handling the request
	for _, el := range ec.Elevators {
		dist := el.GetDistanceIfAssigned(req)
		if dist < minDistance {
			minDistance = dist
			optimalElevator = el
		}
	}

	if optimalElevator == nil {
		return fmt.Errorf("all elevators are currently over capacity to handle %d passengers", req.Passengers)
	}

	fmt.Printf("[Dispatcher] 🟢 Routing (Src: %d -> Dest: %d, Pax: %d) into Optimal => %s\n", req.SourceFloor, req.DestinationFloor, req.Passengers, optimalElevator.ID)
	return optimalElevator.AddRequest(req) // Assign the request to the chosen elevator
}

// TickAll advances the physical simulation by one unit of time for all managed elevators
// Returns true if at least one elevator is still active (moving or processing stops)
func (ec *ElevatorController) TickAll() bool {
	isActive := false
	for _, el := range ec.Elevators {
		el.Tick()
		if el.CurrentDirection != IDLE {
			isActive = true
		}
	}
	return isActive
}
