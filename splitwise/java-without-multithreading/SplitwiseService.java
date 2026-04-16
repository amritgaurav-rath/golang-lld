import java.util.HashMap;
import java.util.Map;

public class SplitwiseService {
    private static SplitwiseService instance;

    private Map<String, User> users;
    private Map<String, Group> groups;
    // Balances maps [UserA_ID][UserB_ID] = Amount that User A owes to User B.
    private Map<String, Map<String, Double>> balances;

    private SplitwiseService() {
        users = new HashMap<>();
        groups = new HashMap<>();
        balances = new HashMap<>();
    }

    public static SplitwiseService getInstance() {
        if (instance == null) {
            instance = new SplitwiseService();
        }
        return instance;
    }

    public void addUser(User u) {
        users.put(u.getId(), u);
        balances.put(u.getId(), new HashMap<>());
    }

    public void addGroup(Group g) {
        groups.put(g.getId(), g);
    }

    public void addExpense(String groupId, Expense expense) throws Exception {
        Group group = groups.get(groupId);
        if (group == null) {
            throw new Exception("group " + groupId + " does not exist");
        }

        Split typeCheck = expense.getSplits().get(0);

        if (typeCheck instanceof EqualSplit) {
            double amountPerUser = expense.getAmount() / expense.getSplits().size();
            amountPerUser = Math.round(amountPerUser * 100.0) / 100.0;
            for (Split split : expense.getSplits()) {
                split.setAmount(amountPerUser);
            }
        } else if (typeCheck instanceof PercentSplit) {
            for (Split split : expense.getSplits()) {
                PercentSplit percentSplit = (PercentSplit) split;
                double amount = (expense.getAmount() * percentSplit.getPercent()) / 100.0;
                split.setAmount(Math.round(amount * 100.0) / 100.0);
            }
        } else if (typeCheck instanceof ExactSplit) {
            double total = 0.0;
            for (Split split : expense.getSplits()) {
                total += split.getAmount();
            }
            if (Math.abs(total - expense.getAmount()) > 0.01) {
                throw new Exception("exact splits total does not match expense amount");
            }
        }

        group.addExpense(expense);

        String paidBy = expense.getPaidBy().getId();
        for (Split split : expense.getSplits()) {
            String splitUser = split.getUser().getId();
            if (paidBy.equals(splitUser)) {
                continue;
            }

            Map<String, Double> splitUserBalances = balances.get(splitUser);
            splitUserBalances.put(paidBy, splitUserBalances.getOrDefault(paidBy, 0.0) + split.getAmount());

            Map<String, Double> paidByBalances = balances.get(paidBy);
            paidByBalances.put(splitUser, paidByBalances.getOrDefault(splitUser, 0.0) - split.getAmount());
        }
    }

    public Transaction settleBalance(String userA, String userB) throws Exception {
        double amountOwed = balances.get(userA).getOrDefault(userB, 0.0);
        if (amountOwed == 0) {
            throw new Exception("no balance to settle between " + userA + " and " + userB);
        }

        String sender, receiver;
        double amount;

        if (amountOwed > 0) {
            sender = userA;
            receiver = userB;
            amount = amountOwed;
        } else {
            sender = userB;
            receiver = userA;
            amount = -amountOwed;
        }

        balances.get(userA).put(userB, 0.0);
        balances.get(userB).put(userA, 0.0);

        return new Transaction(
            "TRX-" + sender + "-" + receiver,
            users.get(sender),
            users.get(receiver),
            amount
        );
    }

    public void printBalances() {
        System.out.println("--- System Balances ---");
        for (String userA : balances.keySet()) {
            Map<String, Double> owesMap = balances.get(userA);
            for (String userB : owesMap.keySet()) {
                double amount = owesMap.get(userB);
                if (amount > 0) {
                    System.out.printf("%s owes %s: $%.2f\n", users.get(userA).getName(), users.get(userB).getName(), amount);
                }
            }
        }
        System.out.println("-----------------------");
    }
}
