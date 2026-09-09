package repository

import (
	"context"

	"github.com/yourorg/yourproject/internal/entity"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	Create(ctx context.Context, u *entity.User) error
}

type userRepo struct {
	// db *pgxpool.Pool
}

func NewUserRepository( /* db *pgxpool.Pool */ ) UserRepository {
	return &userRepo{}
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	_ = ctx // будет использован при реализации запроса к БД

	// TODO: реализовать запрос к БД.
	return &entity.User{
		ID:    id,
		Email: "user@example.com",
	}, nil
}

func (r *userRepo) Create(ctx context.Context, u *entity.User) error {
	_ = ctx // будет использован при реализации запроса к БД

	// TODO: реализовать вставку в БД.
	return nil
}
