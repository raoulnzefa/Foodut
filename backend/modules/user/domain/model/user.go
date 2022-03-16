package model

type User struct {
	ID           int    `form:"id" json:"id"gorm:"primaryKey"`
	Username     string `form:"username" json:"username"`
	Email        string `form:"email" json:"email" gorm:"uniqueIndex:idx_email"`
	Name         string `form:"name" json:"name"`
	Password     string `form:"password" json:"password"`
	ProfilePhoto string `form:"profilePhoto" json:"profilePhoro"`
	Level        int    `form:"level" json:"level"`
}
