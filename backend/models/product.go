package models

type Product struct {
	ID              string   `form:"id" json:"id" gorm:"primaryKey"`
	ProductName     string   `form:"productName" json:"productName"`
	ProductPicture  []string `form:"productPicture" json:"productPicture"`
	ProductPrice    float64  `form:"productPrice" json:"productPrice"`
	ProductRate     float32  `form:"productRate" json:"productRate"`
	ProductStock    int16    `form:"productStock" json:"productStock"`
	SellerName      string   `form:"seller" json:"seller"`
	ProductCategory string   `form:"productCategory" json:"productCategory"`
	Category        Category `gorm:"foreignKey:ProductCategory"`
}

type ProductResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []Product `form:"data" json:"data"`
}
