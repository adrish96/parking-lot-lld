package models

import "github.com/adrish96/parking-lot-lld/consts"

type Ratecard struct {
	ID string
	Rate map[consts.VehicleType]map[int]float64
	DefaultRate map[consts.VehicleType]float64
}


func newRatecard(id string, rate map[consts.VehicleType]map[int]float64, def map[consts.VehicleType]float64) Ratecard {
	return Ratecard{
		ID:   id,
		Rate: rate,
		DefaultRate: def,
	}
}

