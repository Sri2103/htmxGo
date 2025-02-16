package service

import (
	"context"

	"github.com/sri2103/htmx_go/internal/expense/model"
)

type IService interface {
	AddTransaction(context.Context, *model.Expense) error
	GetTransaction(context.Context, int) (*model.Expense, error)
	GetTransactions(context.Context) ([]*model.Expense, error)
	UpdateTransaction(context.Context, *model.Expense) error
	DeleteTransaction(context.Context, int) error
	GetCategories(context.Context) ([]*model.Category, error)
	GetCategory(context.Context, int) (*model.Category, error)
	AddCategory(context.Context, *model.Category) error
	UpdateCategory(context.Context, *model.Category) error
	DeleteCategory(context.Context, int) error
}
