#include "ParkingLot.hpp"
#include <iostream>

ParkingLot::ParkingLot(const std::vector<std::shared_ptr<Level>>& levels) : Levels(levels) {}

ParkingSpot* ParkingLot::ParkVehicle(Vehicle* v) {
    for (const auto& level : Levels) {
        auto avail = level->GetAvailability();
        if (avail[v->Type] > 0) {
            ParkingSpot* spot = level->ParkVehicle(v);
            if (spot != nullptr) {
                return spot;
            }
        }
    }
    return nullptr;
}

Vehicle* ParkingLot::UnparkVehicle(const std::string& spotID) {
    for (const auto& level : Levels) {
        Vehicle* v = level->UnparkVehicle(spotID);
        if (v != nullptr) {
            return v;
        }
    }
    return nullptr;
}

std::map<VehicleType, int> ParkingLot::GetTotalAvailability() const {
    std::map<VehicleType, int> totalAvail;
    for (const auto& level : Levels) {
        auto avail = level->GetAvailability();
        for (const auto& pair : avail) {
            totalAvail[pair.first] += pair.second;
        }
    }
    return totalAvail;
}

void ParkingLot::PrintAvailability() const {
    std::cout << "--- Current Parking Availability ---\n";
    auto avail = GetTotalAvailability();
    for (const auto& pair : avail) {
        if (pair.second > 0 || pair.first != VehicleType::Unknown) {
            std::cout << VehicleTypeToString(pair.first) << " Spots: " << pair.second << "\n";
        }
    }
    std::cout << "------------------------------------\n";
}
