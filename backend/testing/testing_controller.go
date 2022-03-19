package testing

import (
	"math/rand"
	"time"

	dbController "github.com/Foodut/backend/database"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int) string {

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 -_!@#$%^&*"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func CountTableLength(tableName string) int {
	var rv int

	con := dbController.GetConnection()
	con.Raw("SELECT COUNT('id') FROM ?", tableName).Scan(&rv)

	return rv
}
