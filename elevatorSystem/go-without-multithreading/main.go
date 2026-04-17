package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🚀 Initializing Synchronous (Single-Threaded) Elevator System...")

	// Initialize the system with 3 elevators and a max capacity of 10 passengers each
	controller := NewElevatorController(3, 10)

	// Simulate a batch of passenger requests
	requests := []*Request{
		{SourceFloor: 0, DestinationFloor: 5, Passengers: 2},
		{SourceFloor: 2, DestinationFloor: 8, Passengers: 4},
		{SourceFloor: 8, DestinationFloor: 1, Passengers: 3},
		{SourceFloor: 3, DestinationFloor: 6, Passengers: 5},
		{SourceFloor: 6, DestinationFloor: 0, Passengers: 2},
		{SourceFloor: 0, DestinationFloor: 10, Passengers: 9},
	}

	fmt.Println("Sequentially routing all user requests immediately...")

	// Dispatch each request sequentially to the most optimal elevator natively (no WaitGroups)
	for i, req := range requests {
		err := controller.RequestElevator(req)
		if err != nil {
			fmt.Printf("❌ Request %d Failed: %v\n", i+1, err)
		}
	}

	fmt.Println("\nExecuting deterministic Tick loop to spin Elevator motors physically...")
	
	// Single threaded "game loop" that advances the system state iteratively until all elevators are IDLE
	step := 1
	for {
		active := controller.TickAll() // Advance one time step for all elevators
		if !active {
			break
		}
		
		fmt.Printf("--- Time Step %d ---\n", step)
		step++
		time.Sleep(100 * time.Millisecond) // Simulated physics pace
	}
	
	fmt.Println("\n🏁 Synchronous Elevator Simulation Completed.")
}
