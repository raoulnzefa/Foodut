package responses

import (
	entity "github.com/Foodut/backend/entities"
)

type UserResponse struct {
	Status  int           `form:"status" json:"status"`
	Message string        `form:"message" json:"message"`
	Data    []entity.User `form:"data" json:"data"`
}
