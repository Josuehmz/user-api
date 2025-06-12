package dto

type UserPatchDTO struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

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
