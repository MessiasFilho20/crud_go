package main

import (
	"crud/controller"
	"crud/database"
	"crud/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/users", controller.CreateUser)
	app.Get("/users", controller.GetUsers)
	app.Get("/users/:id", controller.GetUser)
	app.Put("/users/:id", controller.UpdateUser)
	app.Delete("/users/:id", controller.Deleteuser)
}

func main() {

	app := fiber.New()

	database.ConectDB()

	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Erro ao migrar o banco de dados:", err)
	}

	SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
