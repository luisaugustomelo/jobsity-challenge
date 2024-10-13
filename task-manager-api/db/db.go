// /db/db.go
package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "taskuser:taskpassword@tcp(localhost:3306)/taskdb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	fmt.Println("Database connection established")
}

func CloseConnection() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatalf("Close database connection error: %v", err)
		}
	}(sqlDB)
}
