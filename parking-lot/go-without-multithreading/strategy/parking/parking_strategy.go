package parking
import "app/parking-lot/go-without-multithreading/entities"

type ParkingStrategy interface {
	FindSpot(levels []*entities.Level, v entities.Vehicle) *entities.ParkingSpot
}
