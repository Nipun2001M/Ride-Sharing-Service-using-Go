package request

type Pay struct{
	PaymentID string `json:"payment_id"`
	RideID int `json:"ride_id"`
}