package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) error

	GetAll(ctx context.Context) ([]User, error)

	GetByID(ctx context.Context, id string) (*User, error)

	Update(ctx context.Context, user *User) error

	Delete(ctx context.Context, id string) error
	UpdatePartial(ctx context.Context, id string, fields map[string]interface{}) error
}
