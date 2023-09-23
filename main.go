package main

import (
	"github.com/Dolity/go-jwt/database"
	"github.com/Dolity/go-jwt/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

    app := fiber.New()
    app.Use(cors.New(cors.Config{
        AllowCredentials: true,
        AllowOrigins:    "http://localhost:5173",
    }))

    routes.Setup(app)  
    database.Connect()
 
    app.Listen(":8000")
}
