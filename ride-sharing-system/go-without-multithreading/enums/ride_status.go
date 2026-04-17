package enums

// RideStatus cleanly binds explicit active states avoiding physical dictionary deletion overhead natively.
// Enables maintaining mathematical logs mathematically identical to DB queries.
type RideStatus int

const (
	Offered RideStatus = iota
	Completed
)
