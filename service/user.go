package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type UsersService struct {
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(DB *sql.DB, validate *validator.Validate) UserService {
	return &UsersService{
		DB:             DB,
		Validate:       validate,
	}
}

const SessionLockValues = 1

func (service *UsersService) FindAllUsers(ctx context.Context) error {

	return nil
}

type UserPromotion struct {
	Id    int64
	Email string
	Name  string
}

func (service *UsersService) CreateUserPromotion(ctx context.Context)  error {
	return  nil
}