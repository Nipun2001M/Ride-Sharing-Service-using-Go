package connection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "nipun1234"
    dbname   = "ride_service"
)

func ConnectToDatabase() (*gorm.DB,error){
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",host, user, password, dbname, port)	
	  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	  if err!=nil{
		return nil,err
	  }
	  return db,nil

}