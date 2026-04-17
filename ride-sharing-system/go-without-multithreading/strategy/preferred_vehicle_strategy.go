package strategy
import (
	"app/ride-sharing-system/go-without-multithreading/entities"
	"app/ride-sharing-system/go-without-multithreading/enums"
)

// PreferredVehicleStrategy routes mathematical iterations prioritizing absolute parameter equivalence natively.
type PreferredVehicleStrategy struct {
	PreferredType enums.VehicleType
}

func (s *PreferredVehicleStrategy) SelectRide(availableRides []*entities.Ride) *entities.Ride {
	for _, ride := range availableRides {
		if ride.Vehicle.Type == s.PreferredType {
			return ride // Physically returns the first matched execution node!
		}
	}
	return nil
}
