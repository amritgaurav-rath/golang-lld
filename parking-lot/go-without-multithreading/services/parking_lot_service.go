package services

import (
	"app/parking-lot/go-without-multithreading/entities"
	"app/parking-lot/go-without-multithreading/enums"
	"app/parking-lot/go-without-multithreading/strategy/fee"
	"app/parking-lot/go-without-multithreading/strategy/parking"
	"fmt"
)

type ParkingLotSystem struct {
	Levels          []*entities.Level
	ParkingStrategy parking.ParkingStrategy
	FeeStrategy     fee.FeeStrategy
}

var instance *ParkingLotSystem

func GetInstance() *ParkingLotSystem {
	if instance == nil {
		instance = &ParkingLotSystem{
			Levels:          make([]*entities.Level, 0),
			ParkingStrategy: &parking.NearestFirstStrategy{}, // base configurations
			FeeStrategy:     &fee.FlatRateFeeStrategy{},     // base configurations
		}
	}
	return instance
}

func (p *ParkingLotSystem) SetLevels(levels []*entities.Level) {
	p.Levels = levels
}

func (p *ParkingLotSystem) SetParkingStrategy(strategy parking.ParkingStrategy) {
	p.ParkingStrategy = strategy
}

func (p *ParkingLotSystem) SetFeeStrategy(strategy fee.FeeStrategy) {
	p.FeeStrategy = strategy
}

func (p *ParkingLotSystem) ParkVehicle(v entities.Vehicle) (*entities.ParkingSpot, error) {
	// Let the mathematical Strategy search loop handle arrays completely decoupled
	spot := p.ParkingStrategy.FindSpot(p.Levels, v)
	if spot != nil {
		err := spot.Park(v)
		if err != nil {
			return nil, err
		}
		for _, level := range p.Levels {
			if level.ID == spot.LevelID {
				level.ModifyAvailability(v.GetType(), -1)
				break
			}
		}
		return spot, nil
	}
	return nil, fmt.Errorf("parking lot fully occupied for %v natively", v.GetType())
}

func (p *ParkingLotSystem) UnparkVehicle(spotID string) (entities.Vehicle, error) {
	for _, level := range p.Levels {
		for _, spot := range level.Spots {
			if spot.ID == spotID {
				if !spot.IsOccupied {
					return nil, fmt.Errorf("spot %s is completely empty mapped", spotID)
				}
				v := spot.Vehicle
				spot.Unpark()
				level.ModifyAvailability(v.GetType(), 1)

				// Bill dynamic checkout
				checkoutFee := p.FeeStrategy.CalculateFee(v)
				fmt.Printf("💳 Fee computed for %s: $%.2f\n", v.GetLicensePlate(), checkoutFee)

				return v, nil
			}
		}
	}
	return nil, fmt.Errorf("spot ID %s wasn't mapped natively", spotID)
}

func (p *ParkingLotSystem) PrintAvailability() {
	fmt.Println("\n--- Current Parking Availability ---")
	for _, l := range p.Levels {
		fmt.Printf("Level %d -> Motorcycles: %d, Cars: %d, Trucks: %d\n",
			l.ID, l.AvailableSpots[enums.Motorcycle], l.AvailableSpots[enums.Car], l.AvailableSpots[enums.Truck])
	}
	fmt.Println("------------------------------------")
}
