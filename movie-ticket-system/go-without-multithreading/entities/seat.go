package entities

import "app/movie-ticket-system/go-without-multithreading/enums"

// Seat captures explicit graph positions inside a specific show theater mapping.
type Seat struct {
	ID     string
	Row    int
	Column int
	Type   enums.SeatType
	Price  float64
	Status enums.SeatStatus
}
