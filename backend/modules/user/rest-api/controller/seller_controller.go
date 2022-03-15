package controller

import (
	"encoding/json"
	"net/http"

	dbController "github.com/Foodut/backend/database"
	usrModel "github.com/Foodut/backend/modules/user/domain/model"
	rspn "github.com/Foodut/backend/responses"
)

// TODO
// not done
// Belom Tambahin productnya juga
// belom di pisah ke service, ke repository,
// dan output JSON nya pikirin lagi, DTO nya pake
func something() usrModel.Seller {
	db := dbController.GetConnection()

	var seller usrModel.Seller

	db.Limit(1).Find(&seller)
	db.Model(&seller).Association("User")
	db.Model(&seller).Association("User").Find(&seller.User)

	return seller
}

func GetSellerWithProducts(writer http.ResponseWriter, req *http.Request) {

	// Get product and check by query
	var seller []usrModel.Seller

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
