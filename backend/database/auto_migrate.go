package database

import (
	prModel "github.com/Foodut/backend/modules/product/domain/model"
	trModel "github.com/Foodut/backend/modules/transaction/domain/model"
	usrModel "github.com/Foodut/backend/modules/user/domain/model"
	"gorm.io/gorm"
)

func Migrate(con *gorm.DB) {

	// Many 2 Many Customer - Product -> Cart
	con.SetupJoinTable(&usrModel.Customer{}, "TempCart", &trModel.Cart{})

	// Many 2 Many Transaction - Product -> Cart
	con.SetupJoinTable(&trModel.Transaction{}, "ProductDetail", &trModel.TransactionDetail{})

	// Automigrate model to database
	con.Debug().AutoMigrate(
		&usrModel.User{},
		&usrModel.Admin{},
		&usrModel.Customer{},
		&usrModel.Seller{},
		// &trModel.DetailTransaksi{},
		// &trModel.Cart{},
		&trModel.Transaction{},
		&prModel.Category{},
		&prModel.Product{},
		&prModel.Picture{},
	)

}
