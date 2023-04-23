package service

import (
	"context"

	"github.com/aliocode/relic-example/internal/domain/user"
	"github.com/google/uuid"
)

type Repo interface {
	CreateUser(ctx context.Context, user user.User) error
	FetchByEmail(ctx context.Context, email string) (user.User, error)
}

type Service struct {
	repo Repo
}

func NewService(repo Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, name string, email string) (user.User, error) {
	newUser := user.User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
	if err := s.repo.CreateUser(ctx, newUser); err != nil {
		return user.User{}, err
	}

	return newUser, nil
}

func (s *Service) FetchByEmail(ctx context.Context, email string) (user.User, error) {
	res, err := s.repo.FetchByEmail(ctx, email)
	if err != nil {
		return user.User{}, err
	}
	return res, nil
}
