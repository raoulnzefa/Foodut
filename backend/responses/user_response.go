package responses

import (
	dto "github.com/Foodut/backend/dto"
)

type UserResponse struct {
	Status  int        `form:"status" json:"status"`
	Message string     `form:"message" json:"message"`
	Data    []dto.User `form:"data" json:"data"`
}
