package parking
import "app/parking-lot/go-without-multithreading/entities"

type NearestFirstStrategy struct{}

func (s *NearestFirstStrategy) FindSpot(levels []*entities.Level, v entities.Vehicle) *entities.ParkingSpot {
	for _, level := range levels {
		if level.AvailableSpots[v.GetType()] > 0 {
			for _, spot := range level.Spots {
				if !spot.IsOccupied && spot.Type == v.GetType() {
					return spot
				}
			}
		}
	}
	return nil
}
