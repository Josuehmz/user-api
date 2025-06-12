package usecase

import (
	"context"

	"github.com/Josuehmz/user-api/internal/domain"
)

// UserUseCaseInterface define el contrato que debe cumplir la lógica de negocio
// relacionada con usuarios, permitiendo su implementación desacoplada.
type UserUseCaseInterface interface {
	// CreateUser crea un nuevo usuario.
	CreateUser(ctx context.Context, user *domain.User) error

	// GetAllUsers retorna todos los usuarios registrados.
	GetAllUsers(ctx context.Context) ([]domain.User, error)

	// GetUserByID busca un usuario por su ID.
	GetUserByID(ctx context.Context, id string) (*domain.User, error)

	// DeleteUser elimina un usuario dado su ID.
	DeleteUser(ctx context.Context, id string) error

	// PatchUser actualiza parcialmente los datos de un usuario.
	PatchUser(ctx context.Context, id string, fields map[string]interface{}) error
}
