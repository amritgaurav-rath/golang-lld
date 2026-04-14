package main

import (
	"sync"
	"time"
)

// User class represents a user of the booking system
type User struct {
	ID    string
	Name  string
	Email string
}

// Movie class represents a movie with properties such as ID, title, description, and duration
type Movie struct {
	ID          string
	Title       string
	Description string
	Duration    int // in minutes
}

// Seat class represents a seat in a show, with properties such as ID, row, column, type, price, and status
type Seat struct {
	ID     string
	Row    int
	Column int
	Type   SeatType
	Price  float64
	Status SeatStatus
}

// Show class represents a movie show in a theater
type Show struct {
	sync.Mutex
	ID        string
	Movie     *Movie
	Theater   *Theater
	StartTime time.Time
	EndTime   time.Time
	Seats     map[string]*Seat
}

// Theater class represents a theater with properties such as ID, name, location, and a list of shows
type Theater struct {
	ID       string
	Name     string
	Location string
	Shows    []*Show
}

// Booking class represents a booking made by a user
type Booking struct {
	ID         string
	User       *User
	Show       *Show
	Seats      []*Seat
	TotalPrice float64
	Status     BookingStatus
}
