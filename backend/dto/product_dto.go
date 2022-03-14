package entities

type Product struct {
	ProductId       string   `form:"productId" json:"productId"`
	ProductName     string   `form:"productName" json:"productName"`
	SellerName      string   `form:"sellerName" json:"sellerName"`
	ProductCategory Category `form:"productCategory" json:"productCategory"`
	ProductStock    int      `form:"productStock" json:"productStock"`
	ProductPrice    float64  `form:"productPrice" json:"productPrice"`
	ProductRate     float32  `form:"productRate" json:"productRate"`
	ProductPicture  []string `form:"productPicture" json:"productPicture"`
}

// Getter Setter
func (p *Product) GetProductCategory() Category {
	return p.ProductCategory
}

func (p *Product) SetProductCategory(productCategory Category) {
	p.ProductCategory = productCategory
}

func (p *Product) GetProductId() string {
	return p.ProductId
}

func (p *Product) SetProductId(productId string) {
	p.ProductId = productId
}

func (p *Product) GetProductName() string {
	return p.ProductName
}

func (p *Product) SetProductName(productName string) {
	p.ProductName = productName
}

func (p *Product) GetProductPicture() []string {
	return p.ProductPicture
}

func (p *Product) SetProductPicture(productPicture []string) {
	p.ProductPicture = productPicture
}

func (p *Product) GetProductPrice() float64 {
	return p.ProductPrice
}

func (p *Product) SetProductPrice(productPrice float64) {
	p.ProductPrice = productPrice
}

func (p *Product) GetProductRate() float32 {
	return p.ProductRate
}

func (p *Product) SetProductRate(productRate float32) {
	p.ProductRate = productRate
}

func (p *Product) GetProductStock() int {
	return p.ProductStock
}

func (p *Product) SetProductStock(productStock int) {
	p.ProductStock = productStock
}

func (p *Product) GetSellerName() string {
	return p.SellerName
}

func (p *Product) SetSellerName(sellerName string) {
	p.SellerName = sellerName
}
