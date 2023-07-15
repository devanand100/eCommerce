package database

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DbConnect () {
		
	Db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if(err != nil){
		log.Fatalf("Database connection error --->%s",err,)
	}

	fmt.Printf("Database connected to ---> %s\n",Db.Migrator().CurrentDatabase())

	DB = Db; 
}