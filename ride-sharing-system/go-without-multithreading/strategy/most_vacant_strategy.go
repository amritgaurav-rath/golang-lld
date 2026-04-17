package strategy
import "app/ride-sharing-system/go-without-multithreading/entities"

// MostVacantStrategy calculates maximum remaining bounds natively ensuring systems minimize booking failures.
type MostVacantStrategy struct{}

func (s *MostVacantStrategy) SelectRide(availableRides []*entities.Ride) *entities.Ride {
	var mostVacant *entities.Ride
	maxSeats := -1

	// Explicitly map natively looping bounds securing max values mathematically.
	for _, ride := range availableRides {
		if ride.AvailableSeats > maxSeats {
			maxSeats = ride.AvailableSeats
			mostVacant = ride
		}
	}
	return mostVacant
}
