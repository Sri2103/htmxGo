package service

import (
	"context"

	userModel "github.com/sri2103/htmx_go/internal/userAuth/model"
)

type IService interface {
	Login(context.Context, string, string) error
	Register(context.Context, *userModel.User) error
	FindUser(context.Context, string) (*userModel.User, error)
}
