package models

type User struct {
	ID           int    `form:"id" json:"id" gorm:"primaryKey"`
	Username     string `form:"username" json:"username"`
	Email        string `form:"email" json:"email"`
	Name         string `form:"name" json:"name"`
	Password     string `form:"password" json:"password"`
	ProfilePhoto string `form:"photo" json:"photo"`
	Level        int    `form:"level" json:"level"`
	Admin        Admin  `gorm:"foreignkey:UserId;references:ID"`
}

type UserResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
