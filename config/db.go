package config

import (
	"belajar/models"
	"belajar/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := utils.Getenv("DB_USERNAME", "")
	password := utils.Getenv("DB_PASS", "")
	database := utils.Getenv("DB_NAME", "")
	host := utils.Getenv("DB_HOST", "")

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Cart{}, &models.CartItem{}, &models.Menu{}, &models.User{})
	return db
}