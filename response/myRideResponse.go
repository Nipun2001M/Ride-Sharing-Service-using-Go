package response

import "database/sql"

type MyRideResponse struct {
	RideId          int     `json:"ride_id"`
	Status          string  `json:"status"`
	Price           float64 `json:"price"`
	RiderID         sql.NullInt64 `json:"rider_id"`
	PickupLocation  string  `json:"pick_up"`
	DropOffLocation string  `json:"drop_off"`
}