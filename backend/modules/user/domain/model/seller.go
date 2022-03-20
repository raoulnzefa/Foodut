package model

import (
	prModel "github.com/Foodut/backend/modules/product/domain/model"
)

type Seller struct {
	UserID      int               `form:"userId" json:"userId" gorm:"primaryKey"`
	User        User              `form:"user" json:"user" gorm:"foreignKey:UserID;references:ID"`
	StoreName   string            `form:"storeName" json:"storeName" gorm:"uniqueIndex:idx_store_name"`
	City        string            `form:"city" json:"city"`
	ListProduct []prModel.Product `form:"listProduct" json:"listProduct" gorm:"foreignKey:SellerID;references:UserID"`
}
