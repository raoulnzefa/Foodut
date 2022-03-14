package database

import (
	picModel "github.com/Foodut/backend/modules/picture/domain/model"
	prModel "github.com/Foodut/backend/modules/product/domain/model"
	trMmodel "github.com/Foodut/backend/modules/transaction/domain/model"
	usrModel "github.com/Foodut/backend/modules/user/domain/model"
	"gorm.io/gorm"
)

func Migrate(con *gorm.DB) {
	// Automigrate model to database
	con.Debug().AutoMigrate(
		&usrModel.User{},
		&usrModel.Admin{},
		&usrModel.Customer{},
		&usrModel.Seller{},
		&prModel.DetailProduct{},
		&trMmodel.Transaction{},
		&prModel.Category{},
		&prModel.Product{},
		&picModel.Picture{},
	)
}
