#include "Level.hpp"

Level::Level(int id, int numMotorcycleSpots, int numCarSpots, int numTruckSpots) : ID(id) {
    AvailableSpots[VehicleType::Motorcycle] = numMotorcycleSpots;
    AvailableSpots[VehicleType::Car] = numCarSpots;
    AvailableSpots[VehicleType::Truck] = numTruckSpots;

    int spotNum = 1;
    for (int i = 0; i < numMotorcycleSpots; ++i) {
        Spots.push_back(std::make_unique<ParkingSpot>("L" + std::to_string(id) + "-M" + std::to_string(spotNum++), VehicleType::Motorcycle, id));
    }
    
    spotNum = 1;
    for (int i = 0; i < numCarSpots; ++i) {
        Spots.push_back(std::make_unique<ParkingSpot>("L" + std::to_string(id) + "-C" + std::to_string(spotNum++), VehicleType::Car, id));
    }
    
    spotNum = 1;
    for (int i = 0; i < numTruckSpots; ++i) {
        Spots.push_back(std::make_unique<ParkingSpot>("L" + std::to_string(id) + "-T" + std::to_string(spotNum++), VehicleType::Truck, id));
    }
}

ParkingSpot* Level::ParkVehicle(Vehicle* v) {
    std::lock_guard<std::mutex> lock(mtx);

    if (AvailableSpots[v->Type] <= 0) {
        return nullptr;
    }

    for (const auto& spot : Spots) {
        if (!spot->IsOccupied && spot->Type == v->Type) {
            if (spot->Park(v)) {
                AvailableSpots[v->Type]--;
                return spot.get();
            }
        }
    }
    return nullptr;
}

Vehicle* Level::UnparkVehicle(const std::string& spotID) {
    std::lock_guard<std::mutex> lock(mtx);

    for (const auto& spot : Spots) {
        if (spot->ID == spotID) {
            if (!spot->IsOccupied) {
                return nullptr;
            }
            Vehicle* v = spot->ParkedVehicle;
            spot->Unpark();
            AvailableSpots[v->Type]++;
            return v;
        }
    }
    return nullptr;
}

std::map<VehicleType, int> Level::GetAvailability() const {
    std::lock_guard<std::mutex> lock(mtx);
    return AvailableSpots;
}
