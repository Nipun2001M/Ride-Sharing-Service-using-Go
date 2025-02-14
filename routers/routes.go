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
	//riders can watch pending rides that only matches theres vehical type and ride's vehical type
	router.HandleFunc("/rides/pending",handlers.GetPendingRides).Methods("GET")
	//take ride by sending selected ride id and rider id
	router.HandleFunc("/rides/take",handlers.TakeRide).Methods("PUT")
	//user take his active-pending-completed-all ride
	
	//user take ride by sending ride id

	//mark ride as completed by sending payment id
	

	return router;


}