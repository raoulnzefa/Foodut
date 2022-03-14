package router

import (
	"github.com/gorilla/mux"
)

func LoadRouter(router *mux.Router) {
	UserRouter(router)
	ProductRouter(router)
	TransactionRouter(router)
}
