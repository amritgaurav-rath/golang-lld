package main

import (
	"fmt"
	"sync"
	"time"
)

// MovieTicketBookingDemo class equivalent: demonstrating the usage of the system
func main() {
	fmt.Println("🚀 Initializing Movie Ticket Booking System...")
	sys := GetMovieTicketBookingSystem()

	// 1. Setup Movies
	movie1 := &Movie{
		ID:          "M1",
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Duration:    148,
	}
	sys.AddMovie(movie1)

	// 2. Setup Theater
	theater1 := &Theater{
		ID:       "T1",
		Name:     "PVR Cinemas",
		Location: "Mumbai",
		Shows:    []*Show{},
	}
	sys.AddTheater(theater1)

	// 3. Setup Show and Seats
	show1 := &Show{
		ID:        "SH1",
		Movie:     movie1,
		Theater:   theater1,
		StartTime: time.Now().Add(1 * time.Hour),
		EndTime:   time.Now().Add(3 * time.Hour),
		Seats:     make(map[string]*Seat),
	}
	sys.AddShow(show1)

	// Configure Seats directly in Show
	for i := 1; i <= 5; i++ {
		seatId := fmt.Sprintf("A%d", i)
		show1.Seats[seatId] = &Seat{ID: seatId, Row: 1, Column: i, Type: SeatTypePremium, Price: 25.0, Status: SeatStatusAvailable}

		seatId = fmt.Sprintf("B%d", i)
		show1.Seats[seatId] = &Seat{ID: seatId, Row: 2, Column: i, Type: SeatTypeNormal, Price: 10.0, Status: SeatStatusAvailable}
	}

	// 4. Setup Users
	user1 := &User{ID: "U1", Name: "Alice", Email: "alice@example.com"}
	user2 := &User{ID: "U2", Name: "Bob", Email: "bob@example.com"}
	user3 := &User{ID: "U3", Name: "Charlie", Email: "charlie@example.com"}

	// 5. Simulate Concurrent Bookings
	fmt.Println("\n--- Starting Concurrent Booking Simulation ---")
	fmt.Println("Alice and Bob attempt to book intersecting seats concurrently [A1, A2].")

	var wg sync.WaitGroup
	wg.Add(2)

	seatRequest := []string{"A1", "A2"}
	var successfulBookingId string
	var bookingMutex sync.Mutex

	go func() {
		defer wg.Done()
		fmt.Println("[Alice] attempting to reserve A1, A2...")
		booking, err := sys.BookTickets(user1, show1, seatRequest)
		if err != nil {
			fmt.Printf("❌ [Alice] Booking Failed: %v\n", err)
		} else {
			fmt.Printf("✅ [Alice] Booking Reserved (PENDING)! Booking ID: %s, Amount: $%.2f\n", booking.ID, booking.TotalPrice)
			bookingMutex.Lock()
			successfulBookingId = booking.ID
			bookingMutex.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("[Bob] attempting to reserve A1, A2...")
		booking, err := sys.BookTickets(user2, show1, seatRequest)
		if err != nil {
			fmt.Printf("❌ [Bob] Booking Failed: %v\n", err)
		} else {
			fmt.Printf("✅ [Bob] Booking Reserved (PENDING)! Booking ID: %s, Amount: $%.2f\n", booking.ID, booking.TotalPrice)
			bookingMutex.Lock()
			successfulBookingId = booking.ID
			bookingMutex.Unlock()
		}
	}()

	wg.Wait()

	// 6. Demonstrate Confirm and Cancel
	fmt.Println("\n--- Payment & Concurrency resolution ---")
	
	if successfulBookingId != "" {
		fmt.Printf("System attempting to confirm the victorious booking %s...\n", successfulBookingId)
		err := sys.ConfirmBooking(successfulBookingId)
		if err != nil {
			fmt.Printf("❌ Failed to confirm booking %s: %v\n", successfulBookingId, err)
		} else {
			fmt.Printf("🎉 Successfully CONFIRMED booking %s!\n", successfulBookingId)
		}
	}

	// Charlie books and then cancels
	fmt.Println("\n[Charlie] books B1, B2 and cancels...")
	charlieBooking, err := sys.BookTickets(user3, show1, []string{"B1", "B2"})
	if err == nil {
		fmt.Printf("✅ [Charlie] Reserved B1, B2. Status: PENDING.\n")
		fmt.Printf("🔄 [Charlie] Cancelling booking...\n")
		cancelErr := sys.CancelBooking(charlieBooking.ID)
		if cancelErr == nil {
			fmt.Printf("🗑️  [Charlie] Cancelled successfully. Seats reverted back to AVAILABLE!\n")
		} else {
			fmt.Printf("❌ Cancel failed: %v\n", cancelErr)
		}
	}

	fmt.Println("\n--- Final Seat Availability Summary ---")
	bookedSeats := 0
	availableSeats := 0
	for _, seat := range show1.Seats {
		if seat.Status == SeatStatusBooked {
			bookedSeats++
		} else {
			availableSeats++
		}
	}
	fmt.Printf("Initial Total Seats: 10\n")
	fmt.Printf("Seats Successfully Booked: %d\n", bookedSeats)
	fmt.Printf("Seats Remaning Available: %d\n", availableSeats)
}
