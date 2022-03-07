package models

type DetailProduct struct {
	ProductId string `form:"productId json:"productId" gorm:"primaryKey"`
	Quantity  int    `form:"quantity json:"quantity"`
	Product   Product
}

type DetailProductRespone struct {
	Status  int             `form:"status" json:"status"`
	Message string          `form:"message" json:"message"`
	Data    []DetailProduct `form:"data" json:"data"`
}
