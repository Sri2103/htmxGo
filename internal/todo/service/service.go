package service

import (
	"context"
	"fmt"

	"github.com/sri2103/htmx_go/internal/todo/model"
	"github.com/sri2103/htmx_go/internal/todo/repository"
)

type IService interface {
	CreateTodo(context.Context, *model.Todo) (int, error)
	GetTodoById(context.Context, int) (*model.Todo, error)
	ReadTodos() ([]*model.Todo, error)
	UpdateTodo(int, *model.Todo) error
	DeleteTodo(context.Context,int) error
}

type service struct {
	repo repository.IRepository
}

func New(repo repository.IRepository) *service {
	return &service{
		repo: repo,
	}
}

// implement IService on service

// CreateTodo service
func (s *service) CreateTodo(ctx context.Context, t *model.Todo) (int, error) {
	if t.Status == "" {
		t.Status = "pending"
	}
	if t == nil || t.Title == "" {
		return 0, fmt.Errorf("invalid data")
	}
	return s.repo.CreateTodo(ctx, t)

}

// getSingleTodo Item
func (s *service) GetTodoById(ctx context.Context, id int) (*model.Todo, error) {
	return s.repo.GetTodoById(ctx, id)
}

// ReadTodos service
func (s *service) ReadTodos() ([]*model.Todo, error) {
	return s.repo.ReadTodos()
}

// updateTodos service
func (s *service) UpdateTodo(id int, t *model.Todo) error {
	return s.repo.UpdateTodo(id, t)
}

// Delete todo
func (s *service) DeleteTodo(ctx context.Context, id int) error {
	return s.repo.DeleteTodo(ctx,id)
}
