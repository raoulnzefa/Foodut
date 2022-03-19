package testing3

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

var N = tc.StringWithCharset(10)
var U = tc.StringWithCharset(10)
var E = tc.StringWithCharset(10)
var P = tc.StringWithCharset(10)
var S = tc.StringWithCharset(10)
var C = tc.StringWithCharset(10)

var beforeLength int
var afterLength int

func TestPostSeller(t *testing.T) {

	con := dbController.GetConnection()

	con.Raw("SELECT COUNT('id') FROM sellers").Scan(&beforeLength)

	dto := dto.PostSeller{
		Name:      N,
		Username:  U,
		Email:     E,
		Password:  P,
		StoreName: S,
		City:      C,
	}

	result := srvc.MapToSeller(dto)

	con.Raw("SELECT COUNT('id') FROM sellers").Scan(&afterLength)

	assert.NoError(t, result.Error)

}

func TestSellerTableLengthCount(t *testing.T) {
	assert.Equal(t, (beforeLength + 1), (afterLength))
}
