#ifndef MODELS_HPP
#define MODELS_HPP

#include <string>

enum class VehicleType {
    Motorcycle,
    Car,
    Truck,
    Unknown
};

inline std::string VehicleTypeToString(VehicleType v) {
    switch (v) {
        case VehicleType::Motorcycle: return "Motorcycle";
        case VehicleType::Car:        return "Car";
        case VehicleType::Truck:      return "Truck";
        default:                      return "Unknown";
    }
}

struct Vehicle {
    std::string LicensePlate;
    VehicleType Type;
};

class ParkingSpot {
public:
    std::string ID;
    VehicleType Type;
    int LevelID;
    bool IsOccupied;
    Vehicle* ParkedVehicle;

    ParkingSpot(std::string id, VehicleType type, int levelID)
        : ID(id), Type(type), LevelID(levelID), IsOccupied(false), ParkedVehicle(nullptr) {}

    bool Park(Vehicle* v) {
        if (IsOccupied) return false;
        ParkedVehicle = v;
        IsOccupied = true;
        return true;
    }

    void Unpark() {
        ParkedVehicle = nullptr;
        IsOccupied = false;
    }
};

#endif // MODELS_HPP
