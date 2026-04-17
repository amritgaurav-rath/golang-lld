package services

import (
	"app/movie-ticket-system/go-without-multithreading/entities"
	"app/movie-ticket-system/go-without-multithreading/enums"
    "app/movie-ticket-system/go-without-multithreading/observer"
    "app/movie-ticket-system/go-without-multithreading/strategy"
	"fmt"
	"time"
)

// MovieTicketBookingSystem (Single-threaded Façade)
type MovieTicketBookingSystem struct {
	movies    map[string]*entities.Movie
	theaters  map[string]*entities.Theater
	shows     map[string]*entities.Show
	bookings  map[string]*entities.Booking
    observers map[string][]observer.MovieObserver // Subscription Matrix
}

var systemInstance *MovieTicketBookingSystem

func GetMovieTicketBookingSystem() *MovieTicketBookingSystem {
	if systemInstance == nil {
		systemInstance = &MovieTicketBookingSystem{
			movies:    make(map[string]*entities.Movie),
			theaters:  make(map[string]*entities.Theater),
			shows:     make(map[string]*entities.Show),
			bookings:  make(map[string]*entities.Booking),
            observers: make(map[string][]observer.MovieObserver),
		}
	}
	return systemInstance
}

func (sys *MovieTicketBookingSystem) Subscribe(movieID string, obs observer.MovieObserver) {
    sys.observers[movieID] = append(sys.observers[movieID], obs)
}

func (sys *MovieTicketBookingSystem) AddMovie(movie *entities.Movie) {
	sys.movies[movie.ID] = movie
    
    // Explicit Observer Notification Dispatch Trigger
    if len(sys.observers[movie.ID]) > 0 {
		for _, obs := range sys.observers[movie.ID] {
			obs.Update(movie.Title)
		}
	}
}

func (sys *MovieTicketBookingSystem) GetMovie(id string) *entities.Movie {
	return sys.movies[id]
}

func (sys *MovieTicketBookingSystem) AddTheater(theater *entities.Theater) {
	sys.theaters[theater.ID] = theater
}

func (sys *MovieTicketBookingSystem) GetTheater(id string) *entities.Theater {
	return sys.theaters[id]
}

func (sys *MovieTicketBookingSystem) AddShow(show *entities.Show) {
	sys.shows[show.ID] = show
	if show.Theater != nil {
		show.Theater.Shows = append(show.Theater.Shows, show)
	}
}

func (sys *MovieTicketBookingSystem) GetShow(id string) *entities.Show {
	return sys.shows[id]
}

func (sys *MovieTicketBookingSystem) GetAvailableShows() []*entities.Show {
	var available []*entities.Show
	for _, show := range sys.shows {
		available = append(available, show)
	}
	return available
}

func (sys *MovieTicketBookingSystem) BookTickets(
    user *entities.User, show *entities.Show, seatIds []string, pricingModel strategy.PricingStrategy) (*entities.Booking, error) {
	
    var selectedSeats []*entities.Seat
	totalPrice := 0.0

	for _, seatId := range seatIds {
		seat, exists := show.Seats[seatId]
		if !exists {
			return nil, fmt.Errorf("seat %s does not exist", seatId)
		}
		if seat.Status != enums.SeatStatusAvailable {
			return nil, fmt.Errorf("seat %s is not available", seatId)
		}
		selectedSeats = append(selectedSeats, seat)
        
        // Push Strategy Calculations cleanly in loop decoupling math from struct fields
		totalPrice += pricingModel.CalculatePrice(seat.Price)
	}

	for _, seat := range selectedSeats {
		seat.Status = enums.SeatStatusBooked
	}

	bookingId := fmt.Sprintf("BKG-%d", time.Now().UnixNano())
	booking := &entities.Booking{
		ID:         bookingId,
		User:       user,
		Show:       show,
		Seats:      selectedSeats,
		TotalPrice: totalPrice,
		Status:     enums.BookingStatusPending,
	}

	sys.bookings[bookingId] = booking
	return booking, nil
}

func (sys *MovieTicketBookingSystem) ConfirmBooking(bookingId string) error {
	booking, exists := sys.bookings[bookingId]
	if !exists {
		return fmt.Errorf("booking not found")
	}
	if booking.Status == enums.BookingStatusPending {
		booking.Status = enums.BookingStatusConfirmed
		return nil
	}
	return fmt.Errorf("booking is not in pending state")
}

func (sys *MovieTicketBookingSystem) CancelBooking(bookingId string) error {
	booking, exists := sys.bookings[bookingId]
	if !exists {
		return fmt.Errorf("booking not found")
	}
	if booking.Status == enums.BookingStatusCancelled {
		return fmt.Errorf("booking is already cancelled")
	}

	for _, seat := range booking.Seats {
		seat.Status = enums.SeatStatusAvailable
	}
	booking.Status = enums.BookingStatusCancelled
	return nil
}

func (sys *MovieTicketBookingSystem) GetBooking(id string) *entities.Booking {
	return sys.bookings[id]
}
