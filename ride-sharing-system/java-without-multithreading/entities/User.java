package entities;

public class User {
    private String id;
    private String name;
    private int takenRides;
    private int offeredRides;

    public User(String id, String name) {
        this.id = id;
        this.name = name;
        this.takenRides = 0;
        this.offeredRides = 0;
    }

    public String getId() { return id; }
    public String getName() { return name; }
    public int getTakenRides() { return takenRides; }
    public int getOfferedRides() { return offeredRides; }
    
    public void incrementTakenRides() { this.takenRides++; }
    public void incrementOfferedRides() { this.offeredRides++; }
}
