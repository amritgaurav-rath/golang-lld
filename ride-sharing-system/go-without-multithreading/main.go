package main

import (
	"app/ride-sharing-system/go-without-multithreading/enums"
	"app/ride-sharing-system/go-without-multithreading/services"
	"app/ride-sharing-system/go-without-multithreading/strategy"
	"fmt"
)

func main() {
	fmt.Println("🚀 Initializing Golang S.O.L.I.D. FAANG SDE II Ride-Sharing Platform...")

	sys := services.GetInstance()

	// 1. Onboard Nodes natively
	sys.AddUser("U1", "Rohan")
	sys.AddUser("U2", "Shashank")
	sys.AddUser("U3", "Nandini")

	// 2. Add vehicles safely mapping explicitly typed values
	sys.AddVehicle("U1", "KA-01-1234", enums.Polo)
	sys.AddVehicle("U2", "TS-05-6239", enums.Activa)
	sys.AddVehicle("U3", "MH-12-876D", enums.XUV)

	fmt.Println("✅ Users and their Vehicles safely mapped.\n")

	// 3. Offer Rides structurally bypassing generic matrices
	sys.OfferRide("R1", "U1", "KA-01-1234", "Hyderabad", "Bangalore", 1)
	sys.OfferRide("R2", "U2", "TS-05-6239", "Bangalore", "Mysore", 1)
	sys.OfferRide("R3", "U3", "MH-12-876D", "Hyderabad", "Bangalore", 4)

	fmt.Println("✅ Rides flawlessly created globally.\n")

	// Rule check: Try offering duplicated vehicle actively natively
	err := sys.OfferRide("R4", "U1", "KA-01-1234", "Bangalore", "Pune", 2)
	if err != nil {
		fmt.Println("🔒 Expected Validation mathematical drop:", err.Error())
	}

	fmt.Println("\n--- Invoking Strategy Routing Algorithms ---")

	// Nandini selects ride searching for an ACTIVA specifically (PreferredVehicleStrategy)
	rideA, err := sys.SelectRide("U3", "Bangalore", "Mysore", 1, &strategy.PreferredVehicleStrategy{PreferredType: enums.Activa})
	if err != nil {
		fmt.Println("❌ Strategy Fallback:", err.Error())
	} else {
		fmt.Printf("✅ Nandini securely matched ACTIVA Ride '%s' via mathematical Strategy parameters!\n", rideA.RideID)
		sys.EndRide(rideA.RideID)
	}

	// Rohan dynamically loops for MostVacantStrategy natively!
	rideB, err := sys.SelectRide("U1", "Hyderabad", "Bangalore", 2, &strategy.MostVacantStrategy{})
	if err != nil {
		fmt.Println("❌ Match fail cleanly:", err.Error())
	} else {
		fmt.Printf("✅ Rohan mathematically locked largest available structural mapping: Ride '%s'! Capacity remaining natively: %d!\n",
			rideB.RideID, rideB.AvailableSeats)
		sys.EndRide(rideB.RideID)
	}

	// 5. Native Statistical runtime dumps
	sys.PrintRideStats()
}
