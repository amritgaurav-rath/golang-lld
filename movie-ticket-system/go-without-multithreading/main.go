package main

import (
	"app/movie-ticket-system/go-without-multithreading/entities"
	"app/movie-ticket-system/go-without-multithreading/enums"
	"app/movie-ticket-system/go-without-multithreading/services"
    "app/movie-ticket-system/go-without-multithreading/observer"
    "app/movie-ticket-system/go-without-multithreading/strategy"
	"fmt"
	"time"
)

func main() {
	fmt.Println("🚀 Initializing Golang Movie Ticket System with Explicit Design Patterns")
	sys := services.GetMovieTicketBookingSystem()

	user1 := &entities.User{ID: "U1", Name: "Alice", Email: "alice@example.com"}
	user2 := &entities.User{ID: "U2", Name: "Bob", Email: "bob@example.com"}
	user3 := &entities.User{ID: "U3", Name: "Charlie", Email: "charlie@example.com"}

    // Observer Hooks
    sys.Subscribe("M1", &observer.UserObserver{UserName: user1.Name})
    sys.Subscribe("M1", &observer.UserObserver{UserName: user2.Name})

	movie1 := &entities.Movie{
		ID:          "M1",
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Duration:    148,
	}

	fmt.Println("\n--- Triggering Observer Push Notifications Natively ---")
    sys.AddMovie(movie1)

	theater1 := &entities.Theater{
		ID:       "T1",
		Name:     "PVR Cinemas",
		Location: "Mumbai",
		Shows:    []*entities.Show{},
	}
	sys.AddTheater(theater1)

	show1 := &entities.Show{
		ID:        "SH1",
		Movie:     movie1,
		Theater:   theater1,
		StartTime: time.Now().Add(1 * time.Hour),
		EndTime:   time.Now().Add(3 * time.Hour),
		Seats:     make(map[string]*entities.Seat),
	}
	sys.AddShow(show1)

	for i := 1; i <= 5; i++ {
		seatId := fmt.Sprintf("A%d", i)
		show1.Seats[seatId] = &entities.Seat{ID: seatId, Row: 1, Column: i, Type: enums.SeatTypePremium, Price: 25.0, Status: enums.SeatStatusAvailable}

		seatId = fmt.Sprintf("B%d", i)
		show1.Seats[seatId] = &entities.Seat{ID: seatId, Row: 2, Column: i, Type: enums.SeatTypeNormal, Price: 10.0, Status: enums.SeatStatusAvailable}
	}

	fmt.Println("\n--- Explicit Strategy Pattern Implementation ---")
	fmt.Println("Alice attempts to dynamically book A1, A2 using WEEKEND Pricing Strategy (+$10 / seat surge).")
	
	bookingAlice, err := sys.BookTickets(user1, show1, []string{"A1", "A2"}, &strategy.WeekendPricingStrategy{})
	if err != nil {
		fmt.Printf("❌ [Alice] Booking Failed: %v\n", err)
	} else {
		fmt.Printf("✅ [Alice] Booking Reserved (PENDING)! Booking ID: %s, Real-Time Amount: $%.2f\n", bookingAlice.ID, bookingAlice.TotalPrice)
		sys.ConfirmBooking(bookingAlice.ID)
	}

	fmt.Println("\nBob attempts to book overlapping Seats A1, A2 via WEEKDAY Pricing Strategy (Standard)...")
	bookingBob, err := sys.BookTickets(user2, show1, []string{"A1", "A2"}, &strategy.WeekdayPricingStrategy{})
	if err != nil {
		fmt.Printf("❌ [Bob] Booking Failed accurately (already natively taken): %v\n", err)
	} else {
		fmt.Printf("✅ [Bob] Booking Reserved! (This shouldn't happen)\n")
		sys.ConfirmBooking(bookingBob.ID)
	}

	fmt.Println("\nCharlie bookings B1, B2 safely on standard Weekday Pricing...")
	charlieBooking, err := sys.BookTickets(user3, show1, []string{"B1", "B2"}, &strategy.WeekdayPricingStrategy{})
	if err == nil {
		fmt.Printf("✅ [Charlie] Reserved B1, B2 dynamically. Total: $%.2f. Status: PENDING.\n", charlieBooking.TotalPrice)
	}
}
