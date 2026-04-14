package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🚀 Initializing Single-Threaded Movie Ticket Booking System...")
	sys := GetMovieTicketBookingSystem()

	movie1 := &Movie{
		ID:          "M1",
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Duration:    148,
	}
	sys.AddMovie(movie1)

	theater1 := &Theater{
		ID:       "T1",
		Name:     "PVR Cinemas",
		Location: "Mumbai",
		Shows:    []*Show{},
	}
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

	for i := 1; i <= 5; i++ {
		seatId := fmt.Sprintf("A%d", i)
		show1.Seats[seatId] = &Seat{ID: seatId, Row: 1, Column: i, Type: SeatTypePremium, Price: 25.0, Status: SeatStatusAvailable}

		seatId = fmt.Sprintf("B%d", i)
		show1.Seats[seatId] = &Seat{ID: seatId, Row: 2, Column: i, Type: SeatTypeNormal, Price: 10.0, Status: SeatStatusAvailable}
	}

	user1 := &User{ID: "U1", Name: "Alice", Email: "alice@example.com"}
	user2 := &User{ID: "U2", Name: "Bob", Email: "bob@example.com"}
	user3 := &User{ID: "U3", Name: "Charlie", Email: "charlie@example.com"}

	fmt.Println("\n--- Starting Sequential Booking Simulation ---")
	fmt.Println("Alice attempts to book A1, A2.")
	
	bookingAlice, err := sys.BookTickets(user1, show1, []string{"A1", "A2"})
	if err != nil {
		fmt.Printf("❌ [Alice] Booking Failed: %v\n", err)
	} else {
		fmt.Printf("✅ [Alice] Booking Reserved (PENDING)! Booking ID: %s, Amount: $%.2f\n", bookingAlice.ID, bookingAlice.TotalPrice)
		sys.ConfirmBooking(bookingAlice.ID)
		fmt.Println("🎉 [Alice] Successfully CONFIRMED booking!")
	}

	fmt.Println("\nBob sequentially attempts to book the overlapping A1, A2...")
	bookingBob, err := sys.BookTickets(user2, show1, []string{"A1", "A2"})
	if err != nil {
		fmt.Printf("❌ [Bob] Booking Failed accurately (already taken): %v\n", err)
	} else {
		fmt.Printf("✅ [Bob] Booking Reserved! (This shouldn't happen)\n")
		sys.ConfirmBooking(bookingBob.ID)
	}

	fmt.Println("\nCharlie books B1, B2 and cancels...")
	charlieBooking, err := sys.BookTickets(user3, show1, []string{"B1", "B2"})
	if err == nil {
		fmt.Printf("✅ [Charlie] Reserved B1, B2. Status: PENDING.\n")
		fmt.Printf("🔄 [Charlie] Cancelling booking...\n")
		cancelErr := sys.CancelBooking(charlieBooking.ID)
		if cancelErr == nil {
			fmt.Printf("🗑️  [Charlie] Cancelled successfully. Seats reverted back to AVAILABLE!\n")
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
