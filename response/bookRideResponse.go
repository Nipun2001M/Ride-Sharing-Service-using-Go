package response

type BookRideResponse struct{
	Rideid int `json:"ride_id"`
	Status string `json:"status"`
	Price float64 `json:"price"`

}