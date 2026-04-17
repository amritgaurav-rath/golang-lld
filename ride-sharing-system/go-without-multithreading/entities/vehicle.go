package entities
import "app/ride-sharing-system/go-without-multithreading/enums"

// Vehicle securely links an owner and physical vehicle mapping structures mathematically.
// Binds strict typing avoiding string manipulation errors natively!
type Vehicle struct {
	OwnerID        string
	RegistrationNo string
	Type           enums.VehicleType
}
