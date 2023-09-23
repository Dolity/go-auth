package routes

import (
	"github.com/Dolity/go-jwt/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/v1/register", controller.Register)
	app.Post("/api/v1/login", controller.Login)
	app.Get("/api/v1/user", controller.User)
	app.Post("/api/v1/logout", controller.Logout)

}