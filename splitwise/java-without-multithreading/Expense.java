import java.util.List;

public class Expense {
    private String id;
    private double amount;
    private String description;
    private User paidBy;
    private List<Split> splits;

    public Expense(String id, double amount, String description, User paidBy, List<Split> splits) {
        this.id = id;
        this.amount = amount;
        this.description = description;
        this.paidBy = paidBy;
        this.splits = splits;
    }

    public String getId() {
        return id;
    }

    public double getAmount() {
        return amount;
    }

    public String getDescription() {
        return description;
    }

    public User getPaidBy() {
        return paidBy;
    }

    public List<Split> getSplits() {
        return splits;
    }
}
