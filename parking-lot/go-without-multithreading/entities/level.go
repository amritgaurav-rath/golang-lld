package entities
import (
    "app/parking-lot/go-without-multithreading/enums"
    "fmt"
)

type Level struct {
	ID             int
	Spots          []*ParkingSpot
	AvailableSpots map[enums.VehicleType]int
}

func NewLevel(id int, numMotorcycleSpots, numCarSpots, numTruckSpots int) *Level {
	level := &Level{
		ID:             id,
		Spots:          make([]*ParkingSpot, 0, numMotorcycleSpots+numCarSpots+numTruckSpots),
		AvailableSpots: make(map[enums.VehicleType]int),
	}

	level.AvailableSpots[enums.Motorcycle] = numMotorcycleSpots
	level.AvailableSpots[enums.Car] = numCarSpots
	level.AvailableSpots[enums.Truck] = numTruckSpots

	spotNum := 1
	for i := 0; i < numMotorcycleSpots; i++ {
		level.Spots = append(level.Spots, &ParkingSpot{ ID: fmt.Sprintf("L%d-M%d", id, spotNum), Type: enums.Motorcycle, LevelID: id })
		spotNum++
	}
	spotNum = 1
	for i := 0; i < numCarSpots; i++ {
		level.Spots = append(level.Spots, &ParkingSpot{ ID: fmt.Sprintf("L%d-C%d", id, spotNum), Type: enums.Car, LevelID: id })
		spotNum++
	}
	spotNum = 1
	for i := 0; i < numTruckSpots; i++ {
		level.Spots = append(level.Spots, &ParkingSpot{ ID: fmt.Sprintf("L%d-T%d", id, spotNum), Type: enums.Truck, LevelID: id })
		spotNum++
	}

	return level
}

// Simple modifier decoupled from execution math
func (l *Level) ModifyAvailability(vType enums.VehicleType, delta int) {
    l.AvailableSpots[vType] += delta
}
