package entities

type DetailProduct struct {
	Product  Product `form:"product" json:"product"`
	Quantity int     `form:"quantity" json:"quantity"`
}

// Getter Setter
func (d *DetailProduct) GetProduct() Product {
	return d.Product
}

func (d *DetailProduct) SetProduct(product Product) {
	d.Product = product
}

func (d *DetailProduct) GetQuantity() int {
	return d.Quantity
}

func (d *DetailProduct) SetQuantity(quantity int) {
	d.Quantity = quantity
}
