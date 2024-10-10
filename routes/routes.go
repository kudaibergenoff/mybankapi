package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kudaibergenoff/mybankapi/internal/http/handlers"
)

type Route struct {
	App *fiber.App
}

func NewRoute(app *fiber.App) *Route {
	return &Route{App: app}
}

func (r *Route) Register() {
	accountRoutes := r.App.Group("/accounts")

	accountRoutes.Get("/", handlers.GetAccounts)
	//accountRoutes.Post("/", handlers.CreateAccount)
	//accountRoutes.Get("/:id", handlers.GetAccountByID)
	//accountRoutes.Put("/:id", handlers.UpdateAccount)
	//accountRoutes.Delete("/:id", handlers.DeleteAccount)
}
