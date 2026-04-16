import java.util.Arrays;

public class Main {
    public static void main(String[] args) {
        System.out.println("🚀 Starting Splitwise (Synchronous Java)");

        SplitwiseService service = SplitwiseService.getInstance();

        User alice = new User("U1", "Alice", "alice@example.com");
        User bob = new User("U2", "Bob", "bob@example.com");
        User charlie = new User("U3", "Charlie", "charlie@example.com");

        service.addUser(alice);
        service.addUser(bob);
        service.addUser(charlie);

        Group group = new Group("G1", "Vacation", Arrays.asList(alice, bob, charlie));
        service.addGroup(group);

        // Equal Split Expense
        Expense expense1 = new Expense(
            "EXP1",
            300.0,
            "Hotel",
            alice,
            Arrays.asList(new EqualSplit(alice), new EqualSplit(bob), new EqualSplit(charlie))
        );

        try {
            service.addExpense(group.getId(), expense1);
        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }

        // Percent Split Expense
        Expense expense2 = new Expense(
            "EXP2",
            100.0,
            "Dinner",
            bob,
            Arrays.asList(new PercentSplit(alice, 20), new PercentSplit(bob, 50), new PercentSplit(charlie, 30))
        );

        try {
            service.addExpense(group.getId(), expense2);
        } catch (Exception e) {
            System.out.println("Error: " + e.getMessage());
        }

        service.printBalances();

        // Settle Bob and Charlie
        System.out.println("Attempting to settle Charlie and Bob...");
        try {
            Transaction transaction = service.settleBalance(charlie.getId(), bob.getId());
            System.out.printf("✅ Settled! %s paid %s $%.2f\n", transaction.getSender().getName(), transaction.getReceiver().getName(), transaction.getAmount());
        } catch (Exception e) {
            System.out.println("Error settling: " + e.getMessage());
        }

        service.printBalances();
    }
}
