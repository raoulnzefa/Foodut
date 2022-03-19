package service

import (
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	repo "github.com/Foodut/backend/modules/transaction/repository"
	dto "github.com/Foodut/backend/modules/transaction/rest-api/dto"
	"gorm.io/gorm"
)

func ToOneCart(uid int, pid int, q int) model.Cart {
	// Parse from JSON DTO -> Database Model
	cart := model.Cart{
		CustomerUserID: uid,
		ProductID:      pid,
		Quantity:       q,
	}

	return cart
}

func MapToCart(cr dto.PostCart) *gorm.DB {

	var carts []model.Cart
	// For each products, map it
	for _, p := range cr.Products {
		cart := ToOneCart(cr.UserId, p.ProductId, p.Quantity)
		carts = append(carts, cart)
	}

	return repo.CreateCart(carts)
}
