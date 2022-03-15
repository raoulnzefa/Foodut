package controller

import (
	"encoding/json"
	"net/http"

	dbController "github.com/Foodut/backend/database"
	usrModel "github.com/Foodut/backend/modules/user/domain/model"
	rspn "github.com/Foodut/backend/responses"
)

func something() usrModel.User {
	db := dbController.GetConnection()

	var user usrModel.User

	db.Limit(1).Find(&user)

	db.Model(&user).Association("Seller")
	db.Model(&user).Association("Seller").Find(&user.Seller)

	db.Model(&user.Seller).Association("ListProduct")
	db.Model(&user.Seller).Association("ListProduct").Find(&user.Seller.ListProduct)

	db.Model(&user.Seller.ListProduct).Association("Picture")

	for i := 0; i < len(user.Seller.ListProduct); i++ {
		db.Model(&user.Seller.ListProduct[i]).Association("Picture").Find(&user.Seller.ListProduct[i].Picture)
	}

	return user
}

func GetSellerWithProducts(writer http.ResponseWriter, req *http.Request) {

	// Get product and check by query
	var seller []usrModel.User

	seller = append(seller, something())

	// Set response
	var response rspn.Response
	if len(seller) > 0 {
		response.Response_200(seller)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
