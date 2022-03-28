package dto

import "github.com/Foodut/backend/modules/product/domain/model"

type ProductDetail struct {
	Product  model.Product `form:"product" json:"product"`
	Quantity int           `form:"quantity" json:"quantity"`
}

type GetProductDetail struct {
	ID              int      `form:"id" json:"id"`
	ProductName     string   `form:"productName" json:"productName"`
	StoreName       string   `form:"storeName" json:"storeName"`
	ProductCategory string   `form:"productCategory" json:"productCategory"`
	ProductStock    int      `form:"productStock" json:"productStock"`
	ProductPrice    float64  `form:"productPrice" json:"productPrice"`
	ProductRate     float32  `form:"productRate" json:"productRate"`
	Quantity        int      `form:"quantity" json:"quantity"`
	ProductPicture  []string `form:"productPicture" json:"productPicture" gorm:"type:text"`
}

type PostProductDetail struct {
	ProductId int `form:"productId" json:"productId"`
	Quantity  int `form:"quantity" json:"quantity"`
}

// Getter Setter
func (d *ProductDetail) GetProduct() model.Product {
	return d.Product
}

func (d *ProductDetail) SetProduct(product model.Product) {
	d.Product = product
}

func (d *ProductDetail) GetQuantity() int {
	return d.Quantity
}

func (d *ProductDetail) SetQuantity(quantity int) {
	d.Quantity = quantity
}
