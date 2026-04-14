package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestParkingLot_ConcurrentParkAndUnpark(t *testing.T) {
	// Initialize 1 level with 10 Motorcycle spots
	level1 := NewLevel(1, 10, 0, 0)
	lot := NewParkingLot([]*Level{level1})

	var wg sync.WaitGroup
	// Park 10 motorcycles concurrently
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			v := &Vehicle{
				LicensePlate: fmt.Sprintf("MOTO-%d", id),
				Type:         Motorcycle,
			}
			_, err := lot.ParkVehicle(v)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		}(i)
	}
	wg.Wait()

	// Verify all spots are taken
	avail := lot.GetTotalAvailability()
	if avail[Motorcycle] != 0 {
		t.Errorf("Expected 0 available motorcycle spots, got %d", avail[Motorcycle])
	}

	// Try parking an 11th motorcycle, should fail
	vFail := &Vehicle{LicensePlate: "MOTO-11", Type: Motorcycle}
	_, err := lot.ParkVehicle(vFail)
	if err == nil {
		t.Error("Expected error when parking in a full lot, got nil")
	}

	// Unpark all 10 motorcycles concurrently
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			spotID := fmt.Sprintf("L1-M%d", id)
			_, err := lot.UnparkVehicle(spotID)
			if err != nil {
				t.Errorf("Unexpected error unparking spot %s: %v", spotID, err)
			}
		}(i)
	}
	wg.Wait()

	// Verify all spots are available again
	avail = lot.GetTotalAvailability()
	if avail[Motorcycle] != 10 {
		t.Errorf("Expected 10 available motorcycle spots, got %d", avail[Motorcycle])
	}
}

func TestParkingLot_FullCapacity(t *testing.T) {
	level1 := NewLevel(1, 0, 1, 0)
	lot := NewParkingLot([]*Level{level1})

	v1 := &Vehicle{LicensePlate: "CAR-1", Type: Car}
	_, err1 := lot.ParkVehicle(v1)
	if err1 != nil {
		t.Errorf("Expected success, got error: %v", err1)
	}

	v2 := &Vehicle{LicensePlate: "CAR-2", Type: Car}
	_, err2 := lot.ParkVehicle(v2)
	if err2 == nil {
		t.Error("Expected failure due to full capacity, got success")
	}
}
