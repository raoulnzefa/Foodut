package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Foodut/backend/modules/user/domain/model"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	rspn "github.com/Foodut/backend/responses"
)

// TODO
// not done
// Belom Tambahin productnya juga
// belom di pisah ke service, ke repository,
// dan output JSON nya pikirin lagi, DTO nya pake

func RequestSellerByStoreWithProducts(writer http.ResponseWriter, req *http.Request) {

	// Check product_name query
	storeName := req.URL.Query()["store_name"]

	// Get list of seller object
	var seller model.Seller = srvc.CombinedSellerProduct(storeName)

	// Set response
	var response rspn.Response
	if seller.UserID != 0 {
		response.Response_200(seller)
	} else {
		response.Response_204()
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
