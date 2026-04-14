package main

import (
	"fmt"
)

// ParkingLot represents the overarching parking system managing multiple levels
type ParkingLot struct {
	Levels []*Level
}

// NewParkingLot creates a Parking Lot with the given initialized levels.
func NewParkingLot(levels []*Level) *ParkingLot {
	return &ParkingLot{
		Levels: levels,
	}
}

// ParkVehicle searches all levels for an available spot of the matching type and attempts to park.
func (p *ParkingLot) ParkVehicle(v *Vehicle) (*ParkingSpot, error) {
	for _, level := range p.Levels {
		// Quick check before attempting to acquire the lock to park
		avail := level.GetAvailability()
		if avail[v.Type] > 0 {
			spot, err := level.ParkVehicle(v)
			if err == nil {
				// Successfully parked
				return spot, nil
			}
			// If err != nil here, it means due to a race condition, another concurrent thread
			// might have taken the last spot between our GetAvailability() check and ParkVehicle() call.
			// In that case, we gracefully continue to check the next levels.
		}
	}
	return nil, fmt.Errorf("parking lot is full for vehicle type: %s", v.Type)
}

// UnparkVehicle frees up the parking spot across the lot by its ID.
func (p *ParkingLot) UnparkVehicle(spotID string) (*Vehicle, error) {
	for _, level := range p.Levels {
		v, err := level.UnparkVehicle(spotID)
		if err == nil {
			return v, nil
		}
		// If err != nil, the spot might be empty or belong to a different level.
		// We continue searching other levels, though a better optimization could
		// decode the level ID directly from the spotID (e.g., L1-M1) and jump to it.
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
