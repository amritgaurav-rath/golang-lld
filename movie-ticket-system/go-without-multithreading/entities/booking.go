package entities

import "app/movie-ticket-system/go-without-multithreading/enums"

// Booking structurally packages transaction states directly tied to a User and precise Seats.
type Booking struct {
	ID         string
	User       *User
	Show       *Show
	Seats      []*Seat
	TotalPrice float64
	Status     enums.BookingStatus
}
