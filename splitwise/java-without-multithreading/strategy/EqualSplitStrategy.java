package strategy;
import entities.Split;
import entities.User;
import java.util.ArrayList;
import java.util.List;

public class EqualSplitStrategy implements SplitStrategy {
    public List<Split> calculateSplits(double amount, User paidBy, List<User> participants, List<Double> splitValues) {
        double splitAmount = Math.round((amount / participants.size()) * 100.0) / 100.0;
        List<Split> splits = new ArrayList<>();
        for (User user : participants) {
            splits.add(new Split(user, splitAmount));
        }
        return splits;
    }
}
