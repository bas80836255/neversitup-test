package user

import (
	"context"
)

//go:generate mockgen -source=./service.go -destination=./service_mock.go -package=user
type Service interface {
	GetUser(ctx context.Context, userID string) (User, error)
}

type service struct {
	repository Repository
}

func (s *service) GetUser(ctx context.Context, userID string) (User, error) {
	return s.repository.FindOne(ctx, userID)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
