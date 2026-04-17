package main

// Request represents a user request for an elevator
type Request struct {
	SourceFloor      int
	DestinationFloor int
	Passengers       int
}
