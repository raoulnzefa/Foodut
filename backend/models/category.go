package models

type Category struct {
	ProductCategory string `form:"productId" json:"productId" gorm:"primaryKey"`
}

type CategoryResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []Product `form:"data" json:"data"`
}
