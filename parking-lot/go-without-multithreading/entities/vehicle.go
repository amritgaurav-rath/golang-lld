package entities
import "app/parking-lot/go-without-multithreading/enums"

type Vehicle interface {
	GetLicensePlate() string
	GetType() enums.VehicleType
}

type BaseVehicle struct {
	LicensePlate string
}

func (b *BaseVehicle) GetLicensePlate() string { return b.LicensePlate }

type CarVehicle struct { BaseVehicle }
func (c *CarVehicle) GetType() enums.VehicleType { return enums.Car }

type MotorcycleVehicle struct { BaseVehicle }
func (m *MotorcycleVehicle) GetType() enums.VehicleType { return enums.Motorcycle }

type TruckVehicle struct { BaseVehicle }
func (t *TruckVehicle) GetType() enums.VehicleType { return enums.Truck }
