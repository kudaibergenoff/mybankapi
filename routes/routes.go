package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kudaibergenoff/mybankapi/config"
	"github.com/kudaibergenoff/mybankapi/internal/http/controllers"
	"github.com/kudaibergenoff/mybankapi/internal/repositories"
	"github.com/kudaibergenoff/mybankapi/internal/services"
)

type Route struct {
	App *fiber.App
}

func NewRoute(app *fiber.App) *Route {
	return &Route{App: app}
}

func (r *Route) Register() {
	cfg := config.LoadConfig()
	db := config.NewDatabase(cfg).DB

	accountRepo := repositories.NewAccountRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	accountService := services.NewAccountService(accountRepo)
	transactionService := services.NewTransactionService(accountRepo, transactionRepo)

	accountController := controllers.NewAccountController(accountService)
	transactionController := controllers.NewTransactionController(transactionService)

	api := r.App.Group("/api")

	api.Post("/accounts", accountController.CreateAccount)
	api.Put("/accounts/:id", accountController.UpdateAccount)
	api.Delete("/accounts/:id", accountController.DeleteAccount)
	api.Get("/accounts/:id", accountController.GetAccountByID)

	api.Post("/accounts/:id/debit", transactionController.DebitAccount)
	api.Post("/accounts/:id/credit", transactionController.CreditAccount)
	api.Post("/transfer", transactionController.TransferFunds)

	api.Post("/accounts/:id/freeze", accountController.FreezeAccount)
	api.Post("/accounts/:id/unfreeze", accountController.UnfreezeAccount)
}
