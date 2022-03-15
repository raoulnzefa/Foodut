package model

type User struct {
	ID           int `gorm:"primaryKey"`
	Username     string
	Email        string `gorm:"uniqueIndex:idx_email"`
	Name         string
	Password     string
	ProfilePhoto string
	Level        int
}
