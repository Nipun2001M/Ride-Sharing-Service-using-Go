package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ride-sharing-service/connection"
	"ride-sharing-service/model"
	"strconv"
)

func TakeRide(w http.ResponseWriter,req *http.Request){
	rideid,_:=strconv.Atoi(req.URL.Query().Get("rideid"))
	riderid,_:=strconv.Atoi(req.URL.Query().Get("riderid"))
	
	result:=connection.Db.Model(&model.Ride{}).Where("ride_id=?",rideid).Updates(map[string]interface{}{
		"rider_id":riderid,
		"status":"ACTIVE",
	})

	if result.Error!=nil{
		http.Error(w,"error in updating rider details and status",http.StatusBadRequest)
	}
	res:=map[string]string{
		"message":fmt.Sprintf("You are sucessfully taken the ride , ride id :%d",rideid),
	}

	json.NewEncoder(w).Encode(res)
}