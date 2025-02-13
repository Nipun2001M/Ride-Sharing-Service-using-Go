package main

import (
	"fmt"
	"log"
	"net/http"
	"ride-sharing-service/connection"

	"github.com/gorilla/mux"
)

func main(){

   

    _,error:=connection.ConnectToDatabase()
    if error!=nil{
        fmt.Println("Error Connection to Databsae...")
    }else {
        fmt.Println("Sucessfully Connected to Database...")
    }

    router:=mux.NewRouter()
    fmt.Println("server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080",router))

   




}