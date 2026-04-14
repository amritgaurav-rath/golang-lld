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

// Vehicle defines the abstract base interface
type Vehicle interface {
	GetLicensePlate() string
	GetType() VehicleType
}

// BaseVehicle provides the root attributes shared across all vehicles
type BaseVehicle struct {
	LicensePlate string
}

func (b *BaseVehicle) GetLicensePlate() string {
	return b.LicensePlate
}

// CarVehicle represents a car, extending BaseVehicle
type CarVehicle struct {
	BaseVehicle
}

func (c *CarVehicle) GetType() VehicleType {
	return Car
}

// MotorcycleVehicle represents a motorcycle, extending BaseVehicle
type MotorcycleVehicle struct {
	BaseVehicle
}

func (m *MotorcycleVehicle) GetType() VehicleType {
	return Motorcycle
}

// TruckVehicle represents a truck, extending BaseVehicle
type TruckVehicle struct {
	BaseVehicle
}

func (t *TruckVehicle) GetType() VehicleType {
	return Truck
}

// VehicleFactory implements the Factory pattern to create distinct vehicle types easily.
func VehicleFactory(vType VehicleType, licensePlate string) (Vehicle, error) {
	base := BaseVehicle{LicensePlate: licensePlate}
	switch vType {
	case Car:
		return &CarVehicle{BaseVehicle: base}, nil
	case Motorcycle:
		return &MotorcycleVehicle{BaseVehicle: base}, nil
	case Truck:
		return &TruckVehicle{BaseVehicle: base}, nil
	default:
		return nil, fmt.Errorf("unknown vehicle type")
	}
}

// ParkingSpot represents a single parking space
type ParkingSpot struct {
	ID         string
	Type       VehicleType // Mapping 1:1 spot type to vehicle type
	IsOccupied bool
	Vehicle    Vehicle
	LevelID    int
}

// Park assigns a vehicle to this spot
func (s *ParkingSpot) Park(v Vehicle) error {
	if s.IsOccupied {
		return fmt.Errorf("spot is already occupied")
	}
	s.Vehicle = v
	s.IsOccupied = true
	return nil
}

// Unpark removes the vehicle from the spot
func (s *ParkingSpot) Unpark() {
	s.Vehicle = nil
	s.IsOccupied = false
}
