package repository

import (
	"context"
	"errors"

	"github.com/sri2103/htmx_go/internal/todo/model"
)

type testRepo struct{}

// ReadDoneTodos implements IRepository.
func (*testRepo) ReadDoneTodos(id int) ([]*model.Todo, error) {
	panic("unimplemented")
}

// ToggleTodoStatus implements IRepository.
func (*testRepo) ToggleTodoStatus(context.Context, int, bool) error {
	panic("unimplemented")
}

var data = []*model.Todo{}

func NewTestRepo() *testRepo {
	r := &testRepo{}
	return r
}

func (r *testRepo) CreateRecords() {
	// do nothing for now
	data = append(data, &model.Todo{
		ID: 1, Title: "Buy Groceries", Status: false,
	})
	data = append(data, &model.Todo{
		ID: 2, Title: "Buy Groceries-2", Status: false,
	})
	data = append(data, &model.Todo{
		ID: 3, Title: "Buy Groceries-3", Status: false,
	})
}

func (r *testRepo) CreateTodo(_ context.Context, _ *model.Todo) (int, error) {
	return 1, nil
}

func (r *testRepo) ReadTodos(_ int) ([]*model.Todo, error) {
	return data, nil
}

func (r *testRepo) UpdateTodo(int, *model.Todo) error {
	return nil
}

func (r *testRepo) DeleteTodo(context.Context, int) error {
	return nil
}

func (r *testRepo) GetTodoById(_ context.Context, id int) (*model.Todo, error) {
	for _, v := range data {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, errors.New("no data found")
}
