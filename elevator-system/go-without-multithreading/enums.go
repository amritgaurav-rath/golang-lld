package main

type Direction int

const (
	IDLE Direction = iota
	UP
	DOWN
)

func (d Direction) String() string {
	switch d {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	default:
		return "IDLE"
	}
}
