package model

import (
	prModel "github.com/Foodut/backend/modules/product/domain/model"
)

type User struct {
	ID            int                   `form:"id" json:"id" gorm:"primaryKey"`
	Username      string                `form:"username" json:"username"`
	Email         string                `form:"email" json:"email" gorm:"uniqueIndex:idx_email"`
	Name          string                `form:"name" json:"name"`
	Password      string                `form:"password" json:"password"`
	ProfilePhoto  string                `form:"profilePhoto" json:"profilePhoto"`
	Level         int                   `form:"level" json:"level"`
	Admin         Admin                 `gorm:"foreignkey:ID;references:ID"`
	Customer      Customer              `gorm:"foreignkey:ID;references:ID"`
	Seller        Seller                `gorm:"foreignkey:ID;references:ID"`
	DetailProduct prModel.DetailProduct `gorm:"foreignkey:UserID;references:ID"`
}
