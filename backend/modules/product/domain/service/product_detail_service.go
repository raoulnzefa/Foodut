package services

import (
	prRepo "github.com/Foodut/backend/modules/product/repository"
	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
	trRepo "github.com/Foodut/backend/modules/transaction/repository"
)

func ProductDetailAssociationWithCart(customerId int) []prDto.GetProductDetail {
	// Read to carts table
	productDetails := prRepo.ReadCartProductDetailsAssociation(customerId)

	if len(productDetails) > 0 {
		// for each product, get its association
		for i := 0; i < len(productDetails); i++ {
			productDetails[i].ProductPicture = prRepo.ReadOneProductPicture(productDetails[i].ID)
		}
	}

	return productDetails
}

func ProductDetailAssociationWithTransaction(transactionId int) []prDto.GetProductDetail {
	// Read to transaction_details table
	productDetails := prRepo.ReadTransactionProductDetailsAssociation(transactionId)

	if len(productDetails) > 0 {
		// for each product, get its association
		for i := 0; i < len(productDetails); i++ {
			// Get static price from transaction detail
			productDetails[i].ProductPrice = trRepo.ReadStaticPrice(
				transactionId,
				productDetails[i].ID,
			)

			// Picture Assocaition
			productDetails[i].ProductPicture = prRepo.ReadOneProductPicture(productDetails[i].ID)
		}
	}

	return productDetails
}
