package models

import "github.com/adrish96/parking-lot-lld/consts"

type ParkingLot struct {
	ID string
	Name string
	Address string
	NumFloors int64
	RateCardOfLot Ratecard
	ParkingLotFloors []ParkingLotFloor
}

type ParkingLotFloor struct {
	FloorNum int64
	Capacity map[consts.VehicleType]int64
	ParkingLotSpots []ParkingSpot
}

type ParkingSpot struct {
	Id string
	IsTaken bool
}

func NewParkingLot(id string, name string, address string, numFloors int64, ratecard Ratecard, lotfloors []ParkingLotFloor) ParkingLot {
	return ParkingLot{
		ID:               id,
		Name:             name,
		Address:          address,
		NumFloors:        numFloors,
		RateCardOfLot:    ratecard,
		ParkingLotFloors: lotfloors,
	}
} 

func NewParkingLotFloor(floorNum int64, capacity map[consts.VehicleType]int64, parkingSpots []ParkingSpot) ParkingLotFloor {
	return ParkingLotFloor{
		FloorNum:        floorNum,
		Capacity:        capacity,
		ParkingLotSpots: parkingSpots,
	}
}

func NewParkingLotSpot(id string) ParkingSpot {
	return ParkingSpot{Id: id}
}


