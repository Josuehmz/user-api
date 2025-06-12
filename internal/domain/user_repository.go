package domain

import "context"

// UserRepository define las operaciones de acceso a datos para usuarios.
type UserRepository interface {
	// Create agrega un nuevo usuario.
	Create(ctx context.Context, user *User) error

	// GetAll retorna todos los usuarios.
	GetAll(ctx context.Context) ([]User, error)

	// GetByID busca un usuario por ID.
	GetByID(ctx context.Context, id string) (*User, error)

	// Update reemplaza los datos de un usuario existente.
	Update(ctx context.Context, user *User) error

	// Delete elimina un usuario por ID.
	Delete(ctx context.Context, id string) error

	// UpdatePartial actualiza campos específicos de un usuario.
	UpdatePartial(ctx context.Context, id string, fields map[string]interface{}) error
}
