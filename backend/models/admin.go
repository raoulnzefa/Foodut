package models

type Admin struct {
	User User `form:"user" json:"user" gorm:"primaryKey"`
}

type AdminResponse struct {
	Status  int     `form:"status" json:"status"`
	Message string  `form:"message" json:"message"`
	Data    []Admin `form:"data" json:"data"`
}
