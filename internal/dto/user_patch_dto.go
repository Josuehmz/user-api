package dto

// UserPatchDTO representa los campos que pueden actualizarse parcialmente de un usuario.
// Se usa en operaciones PATCH para modificar solo los datos provistos.
type UserPatchDTO struct {
	Name  *string `json:"name,omitempty"`  // Nombre opcional
	Email *string `json:"email,omitempty"` // Email opcional
}

// ToUpdateMap convierte el DTO en un mapa con los campos que fueron proporcionados.
// Este mapa se usa para actualizar parcialmente el usuario en la base de datos.
func (u *UserPatchDTO) ToUpdateMap() map[string]interface{} {
	update := make(map[string]interface{})

	if u.Name != nil {
		update["name"] = *u.Name
	}
	if u.Email != nil {
		update["email"] = *u.Email
	}

	return update
}
