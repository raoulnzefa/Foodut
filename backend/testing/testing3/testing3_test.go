package testing3

import (
	"testing"

	dbController "github.com/Foodut/backend/database"
	srvc "github.com/Foodut/backend/modules/user/domain/service"
	dto "github.com/Foodut/backend/modules/user/rest-api/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	con := dbController.GetConnection()

	require.NoError(t, con.Error)
}

func TestPostAdmin(t *testing.T) {

	n, u, e, p :=
		"test Full Name",
		"test Username",
		"test Email2",
		"test Password"

	dto := dto.PostUser{
		Name:     n,
		Username: u,
		Email:    e,
		Password: p,
	}

	result := srvc.MapToAdmin(dto)

	assert.NoError(t, result.Error)

}
