package main

import (
	"app/parking-lot/go-without-multithreading/entities"
	"app/parking-lot/go-without-multithreading/services"
	"app/parking-lot/go-without-multithreading/strategy/fee"
	"app/parking-lot/go-without-multithreading/strategy/parking"
	"fmt"
)

func main() {
	fmt.Println("🚀 Initializing FAANG Tier S.O.L.I.D. Parking Lot (Strategies Active)...")

	level1 := entities.NewLevel(1, 5, 5, 2)
	level2 := entities.NewLevel(2, 3, 4, 1)

	sys := services.GetInstance()
	sys.SetLevels([]*entities.Level{level1, level2})

	// Inject Dynamic Pattern Strategies!
	sys.SetParkingStrategy(&parking.FarthestFirstStrategy{}) // Inverts array to fill from backwards
	sys.SetFeeStrategy(&fee.VehicleBasedFeeStrategy{})       // Generates dynamic billing cleanly

	sys.PrintAvailability()

	fmt.Println("Alice sequence: Parks Car, Parks Motorcycle via FarthestFirstStrategy.")

	aliceCar := &entities.CarVehicle{BaseVehicle: entities.BaseVehicle{LicensePlate: "ALICE-01"}}
	aliceMoto := &entities.MotorcycleVehicle{BaseVehicle: entities.BaseVehicle{LicensePlate: "ALICE-02"}}

	spot1, err := sys.ParkVehicle(aliceCar)
	if err != nil {
		fmt.Println("❌ ERROR: ", err.Error())
	} else {
		fmt.Println("✅ Alice successfully PARKED Car at:", spot1.ID, "(Expect Level 2!)")
	}

	spot2, err := sys.ParkVehicle(aliceMoto)
	if err != nil {
		fmt.Println("❌ ERROR: ", err.Error())
	} else {
		fmt.Println("✅ Alice successfully PARKED Motorcycle at:", spot2.ID, "(Expect Level 2!)")
	}

	sys.PrintAvailability()

	fmt.Println("Alice unparks Car...")
	unparked, err := sys.UnparkVehicle(spot1.ID)
	if err != nil {
		fmt.Println("❌ ERROR:", err.Error())
	} else {
		fmt.Println("👋 Unparked safely:", unparked.GetLicensePlate())
	}

	sys.PrintAvailability()
}
