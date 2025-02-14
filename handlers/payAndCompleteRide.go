package handlers

import (
	"encoding/json"
	"net/http"
	"ride-sharing-service/connection"
	"ride-sharing-service/model"
	"ride-sharing-service/request"
)

func PayAndCompleteRide(w http.ResponseWriter,req * http.Request){
	w.Header().Set("Content-Type","application/json")
	var reqBody request.Pay
	err:=json.NewDecoder(req.Body).Decode(&reqBody)
	if err!=nil{
		http.Error(w,"error in decoding req body",http.StatusBadRequest)
		return
	}
	result:=connection.Db.Model(&model.Ride{}).Where("ride_id=?",reqBody.RideID).Updates(map[string]interface{}{
		"status":"COMPLETED",
		"payment_id":reqBody.PaymentID,
	})
	if result.Error!=nil{
		http.Error(w,"error in updating payment",http.StatusBadRequest)
		return
	}
	res:=map[string]string{
		"message":"sucessfully updated payment and status",
	}
	json.NewEncoder(w).Encode(res)




}