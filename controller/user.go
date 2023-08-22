package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/saza-aluminium/service"
)

type User struct {
}

func NewUserController(userService service.UserService) UserController {
	return &User{
	}
}

func (controller *User) FindAllUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
}

func (controller *User) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
}