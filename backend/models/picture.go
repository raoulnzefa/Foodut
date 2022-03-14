package models

type Picture struct {
	ID          int    `form:"id" json:"id" gorm:"primaryKey"`
	PicturePath string `form:"picturePath" json:"picturePath"`
	ProductID   int
}

type PictureResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []Product `form:"data" json:"data"`
}
