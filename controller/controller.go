package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	FindAllUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
