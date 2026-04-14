package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("🚀 Initializing Concurrent Elevator System...")

	// 3 Elevators with capacity limit of 10
	controller := NewElevatorController(3, 10)

	var wg sync.WaitGroup

	requests := []*Request{
		{SourceFloor: 0, DestinationFloor: 5, Passengers: 2},
		{SourceFloor: 2, DestinationFloor: 8, Passengers: 4},
		{SourceFloor: 8, DestinationFloor: 1, Passengers: 3},
		{SourceFloor: 3, DestinationFloor: 6, Passengers: 5},
		{SourceFloor: 6, DestinationFloor: 0, Passengers: 2},
		{SourceFloor: 0, DestinationFloor: 10, Passengers: 9}, // Capacity test
	}

	// Random concurrent delays
	fmt.Println("Firing concurrent user requests across multiple floors...")

	for i, req := range requests {
		wg.Add(1)
		go func(r *Request, requestID int) {
			defer wg.Done()
			
			// Simulate users pressing the button simultaneously, but with tiny ms drift
			time.Sleep(time.Duration(requestID*15) * time.Millisecond)

			err := controller.RequestElevator(r)
			if err != nil {
				fmt.Printf("❌ Request %d Failed: %v\n", requestID, err)
			}
		}(req, i+1)
	}

	wg.Wait()

	fmt.Println("\nAll requests dispatched! Wait for Elevators to finish physical processing...")
	
	// Hold the main thread open to allow the internal Elevator goroutines to spin and physically hit the floors
	time.Sleep(3 * time.Second)
	
	fmt.Println("\n🏁 Elevator Simulation Completed.")
}
