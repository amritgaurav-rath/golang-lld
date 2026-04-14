#ifndef PARKING_LOT_HPP
#define PARKING_LOT_HPP

#include "Level.hpp"
#include <vector>
#include <memory>
#include <map>

class ParkingLot {
private:
    std::vector<std::shared_ptr<Level>> Levels;

public:
    ParkingLot(const std::vector<std::shared_ptr<Level>>& levels);

    ParkingSpot* ParkVehicle(Vehicle* v);
    Vehicle* UnparkVehicle(const std::string& spotID);
    std::map<VehicleType, int> GetTotalAvailability() const;
    void PrintAvailability() const;
};

#endif // PARKING_LOT_HPP
