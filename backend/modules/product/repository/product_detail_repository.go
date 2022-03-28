package repository

import (
	dbController "github.com/Foodut/backend/database"
	dto "github.com/Foodut/backend/modules/product/rest-api/dto"
)

func ReadCartProductDetailsAssociation(customerId int) []dto.GetProductDetail {
	// Check connection
	con := dbController.GetConnection()

	// Get productDetail from database
	var productDetail []dto.GetProductDetail

	// Extra query by id
	if customerId > 0 {
		con.Raw(
			"SELECT products.id, products.product_name, sellers.store_name, categories.product_category, "+
				"products.product_price, products.product_stock, products.product_rate, carts.quantity "+
				"FROM products, carts, sellers, categories "+
				"WHERE products.id = carts.product_id AND products.seller_id = sellers.user_id "+
				"AND products.category_id = categories.id "+
				"AND carts.customer_user_id = ?", customerId).
			Scan(&productDetail)
	}

	return productDetail
}

func ReadTransactionProductDetailsAssociation(transactionId int) []dto.GetProductDetail {
	// Check connection
	con := dbController.GetConnection()

	// Get productDetail from database
	var productDetail []dto.GetProductDetail

	// Extra query by id
	if transactionId > 0 {
		con.Raw(
			"SELECT products.id, products.product_name, sellers.store_name, categories.product_category, "+
				"products.product_price, products.product_stock, products.product_rate, transaction_details.quantity "+
				"FROM products, transaction_details, sellers, categories "+
				"WHERE products.id = transaction_details.product_id "+
				"AND products.seller_id = sellers.user_id "+
				"AND products.category_id = categories.id "+
				"AND transaction_details.transaction_id = ?", transactionId).
			Scan(&productDetail)
	}

	return productDetail
}
