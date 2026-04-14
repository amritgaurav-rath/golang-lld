package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("🚀 Initializing Parking Lot Simulation (Singleton & Interface Model)...")

	// Initialize 2 levels for the parking lot
	level1 := NewLevel(1, 5, 5, 2)
	level2 := NewLevel(2, 3, 4, 1)

	// Obtain the Singleton instance
	lot := GetParkingLotInstance([]*Level{level1, level2})

	lot.PrintAvailability()

	var wg sync.WaitGroup

	numVehicles := 30
	fmt.Printf("Simulating %d vehicles arriving concurrently...\n\n", numVehicles)

	for i := 1; i <= numVehicles; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			vTypes := []VehicleType{Motorcycle, Car, Truck}
			vType := vTypes[rand.Intn(len(vTypes))]
			plate := fmt.Sprintf("VEH-%03d", id)

			// Factory Pattern to create vehicles dynamically
			v, _ := VehicleFactory(vType, plate)

			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

			spot, err := lot.ParkVehicle(v)
			if err != nil {
				fmt.Printf("❌ %s (%s) rejected: %s\n", v.GetLicensePlate(), v.GetType(), err.Error())
				return
			}
			
			fmt.Printf("✅ %s (%s) PARKED in spot %s.\n", v.GetLicensePlate(), v.GetType(), spot.ID)

			time.Sleep(time.Duration(rand.Intn(40)+10) * time.Millisecond)

			if rand.Intn(10) < 6 {
				_, unparkErr := lot.UnparkVehicle(spot.ID)
				if unparkErr != nil {
					fmt.Printf("⚠️  Error unparking %s: %s\n", v.GetLicensePlate(), unparkErr.Error())
				} else {
					fmt.Printf("👋 %s (%s) LEFT spot %s.\n", v.GetLicensePlate(), v.GetType(), spot.ID)
				}
			}
		}(i)
	}

	wg.Wait()
	
	fmt.Println("\n🏁 Simulation completed.")
	lot.PrintAvailability()
}
