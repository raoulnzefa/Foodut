package models

type Admin struct {
	ID int `form:"id" json:"id" gorm:"primaryKey"`
}

type AdminResponse struct {
	Status  int     `form:"status" json:"status"`
	Message string  `form:"message" json:"message"`
	Data    []Admin `form:"data" json:"data"`
}
