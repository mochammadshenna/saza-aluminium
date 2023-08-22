package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/saza-aluminium/controller"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.PanicHandler = httprouter.New().PanicHandler

	return router
}
