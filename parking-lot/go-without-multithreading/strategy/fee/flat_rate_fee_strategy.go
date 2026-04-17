package fee
import "app/parking-lot/go-without-multithreading/entities"

type FlatRateFeeStrategy struct{}

func (s *FlatRateFeeStrategy) CalculateFee(v entities.Vehicle) float64 {
	return 10.0
}
