package main

import (
	"github.com/gofiber/fiber/v2"
	"projects/APIWithMongoDB_2/app"
	"projects/APIWithMongoDB_2/configs"
	"projects/APIWithMongoDB_2/repository"
	"projects/APIWithMongoDB_2/services"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")

	todoRepositoryDB := repository.NewTodoRepositoryDB(dbClient)

	td := app.TodoHandler{Service: services.NewTodoService(todoRepositoryDB)}
	appRoute.Post("/api/todo", td.Insert)
	appRoute.Get("/api/todos", td.GetAll)
	appRoute.Delete("api/todo/:id", td.Delete)

	appRoute.Listen(":8080")
}
