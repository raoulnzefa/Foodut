package models

type Product struct {
	ProductId       string   `form:"productId" json:"productId" gorm:"primaryKey"`
	ProductCategory Category `form:"productCategory" json:"productCategory"`
	ProductName     string   `form:"productName" json:"productName"`
	ProductPicture  []string `form:"productPicture" json:"productPicture"`
	ProductPrice    float64  `form:"productPrice" json:"productPrice"`
	ProductRate     float32  `form:"productRate" json:"productRate"`
	ProductStock    int16    `form:"productStock" json:"productStock"`
	SellerName      string   `form:"seller" json:"seller"`
}

type ProductResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []Product `form:"data" json:"data"`
}
