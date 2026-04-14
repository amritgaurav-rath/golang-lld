#ifndef LEVEL_HPP
#define LEVEL_HPP

#include <vector>
#include <map>
#include <mutex>
#include <memory>
#include <string>
#include "Models.hpp"

class Level {
private:
    int ID;
    std::vector<std::unique_ptr<ParkingSpot>> Spots;
    std::map<VehicleType, int> AvailableSpots;
    mutable std::mutex mtx;

public:
    Level(int id, int numMotorcycleSpots, int numCarSpots, int numTruckSpots);

    // Returns a pointer to the parked spot, or nullptr if unavailable
    ParkingSpot* ParkVehicle(Vehicle* v);

    // Unparks a vehicle and returns its pointer. Returns nullptr if spot not found or empty.
    Vehicle* UnparkVehicle(const std::string& spotID);

    std::map<VehicleType, int> GetAvailability() const;
    
    int GetID() const { return ID; }
};

#endif // LEVEL_HPP
