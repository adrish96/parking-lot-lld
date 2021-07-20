package controllers

import "github.com/adrish96/udaan-coding-test/logic/parkinglot"

func NewParkingLotController() parkinglot.ParkingLot{
	return parkinglot.ParkingLotImpl{
	}
}



