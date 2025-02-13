package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ride-sharing-service/connection"
	"ride-sharing-service/model"

	"github.com/google/uuid"
)

func AddRider(w http.ResponseWriter,req *http.Request){
	var rider model.Rider
	err:=json.NewDecoder(req.Body).Decode(&rider)
	if err!=nil{
		http.Error(w,"error in decoding body data in rider adding",http.StatusBadRequest)
	}
	rider.RiderID=uuid.New().ClockSequence()
	result:=connection.Db.Create(&rider)
	if result.Error!=nil{
		http.Error(w,"error in creating rider",http.StatusBadRequest)
	}
	response:=map[string]string{
		"message":fmt.Sprintf("sucessfully created rider %d",rider.RiderID),
	}
	json.NewEncoder(w).Encode(response)
	


}