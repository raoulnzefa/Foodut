package controllers

import (
	"fmt"

	dbController "github.com/Foodut/backend/controllers/database"
	model "github.com/Foodut/backend/models"
)

func TestDB() {
	// AddCategory()
	// AddUser()
	// AddCustomer()
	// AddSeller()
	// AddProduct()
}

func AddCategory() {
	db := dbController.GetConnection()

	cate := model.Category{
		ProductCategory: "Basah",
	}

	result := db.Create(&cate)

	fmt.Println("Category ID just inserted {}", cate.CategoryID)

	fmt.Println(result.Error)
}

func AddUser() {
	db := dbController.GetConnection()

	user := model.User{
		ID:           2,
		Username:     "Timothy123",
		Email:        "timothy@gmail.com",
		Name:         "Timothy Ray",
		Password:     "12341234",
		ProfilePhoto: "database/photo/profile/tr.jpg",
		Level:        0, // Seller
	}

	result := db.Create(&user)

	fmt.Println(user.ID)
	fmt.Println(user.Username)

	fmt.Println(result.Error)

}

func AddCustomer() {
	db := dbController.GetConnection()

	cust := model.Customer{
		ID:      2,
		Address: "Jl. Kopo Permai No. 42",
	}

	result := db.Create(&cust)

	fmt.Println(cust.ID)
	fmt.Println(cust.Address)

	fmt.Println(result.Error)

}

func AddSeller() {
	db := dbController.GetConnection()

	sell := model.Seller{
		City:      "Bandung",
		StoreName: "Makaroni Asoy",
	}

	result := db.Create(&sell)

	fmt.Println(sell.ID)
	fmt.Println(sell.StoreName)

	fmt.Println(result.Error)

}

func AddProduct() {
	db := dbController.GetConnection()

	pic := model.Picture{
		PicturePath: "database/picture/product/mrb1.jpg",
		ProductID:   2,
	}

	pic2 := model.Picture{
		PicturePath: "database/picture/product/mrb2.jpg",
		ProductID:   2,
	}

	var picArr []model.Picture
	picArr = append(picArr, pic)
	picArr = append(picArr, pic2)

	prod := model.Product{
		ProductName:  "Makaroni Royco Basah",
		ProductPrice: 5000,
		ProductRate:  4.8,
		ProductStock: 100,
		SellerID:     1,
		CategoryID:   2,
		Picture:      picArr,
	}

	result := db.Create(&prod)

	fmt.Println(prod.CategoryID)

	fmt.Println(result.Error)
}

func GetSeller() {
	db := dbController.GetConnection()

	var user model.User

	db.Limit(1).Find(&user)

	db.Model(&user).Association("Seller")
	db.Model(&user).Association("Seller").Find(&user.Seller)

	db.Model(&user.Seller).Association("ListProduct")
	db.Model(&user.Seller).Association("ListProduct").Find(&user.Seller.ListProduct)

	db.Model(&user.Seller.ListProduct).Association("Picture")

	for i := 0; i < len(user.Seller.ListProduct); i++ {
		db.Model(&user.Seller.ListProduct[i]).Association("Picture").Find(&user.Seller.ListProduct[i].Picture)
	}

	fmt.Println(user)
}
