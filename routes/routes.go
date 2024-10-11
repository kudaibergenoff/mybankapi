package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kudaibergenoff/mybankapi/internal/http/controllers"
)

type Route struct {
	App *fiber.App
}

func NewRoute(app *fiber.App) *Route {
	return &Route{App: app}
}

func (r *Route) Register() {
	accountRoutes := r.App.Group("/accounts")

	accountRoutes.Get("/", controllers.GetAccounts)
	//accountRoutes.Post("/", controllers.CreateAccount)
	//accountRoutes.Get("/:id", controllers.GetAccountByID)
	//accountRoutes.Put("/:id", controllers.UpdateAccount)
	//accountRoutes.Delete("/:id", controllers.DeleteAccount)
}
