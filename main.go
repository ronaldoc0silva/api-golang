package main

import (
	"fiber-mongo-api/configs"

	"fiber-mongo-api/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	// Inicia o servidor na porta 8080
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}

}
