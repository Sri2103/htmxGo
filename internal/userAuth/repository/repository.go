package repository

import (
	"context"

	userModel "github.com/sri2103/htmx_go/internal/userAuth/model"
)

type IRepository interface {
	Login(context.Context, string, string) error
	Register(context.Context, *userModel.User) error
	FindUser(context.Context, string) (*userModel.User, error)
}

type repo struct {
}

func (r *repo) Login(context.Context, string, string) error {
	return nil
}

func (r *repo) RegisterUser(context.Context, string) (*userModel.User, error) {
	return &userModel.User{}, nil
}

func (r *repo) FindUser(context.Context, string) (*userModel.User, error) {
	return &userModel.User{}, nil
}
