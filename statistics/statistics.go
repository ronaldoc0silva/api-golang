package statistics

import (
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

// Instancia o contador
var counter = &Counter{}
