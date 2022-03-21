package testing8

import (
	"testing"

	dbController "github.com/Foodut/backend/database"
	prSrvc "github.com/Foodut/backend/modules/product/domain/service"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	con := dbController.GetConnection()

	require.NoError(t, con.Error)
}

func TestHitoryAssociation(t *testing.T) {

	temp := prSrvc.ProductDetailAssociationWithTransaction(1)

	t.Errorf("Detail Product : %+v", temp)
}
