package models

import (
	"api/service/schemas"
)

func (user User) FromStruct(s schemas.AddUserBody) User {
	return User{
		UserId:    0,
		Name:      s.Name,
		Email:     s.Email,
		CreatedAt: s.CreatedAt,
	}
}
