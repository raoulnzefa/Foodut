package model

type Category struct {
	ID              int       `form:"id" json:"id" gorm:"primaryKey"`
	ProductCategory string    `form:"productCategory" json:"productCategory"`
	Product         []Product `gorm:"foreignKey:CategoryID;references:ID"`
}
