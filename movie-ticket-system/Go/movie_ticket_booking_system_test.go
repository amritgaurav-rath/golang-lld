package main

import (
	"sync"
	"testing"
	"time"
)

func TestMovieTicketBookingSystem_ConcurrentBooking(t *testing.T) {
	ResetMovieTicketBookingSystem()
	sys := GetMovieTicketBookingSystem()

	movie1 := &Movie{ID: "M1", Title: "Test Movie"}
	sys.AddMovie(movie1)

	theater1 := &Theater{ID: "T1", Name: "Test Theater"}
	sys.AddTheater(theater1)

	show1 := &Show{
		ID:        "SH1",
		Movie:     movie1,
		Theater:   theater1,
		StartTime: time.Now().Add(1 * time.Hour),
		EndTime:   time.Now().Add(3 * time.Hour),
		Seats:     make(map[string]*Seat),
	}
	sys.AddShow(show1)

	show1.Seats["A1"] = &Seat{ID: "A1", Type: SeatTypeNormal, Price: 10.0, Status: SeatStatusAvailable}
	show1.Seats["A2"] = &Seat{ID: "A2", Type: SeatTypeNormal, Price: 10.0, Status: SeatStatusAvailable}

	user1 := &User{ID: "U1", Name: "Alice"}
	user2 := &User{ID: "U2", Name: "Bob"}

	var wg sync.WaitGroup
	var successCount int
	var failCount int
	var mu sync.Mutex

	seatRequest := []string{"A1", "A2"}

	// Launch 10 concurrent requests for the exact same seats
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			user := user1
			if id%2 == 0 {
				user = user2
			}
			_, err := sys.BookTickets(user, show1, seatRequest)
			
			mu.Lock()
			if err == nil {
				successCount++
			} else {
				failCount++
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	if successCount != 1 {
		t.Errorf("Expected exactly 1 successful booking for the same seats, got %d", successCount)
	}
	if failCount != 9 {
		t.Errorf("Expected exactly 9 failed bookings, got %d", failCount)
	}

	if show1.Seats["A1"].Status != SeatStatusBooked {
		t.Errorf("Expected A1 to be booked")
	}
}

func TestMovieTicketBookingSystem_ConfirmAndCancel(t *testing.T) {
	ResetMovieTicketBookingSystem()
	sys := GetMovieTicketBookingSystem()

	show := &Show{ID: "SH2", Seats: make(map[string]*Seat)}
	sys.AddShow(show)
	show.Seats["B1"] = &Seat{ID: "B1", Status: SeatStatusAvailable, Price: 15.0}

	user := &User{ID: "U1"}

	// Book
	booking, err := sys.BookTickets(user, show, []string{"B1"})
	if err != nil {
		t.Fatalf("Unexpected booking failure: %v", err)
	}
	if booking.Status != BookingStatusPending {
		t.Errorf("Expected PENDING status, got %s", booking.Status)
	}

	// Confirm
	err = sys.ConfirmBooking(booking.ID)
	if err != nil {
		t.Errorf("Unexpected confirmation failure: %v", err)
	}
	if booking.Status != BookingStatusConfirmed {
		t.Errorf("Expected CONFIRMED status, got %s", booking.Status)
	}

	// Book another and cancel
	show.Seats["B2"] = &Seat{ID: "B2", Status: SeatStatusAvailable}
	booking2, _ := sys.BookTickets(user, show, []string{"B2"})
	
	err = sys.CancelBooking(booking2.ID)
	if err != nil {
		t.Errorf("Unexpected cancel failure: %v", err)
	}
	if booking2.Status != BookingStatusCancelled {
		t.Errorf("Expected CANCELLED status, got %s", booking2.Status)
	}
	if show.Seats["B2"].Status != SeatStatusAvailable {
		t.Errorf("Expected seat B2 to be AVAILABLE after cancellation")
	}
}
