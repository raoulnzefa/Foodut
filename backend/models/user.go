package models

type User struct {
	Username     string `form:"username" json:"username" gorm:"primaryKey"`
	Email        string `form:"email" json:"email"`
	Name         string `form:"name" json:"name"`
	Password     string `form:"password" json:"password"`
	ProfilePhoto string `form:"photo" json:"photo"`
}

type UserResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
