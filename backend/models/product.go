package models

type Product struct {
	ProductCategory string
	ProductId       string
	ProductName     string
	ProductPicture  []string
	ProductPrice    float64
	ProductRate     float32
	ProductStock    int16
	SellerName      string
}
