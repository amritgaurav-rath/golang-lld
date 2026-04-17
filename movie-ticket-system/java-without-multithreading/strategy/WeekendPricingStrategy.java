package strategy;

public class WeekendPricingStrategy implements PricingStrategy {
    @Override
    public double calculatePrice(double basePrice) {
        return basePrice + 10.0; // Implicit $10 Surge
    }
}
