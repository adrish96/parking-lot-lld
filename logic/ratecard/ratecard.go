package ratecard

import (
	"github.com/adrish96/parking-lot-lld/consts"
	"github.com/adrish96/parking-lot-lld/models"
	"math"
	"time"
)

func GetRate(rateCard *models.Ratecard, bookingDuration time.Duration, vehicleType consts.VehicleType) float64 {
	numofHrs := int(math.Ceil(bookingDuration.Hours()))
	ratecard := rateCard.Rate[vehicleType]
	defaultRate := rateCard.DefaultRate[vehicleType]
	finalAmount := 0.00
	var hoursCalc int
	for hour,rate := range ratecard {
		if hour <= numofHrs {
			hoursCalc ++
			finalAmount += rate
		}
	}
	if hoursCalc < numofHrs {
		hrsLeft :=numofHrs-hoursCalc
		finalAmount += float64(hrsLeft) * defaultRate
	}

	return finalAmount
}
