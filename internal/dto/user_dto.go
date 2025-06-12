package dto

import (
	"time"

	"github.com/Josuehmz/user-api/internal/domain"
)

type UserDTO struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"createdat,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updatedat,omitempty"`
}

func FromDomain(u *domain.User) *UserDTO {
	return &UserDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToDomain(dto *UserDTO) *domain.User {
	return &domain.User{
		ID:        dto.ID,
		Name:      dto.Name,
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
