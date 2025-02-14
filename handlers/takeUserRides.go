package handlers

import (
	"encoding/json"
	"net/http"
	"ride-sharing-service/connection"
	"ride-sharing-service/model"
	"ride-sharing-service/response"
	"strconv"
)

func TakeUserRides(w http.ResponseWriter,req *http.Request) {
	userid,_:=strconv.Atoi(req.URL.Query().Get("userid"))
	status:=req.URL.Query().Get("status")
	w.Header().Set("Content-Type","application/json")
	var myrides []model.Ride
	if status!=""{
		if err := connection.Db.Where("user_id = ? AND status = ?", userid,status).Find(&myrides).Error; err != nil {
			http.Error(w, "error in getting matching rides", http.StatusBadRequest)
			return
		}
	}else{
		if err := connection.Db.Where("user_id = ? ",userid).Find(&myrides).Error; err != nil {
			http.Error(w, "error in getting matching rides", http.StatusBadRequest)
			return
		}
	}
	var myRidesRes []response.MyRideResponse
	for _,i:=range myrides{
		myRidesRes=append(myRidesRes, response.MyRideResponse{
			RideId: i.RideID,
			Status: i.Status,
			Price: i.Price,
			PickupLocation: i.PickupLocation,
			DropOffLocation: i.DropOffLocation,
			RiderID: i.RiderID,
		})
	}

	json.NewEncoder(w).Encode(myRidesRes)




}