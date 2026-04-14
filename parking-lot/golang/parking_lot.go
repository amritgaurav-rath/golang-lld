package main

import (
	"fmt"
	"sync"
)

// ParkingLot follows the Singleton pattern handling interactions globally
type ParkingLot struct {
	Levels []*Level
}

var (
	parkingLotInstance *ParkingLot
	parkingLotOnce     sync.Once
)

// GetParkingLotInstance strictly enforces the Singleton pattern natively using sync.Once
func GetParkingLotInstance(levels []*Level) *ParkingLot {
	parkingLotOnce.Do(func() {
		parkingLotInstance = &ParkingLot{
			Levels: levels,
		}
	})
	// If the user tries to overwrite after creation, it just returns the first one created
	return parkingLotInstance
}

// ResetParkingLotInstance is a test-helper purely because tests conflict overriding the singleton state
func ResetParkingLotInstance() {
	parkingLotInstance = nil
	parkingLotOnce = sync.Once{}
}

// ParkVehicle searches all levels for an available spot of the matching type and attempts to park.
func (p *ParkingLot) ParkVehicle(v Vehicle) (*ParkingSpot, error) {
	for _, level := range p.Levels {
		avail := level.GetAvailability()
		if avail[v.GetType()] > 0 {
			spot, err := level.ParkVehicle(v)
			if err == nil {
				return spot, nil
			}
		}
	}
	return nil, fmt.Errorf("parking lot is full for vehicle type: %s", v.GetType())
}

// UnparkVehicle frees up the parking spot across the lot by its ID.
func (p *ParkingLot) UnparkVehicle(spotID string) (Vehicle, error) {
	for _, level := range p.Levels {
		v, err := level.UnparkVehicle(spotID)
		if err == nil {
			return v, nil
		}
	}
	return nil, fmt.Errorf("spot %s not found in parking lot or is already empty", spotID)
}

// GetTotalAvailability aggregates the availability counts across all levels mapping the types.
func (p *ParkingLot) GetTotalAvailability() map[VehicleType]int {
	totalAvail := make(map[VehicleType]int)
	for _, level := range p.Levels {
		avail := level.GetAvailability()
		for k, v := range avail {
			totalAvail[k] += v
		}
	}
	return totalAvail
}

// PrintAvailability prints the current state of availability across the lot in a human-readable format.
func (p *ParkingLot) PrintAvailability() {
	fmt.Println("--- Current Parking Availability ---")
	avail := p.GetTotalAvailability()
	for k, v := range avail {
		fmt.Printf("%s Spots: %d\n", k, v)
	}
	fmt.Println("------------------------------------")
}
