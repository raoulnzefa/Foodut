package testing8

import (
	"testing"

	dbController "github.com/Foodut/backend/database"
	prDto "github.com/Foodut/backend/modules/product/rest-api/dto"
	srvc "github.com/Foodut/backend/modules/transaction/domain/service"
	"github.com/Foodut/backend/modules/transaction/rest-api/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	con := dbController.GetConnection()

	require.NoError(t, con.Error)
}

func TestPostCart(t *testing.T) {

	var temp []prDto.PostProductDetail
	p := prDto.PostProductDetail{
		ProductId: 1,
		Quantity:  2,
	}

	p2 := prDto.PostProductDetail{
		ProductId: 4,
		Quantity:  2,
	}

	temp = append(temp, p)
	temp = append(temp, p2)

	cart := dto.PostCart{
		UserId:   6,
		Products: temp,
	}

	result := srvc.MapToCart(cart)

	require.NoError(t, result.Error)
}

func TestDuplicateEntry(t *testing.T) {

	var temp []prDto.PostProductDetail
	p := prDto.PostProductDetail{
		ProductId: 1,
		Quantity:  2,
	}

	p2 := prDto.PostProductDetail{
		ProductId: 4,
		Quantity:  2,
	}

	temp = append(temp, p)
	temp = append(temp, p2)

	cart := dto.PostCart{
		UserId:   6,
		Products: temp,
	}

	result := srvc.MapToCart(cart)

	assert.ErrorContains(t, result.Error, "Duplicate")
}
