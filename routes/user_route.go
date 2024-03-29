package routes

import (
	"fiber-mongo-api/controllers"
	"sync"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Counter armazena o número de solicitações
type Counter struct {
	count int
	mu    sync.Mutex
}

// Middleware para contar as solicitações
func (c *Counter) CountMiddleware(ctx *fiber.Ctx) error {
	// Incrementa o contador de solicitações
	c.mu.Lock()
	c.count++
	c.mu.Unlock()

	// Chama o próximo manipulador na cadeia
	return ctx.Next()
}

// Manipulador para a rota principal
func mainHandler(ctx *fiber.Ctx) error {
	return ctx.SendString("Olá, esta é a rota principal!")
}

// Manipulador para a rota de estatísticas
func statsHandler(ctx *fiber.Ctx, counter *Counter) error {
	// Recupera o contador de solicitações
	counter.mu.Lock()
	count := counter.count
	counter.mu.Unlock()

	// Retorna o número de solicitações
	return ctx.SendString(fmt.Sprintf("Número total de solicitações: %d", count))
}

func getAllUsersHandler(ctx *fiber.Ctx) error {
	// Lógica para obter todos os usuários
	return ctx.SendString("Listagem de todos os usuários")
}

// Instancia o contador
var counter = &Counter{}

func UserRoute(app *fiber.App) {

	app.Use(counter.CountMiddleware)

	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)

	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	//app.Get("/users", controllers.GetAllUsers)

	app.Get("/stats", func(ctx *fiber.Ctx) error {
		return statsHandler(ctx, counter)
	})

	app.Get("/users", controllers.GetAllUsers, getAllUsersHandler)

}
