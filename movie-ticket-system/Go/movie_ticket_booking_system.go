package main

import (
	"fmt"
	"sync"
	"time"
)

// MovieTicketBookingSystem is the main class that manages the movie ticket booking system.
// It follows the Singleton pattern to ensure only one instance exists.
type MovieTicketBookingSystem struct {
	// Concurrent data structures (sync.Map in Go) to handle concurrent access
	movies   sync.Map // string -> *Movie
	theaters sync.Map // string -> *Theater
	shows    sync.Map // string -> *Show
	bookings sync.Map // string -> *Booking
}

var (
	systemInstance *MovieTicketBookingSystem
	systemOnce     sync.Once
)

func GetMovieTicketBookingSystem() *MovieTicketBookingSystem {
	systemOnce.Do(func() {
		systemInstance = &MovieTicketBookingSystem{}
	})
	return systemInstance
}

func (sys *MovieTicketBookingSystem) AddMovie(movie *Movie) {
	sys.movies.Store(movie.ID, movie)
}

func (sys *MovieTicketBookingSystem) GetMovie(id string) *Movie {
	if val, ok := sys.movies.Load(id); ok {
		return val.(*Movie)
	}
	return nil
}

func (sys *MovieTicketBookingSystem) AddTheater(theater *Theater) {
	sys.theaters.Store(theater.ID, theater)
}

func (sys *MovieTicketBookingSystem) GetTheater(id string) *Theater {
	if val, ok := sys.theaters.Load(id); ok {
		return val.(*Theater)
	}
	return nil
}

func (sys *MovieTicketBookingSystem) AddShow(show *Show) {
	sys.shows.Store(show.ID, show)
	// Mutate theater list
	if show.Theater != nil {
		show.Theater.Shows = append(show.Theater.Shows, show)
	}
}

func (sys *MovieTicketBookingSystem) GetShow(id string) *Show {
	if val, ok := sys.shows.Load(id); ok {
		return val.(*Show)
	}
	return nil
}

func (sys *MovieTicketBookingSystem) GetAvailableShows() []*Show {
	var available []*Show
	sys.shows.Range(func(key, value interface{}) bool {
		available = append(available, value.(*Show))
		return true
	})
	return available
}

func (sys *MovieTicketBookingSystem) BookTickets(user *User, show *Show, seatIds []string) (*Booking, error) {
	// Concurrent access to the shared resource (Show) requires locking
	show.Lock()
	defer show.Unlock()

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
		Status:     BookingStatusPending, // Not confirmed until payment
	}

	sys.bookings.Store(bookingId, booking)
	return booking, nil
}

func (sys *MovieTicketBookingSystem) ConfirmBooking(bookingId string) error {
	if val, ok := sys.bookings.Load(bookingId); ok {
		booking := val.(*Booking)
		// Simulating payment integration logic...
		if booking.Status == BookingStatusPending {
			booking.Status = BookingStatusConfirmed
			return nil
		}
		return fmt.Errorf("booking is not in pending state")
	}
	return fmt.Errorf("booking not found")
}

func (sys *MovieTicketBookingSystem) CancelBooking(bookingId string) error {
	if val, ok := sys.bookings.Load(bookingId); ok {
		booking := val.(*Booking)
		
		if booking.Status == BookingStatusCancelled {
			return fmt.Errorf("booking is already cancelled")
		}

		// Revert the seats to available state concurrently safely
		booking.Show.Lock()
		for _, seat := range booking.Seats {
			seat.Status = SeatStatusAvailable
		}
		booking.Show.Unlock()

		booking.Status = BookingStatusCancelled
		return nil
	}
	return fmt.Errorf("booking not found")
}

func (sys *MovieTicketBookingSystem) GetBooking(id string) *Booking {
	if val, ok := sys.bookings.Load(id); ok {
		return val.(*Booking)
	}
	return nil
}

// ResetMovieTicketBookingSystem is a test-helper purely because tests conflict overriding the singleton state
func ResetMovieTicketBookingSystem() {
	systemInstance = nil
	systemOnce = sync.Once{}
}
