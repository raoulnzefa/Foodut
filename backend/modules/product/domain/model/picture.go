package model

type Picture struct {
	ID          int    `form:"id" json:"id" gorm:"primaryKey"`
	PicturePath string `form:"picturePath" json:"picturePath"`
	ProductID   int    `form:"productId" json:"productId"`
}
