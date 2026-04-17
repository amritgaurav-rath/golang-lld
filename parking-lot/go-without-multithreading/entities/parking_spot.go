package entities
import (
    "app/parking-lot/go-without-multithreading/enums"
    "fmt"
)

type ParkingSpot struct {
	ID         string
	Type       enums.VehicleType
	IsOccupied bool
	Vehicle    Vehicle
	LevelID    int
}

func (s *ParkingSpot) Park(v Vehicle) error {
	if s.IsOccupied {
		return fmt.Errorf("spot mapping internally compromised (Already Occupied)")
	}
	s.Vehicle = v
	s.IsOccupied = true
	return nil
}

func (s *ParkingSpot) Unpark() {
	s.Vehicle = nil
	s.IsOccupied = false
}
