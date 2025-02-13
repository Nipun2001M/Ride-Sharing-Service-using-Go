package handlers

import (
	"encoding/json"
	"net/http"
	"ride-sharing-service/connection"
	"ride-sharing-service/model"
	"ride-sharing-service/response"
)

func GetPendingRides(w http.ResponseWriter, req *http.Request) {
    riderid := req.URL.Query().Get("riderId")

    var rider model.Rider
    if err := connection.Db.First(&rider, riderid).Error; err != nil {
        http.Error(w, "cannot find rider details", http.StatusBadRequest)
        return
    }
    var matchingRides []model.Ride
    if err := connection.Db.Where("status = ? AND vehical_type = ?", "PENDING", rider.VehicleType).Find(&matchingRides).Error; err != nil {
        http.Error(w, "error in getting matching rides", http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
	var matchingridesRes []response.MatchingPendingRidesResponse
	for _,i:=range matchingRides{
		
		matchingridesRes=append(matchingridesRes,response.MatchingPendingRidesResponse{
			RideID:i.RideID,
			Price: i.Price,
			PickUp: i.PickupLocation,
			Drop: i.DropOffLocation,


		} )

	}
    json.NewEncoder(w).Encode(matchingridesRes)

}
