package main

// Request represents a single passenger request for elevator service
type Request struct {
	SourceFloor      int
	DestinationFloor int
	Passengers       int
}
