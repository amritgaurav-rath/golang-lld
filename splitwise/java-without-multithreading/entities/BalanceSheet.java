package entities;
import java.util.HashMap;
import java.util.Map;

public class BalanceSheet {
    private final Map<User, Double> balances;

    public BalanceSheet() {
        this.balances = new HashMap<>();
    }

    public Map<User, Double> getBalances() {
        return balances;
    }

    public void adjustBalance(User user, double amount) {
        balances.put(user, balances.getOrDefault(user, 0.0) + amount);
    }

    public void showBalances(String userName) {
        for (Map.Entry<User, Double> entry : balances.entrySet()) {
            User user = entry.getKey();
            double amount = entry.getValue();
            if (amount < 0) {
                System.out.printf("%s owes %s: $%.2f\n", userName, user.getName(), Math.abs(amount));
            }
        }
    }
}
