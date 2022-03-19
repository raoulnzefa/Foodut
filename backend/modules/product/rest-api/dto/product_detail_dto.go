package dto

type ProductDetail struct {
	Product  Product `form:"product" json:"product"`
	Quantity int     `form:"quantity" json:"quantity"`
}

type PostProductDetail struct {
	ProductId int `form:"productId" json:"productId"`
	Quantity  int `form:"quantity" json:"quantity"`
}

// Getter Setter
func (d *ProductDetail) GetProduct() Product {
	return d.Product
}

func (d *ProductDetail) SetProduct(product Product) {
	d.Product = product
}

func (d *ProductDetail) GetQuantity() int {
	return d.Quantity
}

func (d *ProductDetail) SetQuantity(quantity int) {
	d.Quantity = quantity
}
