package fee

import (
	"app/parking-lot/go-without-multithreading/entities"
	"app/parking-lot/go-without-multithreading/enums"
)

type VehicleBasedFeeStrategy struct{}

func (s *VehicleBasedFeeStrategy) CalculateFee(v entities.Vehicle) float64 {
	switch v.GetType() {
	case enums.Motorcycle: return 5.0
	case enums.Car: return 10.0
	case enums.Truck: return 20.0
	default: return 10.0
	}
}
