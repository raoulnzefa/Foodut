package services

import (
	model "github.com/Foodut/backend/modules/product/domain/model"
	repo "github.com/Foodut/backend/modules/product/repository"
	dto "github.com/Foodut/backend/modules/product/rest-api/dto"

	"gorm.io/gorm"
)

func SearchById(productId []string) []model.Product {
	products := repo.FindAllProducts(productId)

	// Association
	if len(products) > 0 {
		repo.GetProductsAssociation(products)
	}

	return products
}

func SearchByName(name []string) []model.Product {
	products := repo.FindProductsByName(name)

	// If none, try find using 'LIKE'
	if len(products) < 1 {
		products = repo.FindProductsByNameAlike(name)
	}

	// Association
	if len(products) > 0 {
		repo.GetProductsAssociation(products)
	}

	return products
}


func DeleteById(productId string) *gorm.DB {
	deleteFeedback := repo.DeleteProductById(productId)

	return deleteFeedback
}

func MapToProduct(pr dto.PostProduct) *gorm.DB {

	// Convert path string array -> Model Picture
	var pics []model.Picture
	for _, path := range pr.ProductPicture {
		pics = append(pics, model.Picture{
			PicturePath: path,
		})
	}

	// Parse from JSON DTO -> Database Model
	product := model.Product{
		ProductName:  pr.ProductName,
		ProductPrice: pr.ProductPrice,
		ProductStock: pr.ProductStock,
		SellerID:     pr.SellerId,
		CategoryID:   pr.ProductCategory,
		Picture:      pics,
	}

	return repo.CreateProduct(product)
}
