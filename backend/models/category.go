package models

type Category struct {
	ProductCategory string `form:"productCategory" json:"productCategory" gorm:"primaryKey"`
}

type CategoryResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []Product `form:"data" json:"data"`
}
