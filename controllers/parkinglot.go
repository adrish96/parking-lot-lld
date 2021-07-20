package controllers

import "github.com/adrish96/parking-lot-lld/logic/parkinglot"

func NewParkingLotController() parkinglot.ParkingLot{
	return parkinglot.ParkingLotImpl{
	}
}



