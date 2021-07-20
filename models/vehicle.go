package models

import (
	"github.com/adrish96/udaan-coding-test/consts"
	"time"
)

type Vehicle struct {
	ID string
	Type consts.VehicleType
	History []*VehicleParkingHistory
}

type VehicleParkingHistory struct {
	ParkingLotID string
	ParkingLotName string
	Duration time.Duration
	AmtPaid float64
	Date string
	EntryTime time.Time
	ExitTime time.Time
}


func newVehicle(id string, vehicleType consts.VehicleType, history []*VehicleParkingHistory) Vehicle {
	return Vehicle{
		ID:      id,
		Type:    vehicleType,
		History: history,
	}
}