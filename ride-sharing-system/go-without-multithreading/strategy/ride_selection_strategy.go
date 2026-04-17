package strategy
import "app/ride-sharing-system/go-without-multithreading/entities"

// RideSelectionStrategy strongly abstractions dynamic filtering metrics without hardcoding algorithm nodes.
// Explicitly maps out mathematical search implementations decoupled natively from Service interfaces.
type RideSelectionStrategy interface {
	SelectRide(availableRides []*entities.Ride) *entities.Ride
}
