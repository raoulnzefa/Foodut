package model

type Admin struct {
	ID int `form:"id" json:"id" gorm:"primaryKey"`
}
