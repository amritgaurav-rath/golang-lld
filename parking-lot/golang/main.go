package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("🚀 Initializing Parking Lot Simulation...")

	// Initialize 2 levels for the parking lot
	// Level 1: 5 Motorcycle, 5 Car, 2 Truck spots
	level1 := NewLevel(1, 5, 5, 2)
	// Level 2: 3 Motorcycle, 4 Car, 1 Truck spots
	level2 := NewLevel(2, 3, 4, 1)

	lot := NewParkingLot([]*Level{level1, level2})

	lot.PrintAvailability()

	var wg sync.WaitGroup

	// Simulate 30 concurrent vehicles arriving and trying to park
	numVehicles := 30
	fmt.Printf("Simulating %d vehicles arriving concurrently...\n\n", numVehicles)

	for i := 1; i <= numVehicles; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Randomly pick a vehicle type
			vTypes := []VehicleType{Motorcycle, Car, Truck}
			vType := vTypes[rand.Intn(len(vTypes))]
			v := &Vehicle{
				LicensePlate: fmt.Sprintf("VEH-%03d", id),
				Type:         vType,
			}

			// Add a slight jitter to simulate staggered parallel arrivals
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

			spot, err := lot.ParkVehicle(v)
			if err != nil {
				fmt.Printf("❌ %s (%s) rejected: %s\n", v.LicensePlate, v.Type, err.Error())
				return
			}
			
			fmt.Printf("✅ %s (%s) PARKED in spot %s.\n", v.LicensePlate, v.Type, spot.ID)

			// Simulate time spent parked (10 to 50 milliseconds)
			time.Sleep(time.Duration(rand.Intn(40)+10) * time.Millisecond)

			// Some vehicles decide to unpark
			if rand.Intn(10) < 6 { // 60% chance to leave during simulation
				_, unparkErr := lot.UnparkVehicle(spot.ID)
				if unparkErr != nil {
					fmt.Printf("⚠️  Error unparking %s: %s\n", v.LicensePlate, unparkErr.Error())
				} else {
					fmt.Printf("👋 %s (%s) LEFT spot %s.\n", v.LicensePlate, v.Type, spot.ID)
				}
			}
		}(i)
	}

	// Wait for all vehicles to finish their simulation actions
	wg.Wait()
	
	fmt.Println("\n🏁 Simulation completed.")
	lot.PrintAvailability()
}
