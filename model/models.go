package model

import "database/sql"

type User struct {
	UserID    int    `json:"user_id" gorm:"primaryKey"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
	Rides     []Ride `json:"rides" gorm:"foreignKey:UserID"`
}

type Ride struct {
	RideID          int     `json:"ride_id" gorm:"primaryKey"`
	UserID          int     `json:"user_id" gorm:"index"`
	RiderID         sql.NullInt64      `json:"rider_id" gorm:"index"`
	PaymentID       sql.NullString  `json:"payment_id"`
	PickupLocation  string  `json:"pickup_location"`
	DropOffLocation string  `json:"drop_off_location"`
	Distance        float64 `json:"distance"`
	Price           float64 `json:"price"`
	VehicalType     string  `json:"vehical_type"`
	Status          string  `json:"status"`
}

type Rider struct {
	RiderID     int    `json:"rider_id" gorm:"primaryKey"`
	Name        string `json:"name"`
	ContactNo   string `json:"contact_no"`
	VehicleType string `json:"vehicle_type"`
	Rides       []Ride `json:"rides" gorm:"foreignKey:RiderID"`
}
