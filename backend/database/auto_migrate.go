package database

import (
	model "github.com/Foodut/backend/models"
	"gorm.io/gorm"
)

func Migrate(con *gorm.DB) {
	// Automigrate model to database
	con.Debug().AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Customer{},
		&model.Seller{},
		&model.DetailProduct{},
		&model.Transaction{},
		&model.Category{},
		&model.Product{},
		&model.Picture{},
	)
}
