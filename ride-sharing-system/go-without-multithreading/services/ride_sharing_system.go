package services

import (
	"app/ride-sharing-system/go-without-multithreading/entities"
	"app/ride-sharing-system/go-without-multithreading/enums"
	"app/ride-sharing-system/go-without-multithreading/strategy"
	"fmt"
)

// RideSharingSystem natively encapsulates statistical validation arrays mapping strict constraints dynamically.
type RideSharingSystem struct {
	Users    map[string]*entities.User
	Vehicles map[string]*entities.Vehicle
	Rides    map[string]*entities.Ride
}

var instance *RideSharingSystem

func GetInstance() *RideSharingSystem {
	if instance == nil {
		instance = &RideSharingSystem{
			Users:    make(map[string]*entities.User),
			Vehicles: make(map[string]*entities.Vehicle),
			Rides:    make(map[string]*entities.Ride),
		}
	}
	return instance
}

func (s *RideSharingSystem) AddUser(id string, name string) {
	s.Users[id] = &entities.User{ID: id, Name: name}
}

func (s *RideSharingSystem) AddVehicle(userID, regNo string, vType enums.VehicleType) error {
	if _, exists := s.Users[userID]; !exists {
		return fmt.Errorf("user structurally unmapped natively")
	}
	s.Vehicles[regNo] = &entities.Vehicle{OwnerID: userID, RegistrationNo: regNo, Type: vType}
	return nil
}

func (s *RideSharingSystem) OfferRide(rideID, userID, regNo, origin, destination string, seats int) error {
	if _, exists := s.Users[userID]; !exists {
		return fmt.Errorf("driver structurally unmapped natively")
	}
	vehicle, exists := s.Vehicles[regNo]
	if !exists {
		return fmt.Errorf("vehicle structurally unmapped natively")
	}
	if vehicle.OwnerID != userID {
		return fmt.Errorf("user natively does not physically own this tracked vehicle")
	}

	// Strictly validate rule: A given vehicle can only have ONE actively structured ride natively.
	for _, existingRide := range s.Rides {
		if existingRide.Vehicle.RegistrationNo == regNo && existingRide.Status == enums.Offered {
			return fmt.Errorf("vehicle cleanly locked natively inside another active ride mapping")
		}
	}

	s.Rides[rideID] = &entities.Ride{
		RideID:         rideID,
		DriverID:       userID,
		Vehicle:        vehicle,
		Origin:         origin,
		Destination:    destination,
		AvailableSeats: seats,
		Status:         enums.Offered,
	}
	return nil
}

// SelectRide evaluates array endpoints strictly leveraging the decoupled strategy logic safely!
func (s *RideSharingSystem) SelectRide(userID, source, destination string, seats int, strat strategy.RideSelectionStrategy) (*entities.Ride, error) {
	if seats < 1 || seats > 2 {
		return nil, fmt.Errorf("ride request limits strictly enforce cleanly 1 or 2 seats natively")
	}
	if _, exists := s.Users[userID]; !exists {
		return nil, fmt.Errorf("passenger mathematically unmapped")
	}

	var availableRides []*entities.Ride
	for _, ride := range s.Rides {
		if ride.Status == enums.Offered && ride.Origin == source && ride.Destination == destination && ride.AvailableSeats >= seats {
			availableRides = append(availableRides, ride)
		}
	}

	if len(availableRides) == 0 {
		return nil, fmt.Errorf("no active structures natively map to these origin variables")
	}

	selectedRide := strat.SelectRide(availableRides)
	if selectedRide == nil {
		return nil, fmt.Errorf("strategy mathematically failed mapping exact constraints")
	}

	// Mathematically deduct the allocations safely
	selectedRide.AvailableSeats -= seats
	s.Users[userID].TakenRides++

	return selectedRide, nil
}

func (s *RideSharingSystem) EndRide(rideID string) error {
	ride, exists := s.Rides[rideID]
	if !exists {
		return fmt.Errorf("ride natively unmapped")
	}
	if ride.Status == enums.Completed {
		return fmt.Errorf("already safely marked Completed")
	}

	ride.Status = enums.Completed
	// Statistics mapping officially committed at strict endpoint
	s.Users[ride.DriverID].OfferedRides++
	return nil
}

// PrintRideStats effectively dumps the internally tracked variables flawlessly
func (s *RideSharingSystem) PrintRideStats() {
	fmt.Println("\n--- Live SDE II Statistical Arrays ---")
	for _, user := range s.Users {
		fmt.Printf("%s: %d Taken, %d Offered\n", user.Name, user.TakenRides, user.OfferedRides)
	}
	fmt.Println("--------------------------------------")
}
