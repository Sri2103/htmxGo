package service

import (
	"context"

	userModel "github.com/sri2103/htmx_go/internal/userAuth/model"
	"github.com/sri2103/htmx_go/internal/userAuth/repository"
)

type IService interface {
	Login(context.Context, string, string) error
	Register(ctx context.Context, u *userModel.User) (*userModel.User, error)
	FindUser(context.Context, string) (*userModel.User, error)
}

type service struct {
	repo repository.IRepository
}

func NewUserService(repo repository.IRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Login(ctx context.Context, email, password string) error {
	return s.repo.Login(ctx, email, password)
}

func (s *service) Register(ctx context.Context, u *userModel.User) (*userModel.User, error) {
	return s.repo.Register(ctx, u)
}

func (s *service) FindUser(ctx context.Context, email string) (*userModel.User, error) {
	return s.repo.FindUser(ctx, email)
}
