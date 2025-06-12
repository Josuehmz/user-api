package domain

import "time"

// User representa la entidad del usuario dentro del dominio.
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
