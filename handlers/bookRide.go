package handlers

import (
	"encoding/json"
	"net/http"
	"ride-sharing-service/connection"
	"ride-sharing-service/model"
	"ride-sharing-service/response"

	"github.com/google/uuid"
)

func BookRide(w http.ResponseWriter,req *http.Request) {
	var ride model.Ride
	err:=json.NewDecoder(req.Body).Decode(&ride)
	if err!=nil{
		http.Error(w,"error in decoding ride details",http.StatusBadRequest)
	}
	ride.RideID=uuid.New().ClockSequence()
	ride.Status="PENDING"
	ride.Price=ride.Distance*50
	result:=connection.Db.Create(&ride)
	if result.Error!=nil{
		http.Error(w,"error in creation ride",http.StatusBadRequest)
	}
	var res= response.BookRideResponse{
		Rideid: ride.RideID,
		Status: ride.Status,
		Price: ride.Price,

	}
	json.NewEncoder(w).Encode(res)
	



}