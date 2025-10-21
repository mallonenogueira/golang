package controllers

import (
	"go-api-crud/api/http"
	"go-api-crud/repositories"
	"go-api-crud/services"
)

type PaginationResponse struct {
	Page  int         `json:"page"`
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

type UserController struct {
	service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		service: services.NewUserService(*repositories.NewMemoryUserRepository()),
	}
}

func (u *UserController) GetUsers(ctx *http.Context) {
	users := u.service.GetAllUsers()

	ctx.ResponseOk(&PaginationResponse{
		Page:  1,
		Total: len(users),
		Data:  users,
	})
}

func (u *UserController) GetUserByID(ctx *http.Context) {
	user, err := u.service.GetUserByID(ctx.PathValue("id"))

	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.ResponseOk(user)
}

func (u *UserController) CreateUser(ctx *http.Context) {
	var userDTO services.UserEntryDTO

	if !ctx.BodyJson(&userDTO) {
		return
	}

	user, err := u.service.CreateUser(&userDTO)

	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.ResponseCreated(user)
}

func (u *UserController) UpdateUser(ctx *http.Context) {
	var user services.UserEntryDTO

	if !ctx.BodyJson(&user) {
		return
	}

	err := u.service.UpdateUser(ctx.PathValue("id"), &user)

	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.ResponseNoContent()
}

func (u *UserController) DeleteUser(ctx *http.Context) {
	err := u.service.DeleteUser(ctx.PathValue("id"))

	if err != nil {
		ctx.HandleError(err)
		return
	}

	ctx.ResponseNoContent()
}
