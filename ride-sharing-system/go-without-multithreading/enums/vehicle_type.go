package enums

// VehicleType formally tracks SDE II specific platform transport limits exactly.
// It completely decouples String matching requirements for absolute memory stability.
type VehicleType int

const (
	Activa VehicleType = iota
	Polo
	XUV
)
