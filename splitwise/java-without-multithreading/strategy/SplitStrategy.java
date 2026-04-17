package strategy;
import entities.Split;
import entities.User;
import java.util.List;

public interface SplitStrategy {
    List<Split> calculateSplits(double amount, User paidBy, List<User> participants, List<Double> splitValues);
}
