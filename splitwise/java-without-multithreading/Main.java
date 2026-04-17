import java.util.Arrays;
import entities.*;
import strategy.*;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        System.out.println("🚀 Starting Splitwise (Perfect AshishPS1 Repository Clone)");

        SplitwiseService service = SplitwiseService.getInstance();

        User u1 = service.addUser("U1", "John", "john@example.com");
        User u2 = service.addUser("U2", "Jane", "jane@example.com");
        User u3 = service.addUser("U3", "Bob", "bob@example.com");

        List<User> participants = Arrays.asList(u1, u2, u3);
        Group group = service.addGroup("G1", "Trip to Paris", participants);

        // Equal Split
        Expense.ExpenseBuilder equalBuilder = new Expense.ExpenseBuilder()
                .setId("EXP1")
                .setDescription("Dinner")
                .setAmount(300.0)
                .setPaidBy(u1)
                .setParticipants(participants)
                .setSplitStrategy(new EqualSplitStrategy());
        service.createExpense(equalBuilder);

        // Exact Split
        Expense.ExpenseBuilder exactBuilder = new Expense.ExpenseBuilder()
                .setId("EXP2")
                .setDescription("Cab")
                .setAmount(100.0)
                .setPaidBy(u2)
                .setParticipants(participants)
                .setSplitValues(Arrays.asList(20.0, 30.0, 50.0)) // Sums to 100
                .setSplitStrategy(new ExactSplitStrategy());
        service.createExpense(exactBuilder);

        // Percentage Split
        Expense.ExpenseBuilder percentBuilder = new Expense.ExpenseBuilder()
                .setId("EXP3")
                .setDescription("Hotel")
                .setAmount(200.0)
                .setPaidBy(u3)
                .setParticipants(participants)
                .setSplitValues(Arrays.asList(10.0, 40.0, 50.0)) // Sums to 100%
                .setSplitStrategy(new PercentageSplitStrategy());
        service.createExpense(percentBuilder);

        System.out.println("\n--- Current Balances ---");
        service.showBalanceSheet(u1.getId());
        service.showBalanceSheet(u2.getId());
        service.showBalanceSheet(u3.getId());

        System.out.println("\n--- Simplifying Group Debts ---");
        List<Transaction> transactions = service.simplifyGroupDebts(group.getId());
        for (Transaction tx : transactions) {
            System.out.printf("Simplify Settlement Route: %s should precisely pay %s $%.2f\n", tx.getFrom().getName(), tx.getTo().getName(), tx.getAmount());
        }
    }
}
