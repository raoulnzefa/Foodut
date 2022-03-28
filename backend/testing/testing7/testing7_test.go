package testing7

import (
	"testing"

	dbController "github.com/Foodut/backend/database"
	model "github.com/Foodut/backend/modules/user/domain/model"
	repo "github.com/Foodut/backend/modules/user/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	con := dbController.GetConnection()

	require.NoError(t, con.Error)
}

func TestGetStore(t *testing.T) {

	con := dbController.GetConnection()

	storeName := "Lumpia Pencit"

	var expected model.Seller
	con.Where("store_name = ?", storeName).First(&expected)

	actual := repo.ReadSellerByStoreName([]string{storeName})

	assert.True(t, actual.UserID != 0)

	assert.Equal(t, expected.StoreName, actual.StoreName)
}
