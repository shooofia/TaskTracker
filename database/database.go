// database.go

package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "username:password@tcp(localhost:3306)/tasktracker?charset=utf8mb4&parseTime=True&loc=Local"
	// Ganti "username" dan "password" dengan informasi login ke database MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto migrate tabel Task
	err = db.AutoMigrate(&Task{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	return db, nil
}
