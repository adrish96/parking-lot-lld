package controllers

import (
	"github.com/adrish96/parking-lot-lld/consts"
	r "github.com/adrish96/parking-lot-lld/logic/ratecard"
	"github.com/adrish96/parking-lot-lld/models"
	"time"
)


// exit

func GetRate(vehicle models.Vehicle, ratecard models.Ratecard)  float64{
	history := vehicle.History[len(vehicle.History)-1]
	entry := history.EntryTime
	exit := time.Now()
	duration := exit.Sub(entry)

	amt := r.GetRate(&ratecard, duration, vehicle.Type)
	return amt
}


// entry
func GetVehicle(id string, vType consts.VehicleType, lotid string) models.Vehicle {
	history := models.VehicleParkingHistory{
		ParkingLotID:   lotid,
		EntryTime:      time.Now(),
	}
	historyList := []*models.VehicleParkingHistory{}
	historyList = append(historyList, &history)

	return models.Vehicle{
		ID:      id,
		Type:    vType,
		History: historyList,
	}
}