package model

import (
	prModel "github.com/Foodut/backend/modules/product/domain/model"
	trModel "github.com/Foodut/backend/modules/transaction/domain/model"
)

type Customer struct {
	UserID      int                   `form:"userId" json:"userId" gorm:"primaryKey"`
	User        User                  `form:"user" json:"user" gorm:"foreignKey:UserID;references:ID"`
	Address     string                `form:"address" json:"address"`
	Cart        []prModel.Product     `form:"cart" json:"cart" gorm:"many2many:carts"`
	ListHistory []trModel.Transaction `form:"listHistory" json:"listHistory" gorm:"foreignKey:CustomerID;references:UserID"`
}
