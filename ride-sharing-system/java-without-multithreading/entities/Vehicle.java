package entities;
import enums.VehicleType;

public class Vehicle {
    private String ownerId;
    private String registrationNo;
    private VehicleType type;

    public Vehicle(String ownerId, String registrationNo, VehicleType type) {
        this.ownerId = ownerId;
        this.registrationNo = registrationNo;
        this.type = type;
    }

    public String getOwnerId() { return ownerId; }
    public String getRegistrationNo() { return registrationNo; }
    public VehicleType getType() { return type; }
}
