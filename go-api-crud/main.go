package main

import (
	apiHttp "go-api-crud/api/http"
	"go-api-crud/controllers"
	"log"
	"net/http"
)

func main() {
	userController := controllers.NewUserController()

	apiHttp.NewRouter().
		Get("/users/{id}", userController.GetUserByID).
		Get("/users", userController.GetUsers).
		Post("/users", userController.CreateUser).
		Put("/users", userController.UpdateUser).
		Delete("/users/{id}", userController.DeleteUser).
		Register()

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
