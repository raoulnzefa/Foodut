package services

import (
	prRepo "github.com/Foodut/backend/modules/product/repository"
	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
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
			productDetails[i].ProductPicture = prRepo.ReadOneProductPicture(productDetails[i].ID)
		}
	}

	return productDetails
}

func BlendProductDetailForTransaction() {

}

func BlendProductDetailForCart() {

}
