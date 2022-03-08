package models

type Picture struct {
	ID             int    `form:"id" json:"id" gorm:"primaryKey"`
	ProductPicture string `form:"productPicture" json:"productPicture"`
	ProductID      int
}

type PictureResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []Product `form:"data" json:"data"`
}
