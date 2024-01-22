package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sri2103/htmx_go/internal/pkg/database"
	userModel "github.com/sri2103/htmx_go/internal/userAuth/model"
	"github.com/sri2103/htmx_go/internal/userAuth/repository/query"
)

type IRepository interface {
	Login(context.Context, string, string) error
	Register(ctx context.Context, u *userModel.User) (*userModel.User, error)
	FindUser(context.Context, string) (*userModel.User, error)
	DeleteUser(context.Context, int) error
}

type repo struct {
	db *database.DB
}

func NewRepo(db *database.DB) IRepository {
	return &repo{db: db}
}

func (r *repo) Login(ctx context.Context, email string, password string) error {
	var u userModel.User
	Result := r.db.Conn.QueryRowContext(ctx, query.GetUser, email)
	err := Result.Scan(&u.ID,&u.Name,&u.Email,&u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) Register(ctx context.Context, u *userModel.User) (*userModel.User, error) {

	res := r.db.Conn.QueryRowContext(ctx, query.CreateUser, u.Name, u.Email, u.Password)
	err := res.Scan(&u.ID, &u.Name, &u.Email,&u.Password)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) FindUser(ctx context.Context, email string) (*userModel.User, error) {
	var u userModel.User
	res := r.db.Conn.QueryRowContext(ctx, query.GetUser, email)
	err := res.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return &u, err
}


func (u *repo) DeleteUser(context.Context, int) error {
	return nil
}
