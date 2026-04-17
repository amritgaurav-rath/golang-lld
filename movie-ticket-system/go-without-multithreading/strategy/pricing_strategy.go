package strategy

// PricingStrategy Interface explicitly isolating fractional evaluation nodes.
type PricingStrategy interface {
	CalculatePrice(basePrice float64) float64
}

// WeekdayPricingStrategy returns pure unaltered values
type WeekdayPricingStrategy struct{}
func (s *WeekdayPricingStrategy) CalculatePrice(basePrice float64) float64 {
	return basePrice
}

// WeekendPricingStrategy pushes a unified peak markup natively 
type WeekendPricingStrategy struct{}
func (s *WeekendPricingStrategy) CalculatePrice(basePrice float64) float64 {
	return basePrice + 10.0
}
