package testing5

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

const N = "test Full Name"
const U = "test User Name"
const E = "test Email12"
const P = "test Password"

func TestDuplicateEmailAdmin(t *testing.T) {
	usrDto := dto.PostUser{
		Name:     N,
		Username: U,
		Email:    E,
		Password: P,
	}

	result := srvc.MapToAdmin(usrDto)

	assert.ErrorContains(t, result.Error, "fk_admins_user")
}

const S = "test Store Name1234"
const C = "test City"

func TestDuplicateEmailSeller(t *testing.T) {
	dto := dto.PostSeller{
		Name:      N,
		Username:  U,
		Email:     E,
		Password:  P,
		StoreName: S,
		City:      C,
	}

	result := srvc.MapToSeller(dto)

	assert.ErrorContains(t, result.Error, "fk_sellers_user")
}

func TestDuplicateEmailCustomer(t *testing.T) {
	usrDto := dto.PostUser{
		Name:     N,
		Username: U,
		Email:    E,
		Password: P,
	}

	result := srvc.MapToCustomer(usrDto)

	assert.ErrorContains(t, result.Error, "fk_customers_user")
}
