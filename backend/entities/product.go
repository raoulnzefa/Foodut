package entities

type Product struct {
	productCategory Category
	productId       string
	productName     string
	productPicture  []string
	productPrice    float64
	productRate     float32
	productStock    int
	sellerName      string
}

// Getter Setter
func (p *Product) GetProductCategory() Category {
	return p.productCategory
}

func (p *Product) SetProductCategory(productCategory Category) {
	p.productCategory = productCategory
}

func (p *Product) GetProductId() string {
	return p.productId
}

func (p *Product) SetProductId(productId string) {
	p.productId = productId
}

func (p *Product) GetProductName() string {
	return p.productName
}

func (p *Product) SetProductName(productName string) {
	p.productName = productName
}

func (p *Product) GetProductPicture() []string {
	return p.productPicture
}

func (p *Product) SetProductPicture(productPicture []string) {
	p.productPicture = productPicture
}

func (p *Product) GetProductPrice() float64 {
	return p.productPrice
}

func (p *Product) SetProductPrice(productPrice float64) {
	p.productPrice = productPrice
}

func (p *Product) GetProductRate() float32 {
	return p.productRate
}

func (p *Product) SetProductRate(productRate float32) {
	p.productRate = productRate
}

func (p *Product) GetProductStock() int {
	return p.productStock
}

func (p *Product) SetProductStock(productStock int) {
	p.productStock = productStock
}

func (p *Product) GetSellerName() string {
	return p.sellerName
}

func (p *Product) SetSellerName(sellerName string) {
	p.sellerName = sellerName
}
