package entities

// User natively maps statistical tracking points (Offered/Taken) decoupled strictly from execution logics.
type User struct {
	ID           string
	Name         string
	TakenRides   int // Increments securely upon successful passenger matching execution
	OfferedRides int // Tracks exact mathematically completed origin-to-destination events
}
