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

var n = tc.StringWithCharset(10)
var u = tc.StringWithCharset(10)
var e = tc.StringWithCharset(10)
var p = tc.StringWithCharset(10)

var beforeLength int
var afterLength int

func TestPostAdmin(t *testing.T) {

	con := dbController.GetConnection()

	con.Raw("SELECT COUNT('id') FROM admins").Scan(&beforeLength)

	dto := dto.PostUser{
		Name:     n,
		Username: u,
		Email:    e,
		Password: p,
	}

	result := srvc.SendForCreateAdmin(dto)

	con.Raw("SELECT COUNT('id') FROM admins").Scan(&afterLength)

	assert.NoError(t, result.Error)

}

func TestAdminTableLengthCount(t *testing.T) {
	assert.Equal(t, (beforeLength + 1), (afterLength))
}
