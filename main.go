package main

import (
	"fmt"
	"github.com/adrish96/udaan-coding-test/consts"
	"github.com/adrish96/udaan-coding-test/controllers"
	"github.com/adrish96/udaan-coding-test/models"
)

func main() {
	fmt.Println("---Udaan coding test---")

	// create parking lot
	parkinglot := controllers.NewParkingLotController()
	sp1 := parkinglot.NewParkingSpot("sp1")
	sp2 := parkinglot.NewParkingSpot("sp2")

	capacity := map[consts.VehicleType]int64{consts.FourWheeler:10, consts.TwoWheeler:5}
	var spots []models.ParkingSpot
	spots = append(spots, sp1, sp2)
	lotfloor := parkinglot.NewParkingLotFloor(1, capacity, spots)
	lotfloors := []models.ParkingLotFloor{}
	lotfloors = append(lotfloors, lotfloor)

	rate := map[int]float64{1:2.00, 2:1.50}
	defRate := map[consts.VehicleType]float64{consts.FourWheeler:1.00}
	ratecard := models.Ratecard{
		ID:          "",
		Rate: map[consts.VehicleType]map[int]float64{consts.FourWheeler:rate},
		DefaultRate: defRate,
	}
	plot := parkinglot.NewParkingLot("lot1", "lotname", "", 1, ratecard, lotfloors)

	// create vehicle entry
	vehicle := controllers.GetVehicle("vehicle1", consts.FourWheeler, plot.ID)

	// calling ratecard on vehicle exit
	amountToBePaid := controllers.GetRate(vehicle, plot.RateCardOfLot)

	fmt.Printf("final amt := %f", amountToBePaid)

}
