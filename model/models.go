package model

type User struct {
    UserID    string `json:"user_id" gorm:"primaryKey"`
    Name      string `json:"name"`
    ContactNo string `json:"contact_no"`
    Address   string `json:"address"`
    Rides     []Ride `json:"rides" gorm:"foreignKey:UserId"`
}

type Ride struct {
    RideID         string  `json:"ride_id" gorm:"primaryKey"`
    UserID         string  `json:"user_id"`
    RiderID        string  `json:"rider_id"`
    PaymentID      string  `json:"payment_id"`
    PickupLocation string  `json:"pickup_location"`
    DropOffLocation string `json:"drop_off_location"`
    Price         float64  `json:"price"`
}

type Rider struct {
    RiderID     string `json:"rider_id" gorm:"primaryKey"`
    Name        string `json:"name"`
    ContactNo   string `json:"contact_no"`
    VehicleType string `json:"vehicle_type"`
	Rides 		[]Ride `json:"rides" gorm:"foreignKey:RiderID"`
}
