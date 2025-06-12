package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Josuehmz/user-api/internal/domain"
)

// UserUseCase implementa la lógica de negocio para operaciones sobre usuarios.
type UserUseCase struct {
	repo domain.UserRepository // Repositorio que accede a la persistencia.
}

// NewUserUseCase crea una nueva instancia del caso de uso inyectando el repositorio.
func NewUserUseCase(r domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: r}
}

// CreateUser valida y crea un nuevo usuario.
func (uc *UserUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	user.ID = ""
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return uc.repo.Create(ctx, user)
}

// GetAllUsers obtiene todos los usuarios desde el repositorio.
func (uc *UserUseCase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return uc.repo.GetAll(ctx)
}

// GetUserByID obtiene un usuario por su ID.
func (uc *UserUseCase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return uc.repo.GetByID(ctx, id)
}

// DeleteUser elimina un usuario por su ID.
func (uc *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

// PatchUser actualiza parcialmente un usuario con los campos proporcionados.
func (uc *UserUseCase) PatchUser(ctx context.Context, id string, fields map[string]interface{}) error {
	if id == "" {
		return errors.New("user ID is required")
	}
	return uc.repo.UpdatePartial(ctx, id, fields)
}
