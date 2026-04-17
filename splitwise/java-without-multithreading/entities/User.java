package entities;
public class User {
    private final String id;
    private final String name;
    private final String email;
    private final BalanceSheet balanceSheet;

    public User(String id, String name, String email) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.balanceSheet = new BalanceSheet();
    }

    public String getId() { return id; }
    public String getName() { return name; }
    public String getEmail() { return email; }
    public BalanceSheet getBalanceSheet() { return balanceSheet; }

    @Override
    public int hashCode() {
        return id.hashCode();
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) return true;
        if (obj == null || getClass() != obj.getClass()) return false;
        User user = (User) obj;
        return id.equals(user.id);
    }
}
