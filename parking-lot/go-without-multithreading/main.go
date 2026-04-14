package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("🚀 Initializing Single-Threaded Parking Lot Simulation...")

	// Initialize 2 levels for the parking lot
	level1 := NewLevel(1, 5, 5, 2)
	level2 := NewLevel(2, 3, 4, 1)

	// Obtain the Singleton instance
	lot := GetParkingLotInstance([]*Level{level1, level2})

	lot.PrintAvailability()

	numVehicles := 30
	fmt.Printf("Simulating %d vehicles arriving sequentially...\n\n", numVehicles)

	// Sequential execution loop instead of goroutines
	for i := 1; i <= numVehicles; i++ {
		vTypes := []VehicleType{Motorcycle, Car, Truck}
		vType := vTypes[rand.Intn(len(vTypes))]
		plate := fmt.Sprintf("VEH-%03d", i)

		// Factory Pattern to create vehicles dynamically
		v, _ := VehicleFactory(vType, plate)

		spot, err := lot.ParkVehicle(v)
		if err != nil {
			fmt.Printf("❌ %s (%s) rejected: %s\n", v.GetLicensePlate(), v.GetType(), err.Error())
			continue
		}
		
		fmt.Printf("✅ %s (%s) PARKED in spot %s.\n", v.GetLicensePlate(), v.GetType(), spot.ID)

		// Synchronous unparking chance
		if rand.Intn(10) < 6 {
			_, unparkErr := lot.UnparkVehicle(spot.ID)
			if unparkErr != nil {
				fmt.Printf("⚠️  Error unparking %s: %s\n", v.GetLicensePlate(), unparkErr.Error())
			} else {
				fmt.Printf("👋 %s (%s) LEFT spot %s.\n", v.GetLicensePlate(), v.GetType(), spot.ID)
			}
		}
	}

	fmt.Println("\n🏁 Simulation completed.")
	lot.PrintAvailability()
}
