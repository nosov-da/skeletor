package service

import (
	"context"
	"fmt"

	"github.com/yourorg/yourproject/internal/entity"
	"github.com/yourorg/yourproject/internal/repository"
)

type UserService interface {
	Get(ctx context.Context, id int64) (*entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Get(ctx context.Context, id int64) (*entity.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get user %d: %w", id, err)
	}

	return u, nil
}
