package connection

import (
	"fmt"
	"log"
	"ride-sharing-service/model"
)

func RunMigrations() {

	var err = Db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Error migrating User table: %v", err)
	}

	err = Db.AutoMigrate(&model.Rider{})
	if err != nil {
		log.Fatalf("Error migrating Rider table: %v", err)
	}

	err = Db.AutoMigrate(&model.Ride{})
	if err != nil {
		log.Fatalf("Error migrating Ride table: %v", err)
	}

	fmt.Println("Tables migrated successfully...")

}