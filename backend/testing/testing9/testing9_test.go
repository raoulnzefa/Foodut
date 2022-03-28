package testing8

import (
	"testing"

	dbController "github.com/Foodut/backend/database"
	prRepo "github.com/Foodut/backend/modules/product/repository"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	con := dbController.GetConnection()

	require.NoError(t, con.Error)
}

func TestProductDetailForCart(t *testing.T) {

	temp := prRepo.ReadCartProductDetailsAssociation(3)

	t.Errorf("DetailProduct : %+v", temp)
}
