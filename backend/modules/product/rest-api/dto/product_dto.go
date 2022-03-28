package dto

type Product struct {
	ProductId       int      `form:"productId" json:"productId"`
	StoreName       string   `form:"storeName" json:"storeName"`
	ProductName     string   `form:"productName" json:"productName"`
	ProductCategory string   `form:"productCategory" json:"productCategory"`
	ProductStock    int      `form:"productStock" json:"productStock"`
	ProductPrice    float64  `form:"productPrice" json:"productPrice"`
	ProductRate     float32  `form:"productRate" json:"productRate"`
	ProductPicture  []string `form:"productPicture" json:"productPicture"`
}

type PostProduct struct {
	SellerId        int      `form:"sellerId" json:"sellerId"`
	ProductName     string   `form:"productName" json:"productName"`
	ProductCategory int      `form:"productCategory" json:"productCategory"`
	ProductStock    int      `form:"productStock" json:"productStock"`
	ProductPrice    float64  `form:"productPrice" json:"productPrice"`
	ProductPicture  []string `form:"productPicture" json:"productPicture"`
}

//-----------------------------------------------------
// Getter Setter Product
func (p *Product) GetProductId() int {
	return p.ProductId
}

func (p *Product) SetProductId(productId int) {
	p.ProductId = productId
}

func (p *Product) GetSellerName() string {
	return p.StoreName
}

func (p *Product) SetSellerName(sellerName string) {
	p.StoreName = sellerName
}

func (p *Product) GetProductName() string {
	return p.ProductName
}

func (p *Product) SetProductName(productName string) {
	p.ProductName = productName
}

func (p *Product) GetProductCategory() string {
	return p.ProductCategory
}

func (p *Product) SetProductCategory(productCategory string) {
	p.ProductCategory = productCategory
}

func (p *Product) GetProductStock() int {
	return p.ProductStock
}

func (p *Product) SetProductStock(productStock int) {
	p.ProductStock = productStock
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

func (p *Product) GetProductPicture() []string {
	return p.ProductPicture
}

func (p *Product) SetProductPicture(productPicture []string) {
	p.ProductPicture = productPicture
}

//-----------------------------------------------------
// Getter Setter PostProduct
func (p *PostProduct) GetSellerId() int {
	return p.SellerId
}

func (p *PostProduct) SetSellerId(sellerName int) {
	p.SellerId = sellerName
}

func (p *PostProduct) GetProductName() string {
	return p.ProductName
}

func (p *PostProduct) SetProductName(productName string) {
	p.ProductName = productName
}

func (p *PostProduct) GetProductCategory() int {
	return p.ProductCategory
}

func (p *PostProduct) SetProductCategory(productCategory int) {
	p.ProductCategory = productCategory
}

func (p *PostProduct) GetProductStock() int {
	return p.ProductStock
}

func (p *PostProduct) SetProductStock(productStock int) {
	p.ProductStock = productStock
}

func (p *PostProduct) GetProductPrice() float64 {
	return p.ProductPrice
}

func (p *PostProduct) SetProductPrice(productPrice float64) {
	p.ProductPrice = productPrice
}

func (p *PostProduct) GetProductPicture() []string {
	return p.ProductPicture
}

func (p *PostProduct) SetProductPicture(productPicture []string) {
	p.ProductPicture = productPicture
}
