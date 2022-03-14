package model

type DetailProduct struct {
	UserID    int `form:"id" json:"id" gorm:"primaryKey;autoIncrement:false"`
	ProductID int `form:"productId" json:"productId" gorm:"primaryKey;autoIncrement:false"`
	Quantity  int `form:"quantity" json:"quantity"`
	Status    int `form:"status" json:"status"`
	// { 0: wishlist, 1: payed }
}

type DetailProductRespone struct {
	Status  int             `form:"status" json:"status"`
	Message string          `form:"message" json:"message"`
	Data    []DetailProduct `form:"data" json:"data"`
}
