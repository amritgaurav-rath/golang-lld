public enum VehicleType {
    MOTORCYCLE("Motorcycle"),
    CAR("Car"),
    TRUCK("Truck"),
    UNKNOWN("Unknown");

    private final String name;

    VehicleType(String name) {
        this.name = name;
    }

    @Override
    public String toString() {
        return name;
    }
}
