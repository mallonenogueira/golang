package services

import (
	"go-api-crud/errors"
	"go-api-crud/models"
	"go-api-crud/repositories"
)

type UserEntryDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	repository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) GetAllUsers() []models.User {
	return s.repository.GetAll()
}

func (s *UserService) GetUserByID(id string) (*models.User, errors.AppError) {
	if id == "" {
		return nil, errors.NewValidationError("Id é obrigatório.")
	}

	user, exists := s.repository.GetById(id)

	if !exists {
		return nil, errors.NewNotFoundError("Usuário não encontrado.", "User")
	}

	return &user, nil
}

func (u *UserService) CreateUser(userDTO *UserEntryDTO) (*models.User, errors.AppError) {
	user, err := models.NewUser(userDTO.Name, userDTO.Email)

	if err != nil {
		return nil, err
	}

	u.repository.Insert(user)

	return user, nil
}

func (u *UserService) UpdateUser(id string, userEntry *UserEntryDTO) errors.AppError {
	user, err := models.UpdateUser(id, userEntry.Name, userEntry.Email)

	if err != nil {
		return err
	}

	if !u.repository.Update(user) {
		return errors.NewNotFoundError("Usuário não encontrado.", "User")
	}

	return nil
}

func (u *UserService) DeleteUser(id string) errors.AppError {
	if id == "" {
		return errors.NewValidationError("Id é obrigatório.")
	}

	if !u.repository.Delete(id) {
		return errors.NewNotFoundError("Usuário não encontrado.", "User")
	}

	return nil
}
