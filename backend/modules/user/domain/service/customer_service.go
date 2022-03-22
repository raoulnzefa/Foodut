package service

import (
	prSrvc "github.com/Foodut/backend/modules/product/domain/service"
	prRepo "github.com/Foodut/backend/modules/product/repository"
	trSrvc "github.com/Foodut/backend/modules/transaction/domain/service"
	trRepo "github.com/Foodut/backend/modules/transaction/repository"
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"gorm.io/gorm"
)

func SearchCustByIdWithAssociation(custId []string) []model.Customer {
	customers := repo.ReadAllCustomers(custId)

	// Association
	if len(customers) > 0 {
		repo.GetCustomerAssociation(customers)

		// Child association
		for i := 0; i < len(customers); i++ {
			// for each product in cart, get its association
			if len(customers[i].Cart) > 0 {
				prRepo.GetProductsAssociation(customers[i].Cart)
			}
			// for each transaction, get its association
			if len(customers[i].ListHistory) > 0 {
				trRepo.GetTransactionsAssociation(customers[i].ListHistory)

				for j := 0; j < len(customers[i].ListHistory); j++ {
					// for each product in transaction, get its association
					if len(customers[i].ListHistory[j].ProductDetail) > 0 {
						prRepo.GetProductsAssociation(customers[i].ListHistory[j].ProductDetail)
					}
				}
			}

		}
	}

	return customers
}

func SearchCustByIdWithAssociationComplete(custId []string) []dto.GetCustomer {
	customers := repo.ReadAllCustomers(custId)

	var getCustomerDTO []dto.GetCustomer

	// Association
	if len(customers) > 0 {
		repo.GetCustomerAssociation(customers)

		// Child association
		for i := 0; i < len(customers); i++ {

			// Append Mapped Model to Return Object
			getCustomerDTO = append(getCustomerDTO, MapToGetCustomerDTO(customers[i]))

			// for each product in cart, get its association
			if len(customers[i].Cart) > 0 {
				// Get from customer
				cart := prSrvc.ProductDetailAssociationWithCart(customers[i].UserID)

				// Set it to customer collapse
				getCustomerDTO[i].Cart = cart
			}

			// for each transaction, get its association
			if len(customers[i].ListHistory) > 0 {

				for j := 0; j < len(customers[i].ListHistory); j++ {

					// Map from model to dto
					getCustomerDTO[i].ListHistory = append(getCustomerDTO[i].ListHistory, trSrvc.MapToTransactionDTO(customers[i].ListHistory[j]))

					// Get from customer
					historyDetail := prSrvc.ProductDetailAssociationWithTransaction(customers[i].ListHistory[j].ID)

					// Set it to customer complete dto
					getCustomerDTO[i].ListHistory[j].ProductDetail = historyDetail

				}
			}
		}
	}

	return getCustomerDTO
}

func MapToGetCustomerDTO(usr model.Customer) dto.GetCustomer {
	return dto.GetCustomer{
		ID:           usr.UserID,
		Username:     usr.User.Username,
		Email:        usr.User.Email,
		Name:         usr.User.Name,
		Password:     usr.User.Password,
		ProfilePhoto: usr.User.ProfilePhoto,
		Level:        usr.User.Level,
		Address:      usr.Address,
	}
}

func SendForCreateCustomer(usr dto.PostUser) *gorm.DB {
	return repo.CreateCustomer(MapToCustomer(usr))
}

func MapToCustomer(usr dto.PostUser) model.Customer {

	// Parse from JSON DTO -> Database Model
	user := MapToUser(usr, 1)

	cust := model.Customer{
		UserID: user.ID,
		User:   user,
	}

	return cust
}
