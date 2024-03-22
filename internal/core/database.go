package core

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URI")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
