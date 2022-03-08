package models

type Category struct {
	CategoryID      int       `form:"id" json:"id" gorm:"primaryKey"`
	ProductCategory string    `form:"productCategory" json:"productCategory"`
	Product         []Product `gorm:"foreignKey:CategoryID;references:CategoryID"`
}

type CategoryResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []Product `form:"data" json:"data"`
}
