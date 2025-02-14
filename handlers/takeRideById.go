package handlers

import (
	"encoding/json"
	"net/http"
	"ride-sharing-service/connection"
	"ride-sharing-service/model"
	"ride-sharing-service/response"
	"strconv"

	"github.com/gorilla/mux"
)

func GetRidebyId(w http.ResponseWriter,req *http.Request) {
	var params =mux.Vars(req)
	rideid,_:=strconv.Atoi(	params["rideid"])
	w.Header().Set("Content-Type","application/json")
	var ride model.Ride
	result:=connection.Db.First(&ride,rideid)
	if result.Error!=nil{
		http.Error(w,"error in getting ride",http.StatusBadRequest)
		return
	}
	var res=response.MyRideResponse{
		RideId: ride.RideID,
		Status: ride.Status,
		Price: ride.Price,
		RiderID: ride.RiderID,
		PickupLocation: ride.PickupLocation,
		DropOffLocation: ride.DropOffLocation,
	}
	json.NewEncoder(w).Encode(res)
	

}