package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ride-sharing-service/connection"
	"ride-sharing-service/model"

	"github.com/google/uuid"
)

func AddCustomer(w http.ResponseWriter,req *http.Request) {
	var user model.User
	w.Header().Set("Content-Type","application/json")
	err:=json.NewDecoder(req.Body).Decode(&user)
	if err!=nil{
		fmt.Println("error ocured in decoding",err)
	}
	user.UserID=uuid.New().ClockSequence()
	result:=connection.Db.Create(&user)
	if result.Error!=nil{
		http.Error(w,"error in creating book",http.StatusBadRequest)
	}
	 response:=map[string]string{
		"message":	 fmt.Sprintf("sucessfully created user %d",user.UserID),

	 }
	 json.NewEncoder(w).Encode(response)



	

}