package models

import (
	"go-api-crud/errors"
	"math/rand"
	"strconv"
)

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateId() string {
	return strconv.FormatUint(rand.Uint64(), 10)
}

func NewUser(name string, email string) (*User, errors.AppError) {
	if name == "" {
		return nil, &errors.FieldValidationError{Field: "name", Message: "Nome é obrigatório."}
	}

	if email == "" {
		return nil, &errors.FieldValidationError{Field: "email", Message: "Email é obrigatório."}
	}

	return &User{CreateId(), name, email}, nil
}

func UpdateUser(id string, name string, email string) (*User, errors.AppError) {
	if id == "" {
		return nil, &errors.FieldValidationError{Field: "id", Message: "Id é obrigatório."}
	}

	if name == "" {
		return nil, &errors.FieldValidationError{Field: "name", Message: "Nome é obrigatório."}
	}

	if email == "" {
		return nil, &errors.FieldValidationError{Field: "email", Message: "Email é obrigatório."}
	}

	return &User{id, name, email}, nil
}
