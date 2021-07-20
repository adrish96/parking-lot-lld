package parkinglot

import (
	"github.com/adrish96/parking-lot-lld/consts"
	"github.com/adrish96/parking-lot-lld/models"
)

// ParkingLotList in-memory store of Parking Lots
var ParkingLotList []models.ParkingLot

type ParkingLot interface {
	NewParkingLot(id string, name string, address string, numFloors int64, ratecard models.Ratecard, lotfloors []models.ParkingLotFloor) models.ParkingLot
	NewParkingLotFloor(floorNum int64, capacity map[consts.VehicleType]int64, parkingSpots []models.ParkingSpot) models.ParkingLotFloor
	NewParkingSpot(id string) models.ParkingSpot
	IsFull(parkingLotID string) bool

}

type ParkingLotImpl struct {
}

func (p ParkingLotImpl) NewParkingSpot(id string) models.ParkingSpot {
	return models.NewParkingLotSpot(id)
}

func (p ParkingLotImpl) NewParkingLotFloor(floorNum int64, capacity map[consts.VehicleType]int64, parkingSpots []models.ParkingSpot) models.ParkingLotFloor {
	return models.NewParkingLotFloor(floorNum, capacity, parkingSpots)
}

func (p ParkingLotImpl) NewParkingLot(id string, name string, address string, numFloors int64, ratecard models.Ratecard, lotfloors []models.ParkingLotFloor) models.ParkingLot {
	lot := models.NewParkingLot(id, name, address, numFloors, ratecard, lotfloors)
	ParkingLotList = append(ParkingLotList, lot)
	return lot
}

func (p ParkingLotImpl) IsFull(parkingLotID string) bool {
	for _, lot := range ParkingLotList {
		if lot.ID == parkingLotID {
			for _, floor := range lot.ParkingLotFloors {
				for _, spots := range floor.ParkingLotSpots {
					if spots.IsTaken == false {
						return false
					}
				}
			}
		}
	}
	return true
}
