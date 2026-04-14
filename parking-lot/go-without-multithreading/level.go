package main

import (
	"fmt"
)

// Level represents a single floor in the parking lot
type Level struct {
	ID             int
	Spots          []*ParkingSpot
	AvailableSpots map[VehicleType]int // Tracks available spots
}

// NewLevel initializes a new level with a predefined number of spots for each type
func NewLevel(id int, numMotorcycleSpots, numCarSpots, numTruckSpots int) *Level {
	level := &Level{
		ID:             id,
		Spots:          make([]*ParkingSpot, 0, numMotorcycleSpots+numCarSpots+numTruckSpots),
		AvailableSpots: make(map[VehicleType]int),
	}

	level.AvailableSpots[Motorcycle] = numMotorcycleSpots
	level.AvailableSpots[Car] = numCarSpots
	level.AvailableSpots[Truck] = numTruckSpots

	spotNum := 1
	for i := 0; i < numMotorcycleSpots; i++ {
		level.Spots = append(level.Spots, &ParkingSpot{
			ID:      fmt.Sprintf("L%d-M%d", id, spotNum),
			Type:    Motorcycle,
			LevelID: id,
		})
		spotNum++
	}
	spotNum = 1
	for i := 0; i < numCarSpots; i++ {
		level.Spots = append(level.Spots, &ParkingSpot{
			ID:      fmt.Sprintf("L%d-C%d", id, spotNum),
			Type:    Car,
			LevelID: id,
		})
		spotNum++
	}
	spotNum = 1
	for i := 0; i < numTruckSpots; i++ {
		level.Spots = append(level.Spots, &ParkingSpot{
			ID:      fmt.Sprintf("L%d-T%d", id, spotNum),
			Type:    Truck,
			LevelID: id,
		})
		spotNum++
	}

	return level
}

// ParkVehicle attempts to park a vehicle on this level. Single-threaded execution.
func (l *Level) ParkVehicle(v Vehicle) (*ParkingSpot, error) {
	if l.AvailableSpots[v.GetType()] <= 0 {
		return nil, fmt.Errorf("no spots available for %s on level %d", v.GetType(), l.ID)
	}

	for _, spot := range l.Spots {
		if !spot.IsOccupied && spot.Type == v.GetType() {
			err := spot.Park(v)
			if err != nil {
				return nil, err
			}
			l.AvailableSpots[v.GetType()]--
			return spot, nil
		}
	}

	return nil, fmt.Errorf("unexpected error: spot count mismatch on level %d", l.ID)
}

// UnparkVehicle attempts to free a spot given the spotID. Single-threaded execution.
func (l *Level) UnparkVehicle(spotID string) (Vehicle, error) {
	for _, spot := range l.Spots {
		if spot.ID == spotID {
			if !spot.IsOccupied {
				return nil, fmt.Errorf("spot %s is already empty", spotID)
			}
			v := spot.Vehicle
			spot.Unpark()
			l.AvailableSpots[v.GetType()]++
			return v, nil
		}
	}
	return nil, fmt.Errorf("spot %s not found on level %d", spotID, l.ID)
}

// GetAvailability returns a copy of the available spots on this level.
func (l *Level) GetAvailability() map[VehicleType]int {
	avail := make(map[VehicleType]int)
	for k, v := range l.AvailableSpots {
		avail[k] = v
	}
	return avail
}
