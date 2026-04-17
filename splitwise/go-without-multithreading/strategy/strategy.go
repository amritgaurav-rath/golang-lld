package strategy

import (
	"app/splitwise/go-without-multithreading/entities"
	"fmt"
	"math"
)

// SplitStrategy is the pure OCP engine interface. Calculations are delegated here.
type SplitStrategy interface {
	CalculateSplits(amount float64, participants []*entities.User, splitValues []float64) ([]entities.Split, error)
}

// EqualSplitStrategy divides an expense evenly.
type EqualSplitStrategy struct{}

func (e *EqualSplitStrategy) CalculateSplits(amount float64, participants []*entities.User, splitValues []float64) ([]entities.Split, error) {
	if len(participants) == 0 {
		return nil, fmt.Errorf("participants required")
	}

	amountPerUser := math.Round((amount/float64(len(participants)))*100) / 100
	var splits []entities.Split
	for _, u := range participants {
		splits = append(splits, entities.Split{User: u, Amount: amountPerUser})
	}
	return splits, nil
}

// PercentSplitStrategy partitions using explicit fraction mapping.
type PercentSplitStrategy struct{}

func (p *PercentSplitStrategy) CalculateSplits(amount float64, participants []*entities.User, splitValues []float64) ([]entities.Split, error) {
	if len(participants) == 0 || len(splitValues) != len(participants) {
		return nil, fmt.Errorf("split values must match participants exactly")
	}

	var total float64
	for _, v := range splitValues {
		total += v
	}
	if math.Abs(total-100.0) > 0.01 {
		return nil, fmt.Errorf("percentage splits must exactly equal 100")
	}

	var splits []entities.Split
	for i, u := range participants {
		calcAmount := math.Round((amount*splitValues[i]/100.0)*100) / 100
		splits = append(splits, entities.Split{User: u, Amount: calcAmount})
	}
	return splits, nil
}
