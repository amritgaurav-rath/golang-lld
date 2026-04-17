/**
 * Represents a definitive settlement operation between two nodes in the ledger.
 */
package entities;
public class Transaction {
    private String id;
    private User from;
    private User to;
    private double amount;

    public Transaction(String id, User from, User to, double amount) {
        this.id = id;
        this.from = from;
        this.to = to;
        this.amount = amount;
    }

    public String getId() { return id; }
    public User getFrom() { return from; }
    public User getTo() { return to; }
    public double getAmount() { return amount; }
}
