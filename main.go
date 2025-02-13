package main

import (
	"fmt"
	"log"
	"net/http"
	"ride-sharing-service/connection"
	"ride-sharing-service/routers"
)

func main(){

   

    _,error:=connection.ConnectToDatabase()
    if error!=nil{
        fmt.Println("Error Connection to Databsae...")
    }else {
        fmt.Println("Sucessfully Connected to Database...")
    }

	// err:= db.AutoMigrate(&model.User{}, &model.Ride{}, &model.Rider{})
    // if err!=nil{
    //     fmt.Println("error in migration tables ",err)
    // }
    connection.RunMigrations()
    router:=routers.GetRouter()
    fmt.Println("server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080",router))

   




}