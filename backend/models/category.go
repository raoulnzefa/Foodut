package models

type Category struct {
	productCategory string
}

// Getter Setter
func (c *Category) GetProductCategory() string {
	return c.productCategory
}

func (c *Category) SetProductCategory(productCategory string) {
	c.productCategory = productCategory
}
