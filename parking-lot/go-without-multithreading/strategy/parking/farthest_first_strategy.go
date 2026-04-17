package parking
import "app/parking-lot/go-without-multithreading/entities"

type FarthestFirstStrategy struct{}

func (s *FarthestFirstStrategy) FindSpot(levels []*entities.Level, v entities.Vehicle) *entities.ParkingSpot {
	for i := len(levels) - 1; i >= 0; i-- {
		level := levels[i]
		if level.AvailableSpots[v.GetType()] > 0 {
			spots := level.Spots
			for j := len(spots) - 1; j >= 0; j-- {
				spot := spots[j]
				if !spot.IsOccupied && spot.Type == v.GetType() {
					return spot
				}
			}
		}
	}
	return nil
}
