package model

type Admin struct {
	UserID int  `form:"userId" json:"userId" gorm:"primaryKey"`
	User   User `form:"user" json:"user" gorm:"foreignKey:UserID;references:ID"`
}
