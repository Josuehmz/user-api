package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Josuehmz/user-api/internal/domain"
)

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	user.ID = ""
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *UserUseCase) PatchUser(ctx context.Context, id string, fields map[string]interface{}) error {
	if id == "" {
		return errors.New("user ID is required")
	}
	return uc.repo.UpdatePartial(ctx, id, fields)
}
