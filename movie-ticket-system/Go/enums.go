package main

type SeatType string

const (
	SeatTypeNormal  SeatType = "NORMAL"
	SeatTypePremium SeatType = "PREMIUM"
)

type SeatStatus string

const (
	SeatStatusAvailable SeatStatus = "AVAILABLE"
	SeatStatusBooked    SeatStatus = "BOOKED"
)

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "PENDING"
	BookingStatusConfirmed BookingStatus = "CONFIRMED"
	BookingStatusCancelled BookingStatus = "CANCELLED"
)
