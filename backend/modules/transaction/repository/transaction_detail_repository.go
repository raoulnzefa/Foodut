package repository

import (
	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	"github.com/Foodut/backend/modules/transaction/rest-api/dto"
	"gorm.io/gorm"
)

func ReadStaticPrice(tid int, pid int) float64 {
	// Check connection
	con := dbController.GetConnection()

	var staticPrice float64

	// Update object in database
	con.Raw(
		"SELECT `price` "+
			"FROM `transaction_details` "+
			"WHERE transaction_id = ? "+
			"AND product_id = ?",
		tid,
		pid,
	).Scan(&staticPrice)

	return staticPrice
}

func ReadAllTransactionDetail() []dto.TransactionDecentralize {
	// Check connection
	con := dbController.GetConnection()

	var TransactionDecentralize []dto.TransactionDecentralize

	con.Raw(
		"SELECT pictures.picture_path, products.product_name, transaction_details.price, " +
			"transaction_details.quantity, transactions.address, transactions.payment_option, " +
			"transactions.transaction_date " +
			"FROM `transactions`, `transaction_details` , `products`, `pictures` " +
			"WHERE transactions.id = transaction_details.transaction_id " +
			"AND transaction_details.product_id = products.id " +
			"AND transaction_details.product_id = pictures.product_id " +
			"GROUP BY transactions.transaction_date, products.product_name " +
			"ORDER BY transactions.transaction_date DESC").
		Scan(&TransactionDecentralize)

	return TransactionDecentralize
}

func ReadTDOrderBySeller(sellerId []string) []dto.OrderDetail {
	// Check connection
	con := dbController.GetConnection()

	var OrderDetail []dto.OrderDetail

	if sellerId != nil {
		con.Raw(
			"SELECT transaction_details.*, products.product_name, transactions.customer_id, "+
				"transactions.payment_option, transactions.transaction_date, "+
				"users.name, transactions.address "+
				"FROM `transactions`, `transaction_details` , `products`, `users` "+
				"WHERE users.id = transactions.customer_id "+
				"AND transactions.id = transaction_details.transaction_id "+
				"AND transaction_details.product_id = products.id "+
				"AND status = 'ORDER' "+
				"AND transaction_details.seller_id = ?",
			sellerId[0]).
			Scan(&OrderDetail)
	} else {
		con.Raw(
			"SELECT transaction_details.*, products.product_name, transactions.customer_id, " +
				"transactions.payment_option, transactions.transaction_date, " +
				"users.name, transactions.address " +
				"FROM `transactions`, `transaction_details` , `products`, `users` " +
				"WHERE users.id = transactions.customer_id " +
				"AND transactions.id = transaction_details.transaction_id " +
				"AND transaction_details.product_id = products.id " +
				"AND status = 'ORDER' ").
			Scan(&OrderDetail)
	}

	return OrderDetail
}

func UpdateTDAfterTransaction(transactionDetails []model.TransactionDetail) *gorm.DB {
	// Check connection
	con := dbController.GetConnection()

	// Update object in database
	result := con.Save(&transactionDetails)

	return result
}
