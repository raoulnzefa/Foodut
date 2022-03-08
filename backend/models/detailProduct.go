package models

type DetailProduct struct {
	UserID    int `form:"id" json:"id" gorm:"primaryKey"`
	ProductID int `form:"productId" json:"productId"`
	Quantity  int `form:"quantity" json:"quantity"`
}

type DetailProductRespone struct {
	Status  int             `form:"status" json:"status"`
	Message string          `form:"message" json:"message"`
	Data    []DetailProduct `form:"data" json:"data"`
}
