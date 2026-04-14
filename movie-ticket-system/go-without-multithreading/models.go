package main

import "time"

type User struct {
	ID    string
	Name  string
	Email string
}

type Movie struct {
	ID          string
	Title       string
	Description string
	Duration    int
}

type Seat struct {
	ID     string
	Row    int
	Column int
	Type   SeatType
	Price  float64
	Status SeatStatus
}

type Show struct {
	// Mutex removed for single-threaded executing version
	ID        string
	Movie     *Movie
	Theater   *Theater
	StartTime time.Time
	EndTime   time.Time
	Seats     map[string]*Seat
}

type Theater struct {
	ID       string
	Name     string
	Location string
	Shows    []*Show
}

type Booking struct {
	ID         string
	User       *User
	Show       *Show
	Seats      []*Seat
	TotalPrice float64
	Status     BookingStatus
}
