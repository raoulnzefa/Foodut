package services

import (
	"strconv"

	model "github.com/Foodut/backend/modules/product/domain/model"
	repo "github.com/Foodut/backend/modules/product/repository"
	dto "github.com/Foodut/backend/modules/product/rest-api/dto"
	trModel "github.com/Foodut/backend/modules/transaction/domain/model"

	"gorm.io/gorm"
)

func SearchById(productId []string) []model.Product {
	products := repo.ReadAllProducts(productId)

	// Association
	if len(products) > 0 {
		repo.GetProductsAssociation(products)
	}

	return products
}

func SearchByName(name []string) []model.Product {
	products := repo.ReadProductsByName(name)

	// If none, try find using 'LIKE'
	if len(products) < 1 {
		products = repo.ReadProductsByNameAlike(name)
	}

	// Association
	if len(products) > 0 {
		repo.GetProductsAssociation(products)
	}

	return products
}

func SearchByCategoryId(id []string) []model.Product {
	// Convert categoryId to int
	categoryId, err := strconv.Atoi(id[0])

	var products []model.Product
	if err == nil {
		products = repo.ReadProductsByCategoryId(categoryId)
	}

	// Association
	if len(products) > 0 {
		repo.GetProductsAssociation(products)
	}

	return products
}

func SearchByCategoryName(categoryName []string) []model.Product {
	products := repo.ReadProductsByCategoryName(categoryName)

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

func SendForUpdateProductStockAfterTransaction(carts []trModel.Cart, products []model.Product) *gorm.DB {
	var updateRes *gorm.DB

	for i := 0; i < len(carts); i++ {
		updateRes = repo.UpdateProductStock(
			products[i].ID,
			(products[i].ProductStock - carts[i].Quantity),
		)
	}

	return updateRes
}

func SendForCreateProduct(pr dto.PostProduct) *gorm.DB {
	return repo.CreateProduct(MapToProduct(pr))
}

func MapToProduct(pr dto.PostProduct) model.Product {

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
		Description:  pr.ProductDescription,
		SellerID:     pr.SellerId,
		CategoryID:   pr.ProductCategory,
		Picture:      pics,
	}

	return product
}

func EditProduct(product model.Product, editProductDto dto.EditProduct) *gorm.DB {
	productName := editProductDto.ProductName
	productPrice := editProductDto.ProductPrice
	productStock := editProductDto.ProductStock
	productCategory := editProductDto.ProductCategory

	if productName != "" {
		product.ProductName = productName
	}
	if productPrice > -1 {
		product.ProductPrice = productPrice
	}
	if productStock > -1 {
		product.ProductStock = productStock
	}
	if productCategory > 0 {
		product.CategoryID = productCategory
	}

	if product.ID > 0 {
		result := repo.SaveProduct(product)
		return result
	} else {
		return nil
	}
}
