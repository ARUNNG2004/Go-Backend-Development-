package routes

import (
	"go-user-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handler.UserHandler) {
	app.Post("/users", h.CreateUser)
	app.Get("/users/:id", h.GetUser)
	app.Put("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)
	app.Get("/users", h.ListUsers)
}
