package entities

type DetailProduct struct {
	product  Product
	quantity int
}

// Getter Setter
func (d *DetailProduct) GetProduct() Product {
	return d.product
}

func (d *DetailProduct) SetProduct(product Product) {
	d.product = product
}

func (d *DetailProduct) GetQuantity() int {
	return d.quantity
}

func (d *DetailProduct) SetQuantity(quantity int) {
	d.quantity = quantity
}
