package user

import (
	"context"
	"errors"
	"time"
)

//go:generate mockgen -source=./repository.go -destination=./repository_mock.go -package=user
type Repository interface {
	FindOne(ctx context.Context, userID string) (User, error)
}

type repository struct {
	mapUsers map[string]User
}

func (r *repository) FindOne(_ context.Context, userID string) (User, error) {
	user, ok := r.mapUsers[userID]
	if !ok {
		return user, errors.New("not found")
	}

	return user, nil
}

func NewRepository() Repository {
	adminUser := User{
		ID:         0,
		Name:       "admin",
		CreateDate: time.Now(),
		CreateBy:   "system",
	}

	user1 := User{
		ID:         1,
		Name:       "user1",
		CreateDate: time.Now(),
		CreateBy:   "system",
	}

	return &repository{
		mapUsers: map[string]User{"0": adminUser, "1": user1},
	}
}
