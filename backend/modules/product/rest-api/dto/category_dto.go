package dto

type Category struct {
	ProductCategory string `form:"productCategory" json:"productCategory"`
}

type GetCategory struct {
	ID              int    `form:"id" json:"id" gorm:"primaryKey"`
	ProductCategory string `form:"productCategory" json:"productCategory"`
}
