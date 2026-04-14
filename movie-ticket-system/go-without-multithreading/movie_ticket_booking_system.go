package main

import (
	"fmt"
	"time"
)

// MovieTicketBookingSystem (Single-threaded version)
type MovieTicketBookingSystem struct {
	movies   map[string]*Movie
	theaters map[string]*Theater
	shows    map[string]*Show
	bookings map[string]*Booking
}

var systemInstance *MovieTicketBookingSystem

func GetMovieTicketBookingSystem() *MovieTicketBookingSystem {
	if systemInstance == nil {
		systemInstance = &MovieTicketBookingSystem{
			movies:   make(map[string]*Movie),
			theaters: make(map[string]*Theater),
			shows:    make(map[string]*Show),
			bookings: make(map[string]*Booking),
		}
	}
	return systemInstance
}

func (sys *MovieTicketBookingSystem) AddMovie(movie *Movie) {
	sys.movies[movie.ID] = movie
}

func (sys *MovieTicketBookingSystem) GetMovie(id string) *Movie {
	return sys.movies[id]
}

func (sys *MovieTicketBookingSystem) AddTheater(theater *Theater) {
	sys.theaters[theater.ID] = theater
}

func (sys *MovieTicketBookingSystem) GetTheater(id string) *Theater {
	return sys.theaters[id]
}

func (sys *MovieTicketBookingSystem) AddShow(show *Show) {
	sys.shows[show.ID] = show
	if show.Theater != nil {
		show.Theater.Shows = append(show.Theater.Shows, show)
	}
}

func (sys *MovieTicketBookingSystem) GetShow(id string) *Show {
	return sys.shows[id]
}

func (sys *MovieTicketBookingSystem) GetAvailableShows() []*Show {
	var available []*Show
	for _, show := range sys.shows {
		available = append(available, show)
	}
	return available
}

func (sys *MovieTicketBookingSystem) BookTickets(user *User, show *Show, seatIds []string) (*Booking, error) {
	var selectedSeats []*Seat
	totalPrice := 0.0

	// 1. Verify availability
	for _, seatId := range seatIds {
		seat, exists := show.Seats[seatId]
		if !exists {
			return nil, fmt.Errorf("seat %s does not exist", seatId)
		}
		if seat.Status != SeatStatusAvailable {
			return nil, fmt.Errorf("seat %s is not available", seatId)
		}
		selectedSeats = append(selectedSeats, seat)
		totalPrice += seat.Price
	}

	// 2. Mark seats as Booked
	for _, seat := range selectedSeats {
		seat.Status = SeatStatusBooked
	}

	// 3. Create Pending Booking
	bookingId := fmt.Sprintf("BKG-%d", time.Now().UnixNano())
	booking := &Booking{
		ID:         bookingId,
		User:       user,
		Show:       show,
		Seats:      selectedSeats,
		TotalPrice: totalPrice,
		Status:     BookingStatusPending,
	}

	sys.bookings[bookingId] = booking
	return booking, nil
}

func (sys *MovieTicketBookingSystem) ConfirmBooking(bookingId string) error {
	booking, exists := sys.bookings[bookingId]
	if !exists {
		return fmt.Errorf("booking not found")
	}
	if booking.Status == BookingStatusPending {
		booking.Status = BookingStatusConfirmed
		return nil
	}
	return fmt.Errorf("booking is not in pending state")
}

func (sys *MovieTicketBookingSystem) CancelBooking(bookingId string) error {
	booking, exists := sys.bookings[bookingId]
	if !exists {
		return fmt.Errorf("booking not found")
	}
	if booking.Status == BookingStatusCancelled {
		return fmt.Errorf("booking is already cancelled")
	}

	for _, seat := range booking.Seats {
		seat.Status = SeatStatusAvailable
	}
	booking.Status = BookingStatusCancelled
	return nil
}

func (sys *MovieTicketBookingSystem) GetBooking(id string) *Booking {
	return sys.bookings[id]
}
