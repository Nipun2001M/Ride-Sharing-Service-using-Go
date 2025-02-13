package response

type MatchingPendingRidesResponse struct{
	RideID		int 	`json:"ride_id"`
	Price		float64	`json:"pickup_location"` 
	PickUp		string	`json:"drop_off_location"`
	Drop		string	`json:"price"`


}