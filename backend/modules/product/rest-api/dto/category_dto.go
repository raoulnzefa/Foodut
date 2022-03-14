package entities

type Category struct {
	ProductCategory string `form:"productCategory" json:"productCategory"`
}

// Getter Setter
func (c *Category) GetProductCategory() string {
	return c.ProductCategory
}

func (c *Category) SetProductCategory(productCategory string) {
	c.ProductCategory = productCategory
}
