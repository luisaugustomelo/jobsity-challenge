package migrations

import (
	"log"
	"task-manager-api/db"
	"task-manager-api/models"
)

func Migrate() {
	err := db.DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Error to do migration: %v", err)
	}
	log.Println("Migration was successful!")
}

func DropTable() {
	err := db.DB.Migrator().DropTable(&models.Task{})
	if err != nil {
		log.Fatalf("Error to drop migration: %v", err)
	}
	log.Println("Drop table was successful!")
}
