package testing1

import (
	"testing"

	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/transaction/domain/model"
	repo "github.com/Foodut/backend/modules/transaction/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	con := dbController.GetConnection()

	require.NoError(t, con.Error)
}

func TestGetTransaction(t *testing.T) {
	con := dbController.GetConnection()

	var expected []model.Transaction
	con.Find(&expected)

	actual := repo.FindAllTransaction(nil)

	assert.Equal(t, len(expected), len(actual))
}
