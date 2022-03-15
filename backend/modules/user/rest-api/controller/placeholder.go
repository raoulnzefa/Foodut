package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbController "github.com/Foodut/backend/database"
	prModel "github.com/Foodut/backend/modules/product/domain/model"
	usrModel "github.com/Foodut/backend/modules/user/domain/model"
	usrDto "github.com/Foodut/backend/modules/user/rest-api/dto"
	response "github.com/Foodut/backend/responses"
)

func TestDB(writer http.ResponseWriter, req *http.Request) {
	// AddCategory()
	// AddUser()
	// AddCustomer()
	// AddSeller()
	// AddProduct()
	GetSeller()
}

func AddCategory() {
	db := dbController.GetConnection()

	cate := prModel.Category{
		ProductCategory: "Basah",
	}

	result := db.Create(&cate)

	fmt.Println("Category ID just inserted {}", cate.ID)

	fmt.Println(result.Error)
}

func AddUser() {
	db := dbController.GetConnection()

	user := usrModel.User{
		// ID:           3,
		Username:     "Jeddi123",
		Email:        "jeddi@gmail.com",
		Name:         "Jedediah Fanuel",
		Password:     "43214321",
		ProfilePhoto: "database/photo/profile/jf.jpg",
		Level:        1, // Seller
	}

	result := db.Create(&user)

	fmt.Println(user.ID)
	fmt.Println(user.Username)

	fmt.Println(result.Error)

}

func AddCustomer() {
	db := dbController.GetConnection()

	cust := usrModel.Customer{
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

	sell := usrModel.Seller{
		ID:        3,
		City:      "Jayapura",
		StoreName: "Papeda Legit",
	}

	result := db.Create(&sell)

	fmt.Println(sell.ID)
	fmt.Println(sell.StoreName)

	fmt.Println(result.Error)

}

func AddProduct() {
	db := dbController.GetConnection()

	pic := prModel.Picture{
		PicturePath: "database/picture/product/pad3.jpg",
		ProductID:   5,
	}

	pic2 := prModel.Picture{
		PicturePath: "database/picture/product/pad1.jpg",
		ProductID:   5,
	}

	pic3 := prModel.Picture{
		PicturePath: "database/picture/product/pad2.jpg",
		ProductID:   5,
	}

	var picArr []prModel.Picture
	picArr = append(picArr, pic)
	picArr = append(picArr, pic2)
	picArr = append(picArr, pic3)

	prod := prModel.Product{
		ProductName:  "Papeda Cake January",
		ProductPrice: 25000,
		ProductRate:  4.9,
		ProductStock: 90,
		SellerID:     3,
		CategoryID:   1,
		Picture:      picArr,
	}

	result := db.Create(&prod)

	fmt.Println(result.Error)
}

func GetSeller() {
	db := dbController.GetConnection()

	var user usrModel.User

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

func GetUserResponse(writer http.ResponseWriter, req *http.Request) {

	db := dbController.GetConnection()

	var userRepo usrModel.User

	db.Limit(1).Find(&userRepo)

	var user usrDto.User
	var users []usrDto.User

	user.SetEmail(userRepo.Email)
	user.SetName(userRepo.Name)
	user.SetPassword(userRepo.Password)
	user.SetProfilePhoto(userRepo.ProfilePhoto)
	user.SetUsername(userRepo.Username)
	user.SetLevel(userRepo.Level)

	users = append(users, user)

	var userResponse response.Response
	userResponse.Response_200(users)

	fmt.Println(userResponse)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(userResponse)
}
