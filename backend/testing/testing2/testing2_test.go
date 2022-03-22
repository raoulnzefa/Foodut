package testing2

import (
	"testing"

	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/product/domain/model"
	repo "github.com/Foodut/backend/modules/product/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	con := dbController.GetConnection()

	require.NoError(t, con.Error)
}

func TestFindAllProduct(t *testing.T) {
	con := dbController.GetConnection()

	var expected []model.Product
	con.Find(&expected)

	actual := repo.ReadAllProducts(nil)

	assert.Equal(t, len(expected), len(actual))
}
