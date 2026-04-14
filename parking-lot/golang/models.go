package main

import "fmt"

// VehicleType defines the types of vehicles supported
type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Truck
)

func (v VehicleType) String() string {
	switch v {
	case Motorcycle:
		return "Motorcycle"
	case Car:
		return "Car"
	case Truck:
		return "Truck"
	default:
		return "Unknown"
	}
}

// Vehicle represents a vehicle parking in the lot
type Vehicle struct {
	LicensePlate string
	Type         VehicleType
}

// ParkingSpot represents a single parking space
type ParkingSpot struct {
	ID         string
	Type       VehicleType // using VehicleType as SpotType for simplicity mapping 1:1
	IsOccupied bool
	Vehicle    *Vehicle
	LevelID    int // Using int ID for easier identification
}

// Park assigns a vehicle to this spot
func (s *ParkingSpot) Park(v *Vehicle) error {
	if s.IsOccupied {
		return fmt.Errorf("spot is already occupied")
	}
	// Note: Validation against spot type is typically done at the level before parking.
	s.Vehicle = v
	s.IsOccupied = true
	return nil
}

// Unpark removes the vehicle from the spot
func (s *ParkingSpot) Unpark() {
	s.Vehicle = nil
	s.IsOccupied = false
}
