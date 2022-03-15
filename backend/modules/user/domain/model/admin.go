package model

type Admin struct {
	UserID int  `gorm:"primaryKey"`
	User   User `gorm:"foreignKey:UserID;references:ID"`
}
