package entities

import (
	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
)

type Seller struct {
	User        User            `form:"user" json:"user"`
	StoreName   string          `form:"storeName" json:"storeName"`
	City        string          `form:"city" json:"city"`
	ListProduct []prDto.Product `form:"listProduct" json:"listProduct"`
}

type SellerMinimal struct {
	UserID    int    `form:"id" json:"id"`
	StoreName string `form:"storeName" json:"storeName"`
	City      string `form:"city" json:"city"`
}

type SellerMinimalProducts struct {
	UserID      int             `form:"id" json:"id"`
	StoreName   string          `form:"storeName" json:"storeName"`
	City        string          `form:"city" json:"city"`
	ListProduct []prDto.Product `form:"listProduct" json:"listProduct"`
}

type PostSeller struct {
	Name      string `form:"name" json:"name"`
	Username  string `form:"username" json:"username"`
	Email     string `form:"email" json:"email"`
	Password  string `form:"password" json:"password"`
	StoreName string `form:"storeName" json:"storeName"`
	City      string `form:"city" json:"city"`
}

//-------------------------------------------------
// Getter Setter Full
func (s *Seller) GetUser() User {
	return s.User
}

func (s *Seller) SetUser(user User) {
	s.User = user
}

func (s *Seller) GetStoreName() string {
	return s.StoreName
}

func (s *Seller) SetStoreName(storeName string) {
	s.StoreName = storeName
}

func (s *Seller) GetCity() string {
	return s.City
}

func (s *Seller) SetCity(city string) {
	s.City = city
}

func (s *Seller) GetListProduct() []prDto.Product {
	return s.ListProduct
}

func (s *Seller) SetListProduct(listProduct []prDto.Product) {
	s.ListProduct = listProduct
}

//-------------------------------------------------
// Getter Setter Minimal
func (s *SellerMinimal) GetUserId() int {
	return s.UserID
}

func (s *SellerMinimal) SetUserId(id int) {
	s.UserID = id
}

func (s *SellerMinimal) GetStoreName() string {
	return s.StoreName
}

func (s *SellerMinimal) SetStoreName(storeName string) {
	s.StoreName = storeName
}

func (s *SellerMinimal) GetCity() string {
	return s.City
}

func (s *SellerMinimal) SetCity(city string) {
	s.City = city
}

//-------------------------------------------------
// Getter Setter Minimal Products
func (s *SellerMinimalProducts) GetUserId() int {
	return s.UserID
}

func (s *SellerMinimalProducts) SetUserId(id int) {
	s.UserID = id
}

func (s *SellerMinimalProducts) GetStoreName() string {
	return s.StoreName
}

func (s *SellerMinimalProducts) SetStoreName(storeName string) {
	s.StoreName = storeName
}

func (s *SellerMinimalProducts) GetCity() string {
	return s.City
}

func (s *SellerMinimalProducts) SetCity(city string) {
	s.City = city
}

func (s *SellerMinimalProducts) GetListProduct() []prDto.Product {
	return s.ListProduct
}

func (s *SellerMinimalProducts) SetListProduct(listProduct []prDto.Product) {
	s.ListProduct = listProduct
}
