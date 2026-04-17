package fee
import "app/parking-lot/go-without-multithreading/entities"

type FeeStrategy interface {
	CalculateFee(v entities.Vehicle) float64
}
