package strategy;
import entities.Split;
import entities.User;
import java.util.ArrayList;
import java.util.List;

public class PercentageSplitStrategy implements SplitStrategy {
    public List<Split> calculateSplits(double amount, User paidBy, List<User> participants, List<Double> splitValues) {
        if (splitValues == null || splitValues.size() != participants.size()) {
            throw new IllegalArgumentException("Invalid percentage values provided");
        }
        
        double totalPercent = 0.0;
        for (double val : splitValues) totalPercent += val;
        
        if (Math.abs(totalPercent - 100.0) > 0.01) {
            throw new IllegalArgumentException("Total percentage does not add up to 100");
        }
        
        List<Split> splits = new ArrayList<>();
        for (int i = 0; i < participants.size(); i++) {
            double splitAmount = Math.round((amount * splitValues.get(i) / 100.0) * 100.0) / 100.0;
            splits.add(new Split(participants.get(i), splitAmount));
        }
        return splits;
    }
}
