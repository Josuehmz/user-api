package usecase

import (
	"context"

	"github.com/Josuehmz/user-api/internal/domain"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
	PatchUser(ctx context.Context, id string, fields map[string]interface{}) error
}
