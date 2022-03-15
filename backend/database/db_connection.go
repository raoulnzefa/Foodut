package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var con *gorm.DB

/**
  Open initialize db session based on dialector.
  If error occured, the program will panic!.
*/
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

/**
  This is 'singleton' like connection.
  If the connection is nil, then
  create a new connection.
*/
func GetConnection() *gorm.DB {
	if con == nil {
		con = connect()
	}
	return con
}

/**
  Close closes the database and prevents new queries from starting. Close then waits for all queries that have started processing on the server to finish.

  It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.
*/
func CloseConnection(con *gorm.DB) {
	if con == nil {
		sqlDB, err := con.DB()

		if err != nil {
			log.Panic("failed to close connection", err)
		}

		sqlDB.Close()
	}
}
