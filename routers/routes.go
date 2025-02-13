package routers

import (
	"ride-sharing-service/handlers"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	//add new customer to system
	router.HandleFunc("/users",handlers.AddCustomer).Methods("POST")
	//add new vehical to system
	router.HandleFunc("/riders",handlers.AddRider).Methods("POST")
	//book a ride
	router.HandleFunc("/rides",handlers.BookRide).Methods("POST")
	

	return router;


}