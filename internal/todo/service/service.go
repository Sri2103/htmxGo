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
	ReadTodos(int) ([]*model.Todo, error)
	ReadDoneTodos(id int) ([]*model.Todo, error)
	UpdateTodo(int, *model.Todo) error
	DeleteTodo(context.Context, int) error
	ToggleTodoStatus(context.Context, int, bool) error
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
func (s *service) ReadTodos(id int) ([]*model.Todo, error) {
	return s.repo.ReadTodos(id)
}

// ReadTodos service
func (s *service) ReadDoneTodos(id int) ([]*model.Todo, error) {
	return s.repo.ReadDoneTodos(id)
}

// updateTodos service
func (s *service) UpdateTodo(id int, t *model.Todo) error {
	return s.repo.UpdateTodo(id, t)
}

// Delete todo
func (s *service) DeleteTodo(ctx context.Context, id int) error {
	return s.repo.DeleteTodo(ctx, id)
}

func (s *service) ToggleTodoStatus(ctx context.Context, id int, status bool) error {
	return s.repo.ToggleTodoStatus(ctx, id, status)
}
