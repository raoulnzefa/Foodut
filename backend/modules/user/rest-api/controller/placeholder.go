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
	// AddAdmin()
	// AddCustomer()
	// AddSeller()
	// AddProduct()
	GetAdmin()
	GetSeller()
	GetCustomer()
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

	// user := usrModel.User{
	// 	// ID:           3,
	// 	Username:     "Jeddi123",
	// 	Email:        "jeddi@gmail.com",
	// 	Name:         "Jedediah Fanuel",
	// 	Password:     "43214321",
	// 	ProfilePhoto: "database/photo/profile/jf.jpg",
	// 	Level:        2, // Admin
	// }

	// user := usrModel.User{
	// 	// ID:           3,
	// 	Username:     "Fedly123",
	// 	Email:        "fedly@gmail.com",
	// 	Name:         "Fedly Septian",
	// 	Password:     "12345678",
	// 	ProfilePhoto: "database/photo/profile/fs.jpg",
	// 	Level:        1, // Seller
	// }

	user := usrModel.User{
		// ID:           3,
		Username:     "Timothy123",
		Email:        "timothy@gmail.com",
		Name:         "Timothy Ray",
		Password:     "12344321",
		ProfilePhoto: "database/photo/profile/tr.jpg",
		Level:        0, // Customer
	}

	result := db.Create(&user)

	fmt.Println(user.ID)
	fmt.Println(user.Username)

	fmt.Println(result.Error)

}

func AddAdmin() {
	db := dbController.GetConnection()

	user := usrModel.Admin{
		UserID: 1,
	}

	result := db.Create(&user)

	fmt.Println(user.UserID)

	fmt.Println(result.Error)

}

func AddCustomer() {
	db := dbController.GetConnection()

	cust := usrModel.Customer{
		UserID:  3,
		Address: "Jl. Kopo Permai No. 42",
	}

	result := db.Create(&cust)

	fmt.Println(cust.Address)

	fmt.Println(result.Error)

}

func AddSeller() {
	db := dbController.GetConnection()

	sell := usrModel.Seller{
		UserID:    2,
		City:      "Bandung",
		StoreName: "Makaroni Asoy",
	}

	result := db.Create(&sell)

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

func GetAdmin() {
	db := dbController.GetConnection()

	var adm usrModel.Admin

	db.Limit(1).Find(&adm)

	db.Model(&adm).Association("User")
	db.Model(&adm).Association("User").Find(&adm.User)

	fmt.Printf("%+v \n", adm)
}

func GetSeller() {
	db := dbController.GetConnection()

	var seller usrModel.Seller

	db.Limit(1).Find(&seller)

	db.Model(&seller).Association("User")
	db.Model(&seller).Association("User").Find(&seller.User)

	fmt.Printf("%+v \n", seller)
}

func GetCustomer() {
	db := dbController.GetConnection()

	var cust usrModel.Customer

	db.Limit(1).Find(&cust)

	db.Model(&cust).Association("User")
	db.Model(&cust).Association("User").Find(&cust.User)

	fmt.Printf("%+v \n", cust)
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
