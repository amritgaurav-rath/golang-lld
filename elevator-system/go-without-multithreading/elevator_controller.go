package main

import (
	"fmt"
	"math"
)

type ElevatorController struct {
	Elevators []*Elevator
}

func NewElevatorController(numElevators int, capacity int) *ElevatorController {
	elevators := make([]*Elevator, numElevators)
	for i := 0; i < numElevators; i++ {
		elevators[i] = NewElevator(fmt.Sprintf("E%d", i+1), capacity)
	}
	return &ElevatorController{
		Elevators: elevators,
	}
}

func (ec *ElevatorController) RequestElevator(req *Request) error {
	var optimalElevator *Elevator
	minDistance := math.MaxInt32

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
	return optimalElevator.AddRequest(req)
}

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
