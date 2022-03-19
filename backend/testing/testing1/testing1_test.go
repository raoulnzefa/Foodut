package testing1

import (
	"testing"

	dbController "github.com/Foodut/backend/database"
)

func TestDatabaseConnection(t *testing.T) {
	con := dbController.GetConnection()

	if con.Error != nil {
		t.Fatalf(
			"Error (%+v) ||| \n"+
				"Expected to Connect",
			con.Error,
		)
	}
}
