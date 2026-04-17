package entities
import "app/ride-sharing-system/go-without-multithreading/enums"

// Ride explicitly represents active physical operations in memory entirely skipping NoSQL.
// AvailableSeats dynamically updates natively ensuring physical locks prevent overbooking calculations.
type Ride struct {
	RideID         string
	DriverID       string
	Vehicle        *Vehicle
	Origin         string
	Destination    string
	AvailableSeats int
	Status         enums.RideStatus
}
