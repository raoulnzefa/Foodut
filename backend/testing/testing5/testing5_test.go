package testing5

import (
	"testing"

	dbController "github.com/Foodut/backend/database"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	tc "github.com/Foodut/backend/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	con := dbController.GetConnection()

	require.NoError(t, con.Error)
}

const N = "test Duplicate Full Name"
const U = "test Duplicate User Name"
const P = "test DuplicatePassword"
const S = "test Duplicate Store Name"
const C = "test Duplicate City"

var E = tc.StringWithCharset(11)

func TestDuplicateStore(t *testing.T) {
	dto := dto.PostSeller{
		Name:      N,
		Username:  U,
		Email:     E,
		Password:  P,
		StoreName: N,
		City:      C,
	}

	result := srvc.MapToSeller(dto)

	assert.ErrorContains(t, result.Error, "idx_store_name")
}
