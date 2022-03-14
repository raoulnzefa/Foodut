package controllers

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var con *gorm.DB

func connect() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level  // Ignore ErrRecordNotFound error for logger
			Colorful:      true,          // Disable color
		},
	)
	dsn := "root@tcp(localhost:3306)/db_foodut?charset=utf8mb4&parseTime=true&loc=Asia%2FJakarta"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func GetConnection() *gorm.DB {
	if con == nil {
		con = connect()
	}
	return con
}

func CloseConnection(con *gorm.DB) {
	if con == nil {
		sqlDB, err := con.DB()

		if err != nil {
			log.Panic("failed to close connection", err)
		}

		sqlDB.Close()
	}
}

