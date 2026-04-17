package strategy;
import entities.Split;
import entities.User;
import java.util.ArrayList;
import java.util.List;

public class ExactSplitStrategy implements SplitStrategy {
    public List<Split> calculateSplits(double amount, User paidBy, List<User> participants, List<Double> splitValues) {
        if (splitValues == null || splitValues.size() != participants.size()) {
            throw new IllegalArgumentException("Invalid split values provided");
        }
        
        double total = 0.0;
        for (double val : splitValues) total += val;
        
        if (Math.abs(total - amount) > 0.01) {
            throw new IllegalArgumentException("Exact splits total does not match expense amount");
        }
        
        List<Split> splits = new ArrayList<>();
        for (int i = 0; i < participants.size(); i++) {
            splits.add(new Split(participants.get(i), splitValues.get(i)));
        }
        return splits;
    }
}
