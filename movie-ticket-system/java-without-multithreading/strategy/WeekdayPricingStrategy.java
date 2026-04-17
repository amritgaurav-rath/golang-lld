package strategy;

public class WeekdayPricingStrategy implements PricingStrategy {
    @Override
    public double calculatePrice(double basePrice) {
        return basePrice; // No surge mapping statically
    }
}
