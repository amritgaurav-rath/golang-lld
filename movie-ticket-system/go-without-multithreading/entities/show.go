package entities

import "time"

// Show intersects a Movie and Theater allocating a specific timeline grid.
type Show struct {
	ID        string
	Movie     *Movie
	Theater   *Theater
	StartTime time.Time
	EndTime   time.Time
	Seats     map[string]*Seat
}
