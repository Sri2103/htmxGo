package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sri2103/htmx_go/internal/pkg/database"
	"github.com/sri2103/htmx_go/internal/todo/model"
	"github.com/sri2103/htmx_go/internal/todo/repository/query.go"
)

type IRepository interface {
	CreateTodo(context.Context, *model.Todo) (int, error)
	GetTodoById(context.Context, int) (*model.Todo, error)
	ReadTodos(id int) ([]*model.Todo, error)
	UpdateTodo(int, *model.Todo) error
	DeleteTodo(context.Context, int) error
}

type repo struct {
	db *database.DB
}

func NewRepo(db *database.DB) *repo {
	return &repo{
		db: db,
	}
}

// Create a new Todo item in the database
func (r *repo) CreateTodo(ctx context.Context, t *model.Todo) (int, error) {
	Result := r.db.Conn.QueryRow(query.CreateTodo, t.Title, t.Description, t.Status,t.UserID)
	err := Result.Scan(&t.ID) // get id of created todo
	if err != nil {
		return 0, fmt.Errorf("cannot get LastInsertId %w", err)
	}
	return int(t.ID), nil
}

func (r *repo) ReadTodos(id int) ([]*model.Todo, error) {
	var todos []*model.Todo
	rows, err := r.db.Conn.QueryContext(context.Background(), query.GetTodo,id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var todo model.Todo
		if err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)

	}
	return todos, nil

}

// get a single Todo ById
func (r *repo) GetTodoById(ctx context.Context, id int) (*model.Todo, error) {
	var todo model.Todo
	row := r.db.Conn.QueryRow(query.GetTodoById, id)
	
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status)
	if err != nil {
		return nil, fmt.Errorf("cannot scan data from db: %v",err)
	}
	return &todo, nil
}

// update todo
func (r *repo) UpdateTodo(id int, t *model.Todo) error {
	_, err := r.GetTodoById(context.Background(), id)
	if err == sql.ErrNoRows {
		return sql.ErrNoRows
	}

	_, err = r.db.Conn.Exec(query.UpdateTodoData, t.Title, t.Description,id)

	if err != nil {
		return err
	}

	return nil

}

// delete Todo
func (r *repo) DeleteTodo(ctx context.Context, id int) error {
	_, err := r.GetTodoById(context.Background(), id)
	if err == sql.ErrNoRows {
		return sql.ErrNoRows
	}
	res, err := r.db.Conn.Exec(query.DeleteTodo, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no record deleted")
	}
	return nil
}
